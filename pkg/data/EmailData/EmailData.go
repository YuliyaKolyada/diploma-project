package EmailData

import (
	"github.com/YuliyaKolyada/diploma-project/pkg/references/countryReference"
	"github.com/YuliyaKolyada/diploma-project/pkg/references/providerReference"
	"strconv"
)

type EmailData struct {
	Country      string `json:"country"`       // alpha-2 — код страны
	Provider     string `json:"provider"`      // название компании-провайдера
	DeliveryTime int    `json:"delivery_time"` // среднее время доставки писем в миллисекундах
}

func NewEmailData(country string, provider string, deliveryTime int) *EmailData {
	data := new(EmailData)
	data.Country = country
	data.Provider = provider
	data.DeliveryTime = deliveryTime
	return data
}

// Возвращает список валидных данных о системе Email
func Parse(fields []string) (*EmailData, bool) {
	if len(fields) != 3 {
		return nil, false
	}
	if !countryReference.Contains(fields[0]) {
		return nil, false
	}
	provider, ok := providerReference.Get(fields[1])
	if !ok || !provider.IsEmail {
		return nil, false
	}
	deliveryTime, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, false
	}

	d := NewEmailData(fields[0], fields[1], deliveryTime)
	return d, true
}
