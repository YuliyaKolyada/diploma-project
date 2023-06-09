package SMSData

import (
	"github.com/YuliyaKolyada/diploma-project/pkg/references/countryReference"
	"github.com/YuliyaKolyada/diploma-project/pkg/references/providerReference"
	"strings"
)

type SMSData struct {
	Country      string `json:"country"`       // alpha-2 — код страны
	Bandwidth    string `json:"bandwidth"`     // пропускная способность канала от 0 до 100%
	ResponseTime string `json:"response_time"` // среднее время ответа в миллисекундах
	Provider     string `json:"provider"`      // название компании-провайдера
}

func NewSMSData(country string, bandwidth string, responseTime string, provider string) *SMSData {
	data := new(SMSData)
	data.Country = country
	data.Bandwidth = bandwidth
	data.ResponseTime = responseTime
	data.Provider = provider
	return data
}

// Возвращает список валидных данных о системе SMS
func Parse(fields []string) (*SMSData, bool) {
	if len(fields) != 4 {
		return nil, false
	}
	if !countryReference.Contains(fields[0]) {
		return nil, false
	}
	provider, ok := providerReference.Get(fields[3])
	if !ok || !provider.IsSMS {
		return nil, false
	}
	d := NewSMSData(fields[0], fields[1], fields[2], fields[3])
	return d, true
}

func (d *SMSData) ModifyCountry() {
	d.Country, _ = countryReference.Get(d.Country)
}

func SortByProvider(data []SMSData) []SMSData {
	for i := 1; i < len(data); i++ {
		j := i
		for j > 0 && strings.Compare(data[j].Provider, data[j-1].Provider) < 0 {
			data[j], data[j-1] = data[j-1], data[j]
			j--
		}
	}
	return data
}

func SortByCountry(data []SMSData) []SMSData {
	for i := 1; i < len(data); i++ {
		j := i
		for j > 0 && strings.Compare(data[j].Country, data[j-1].Country) < 0 {
			data[j], data[j-1] = data[j-1], data[j]
			j--
		}
	}
	return data
}
