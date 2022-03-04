package providers

import (
	"bed/helpers"
	"strconv"
)

type Twelvedata struct{}

func (provider Twelvedata) GetSymbols() []string {
	return []string{"EUR/USD", "USD/JPY", "GBP/USD", "AUD/USD", "USD/CAD", "USD/CHF", "NZD/USD", "EUR/GBP", "EUR/AUD",
		"EUR/CHF", "EUR/JPY", "EUR/NZD", "GBP/EUR", "GBP/JPY", "GBP/AUD", "GBP/CAD", "GBP/CHF", "GBP/NZD",
		"CAD/CHF"}
}

func (provider Twelvedata) GetValues(symbol string, interval string, period int) finalResult {
	url := "https://api.twelvedata.com/ema?symbol=" + symbol + "&interval=" + interval + "&time_period=" + strconv.Itoa(period) + "&apikey=" + helpers.Env("TWELVEDATA_API_KEY")

	jsonResult := request(url)

	result := finalResult{}

	for _, item := range jsonResult["values"].([]interface{}) {
		date := item.(map[string]interface{})["datetime"].(string)
		value := item.(map[string]interface{})["ema"].(string)
		floatValue, _ := strconv.ParseFloat(value, 64)

		result = append(result, resultItem{
			date:  date,
			Value: floatValue,
		})
	}

	return result
}
