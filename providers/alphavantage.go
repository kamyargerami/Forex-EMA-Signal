package providers

import (
	"bed/helpers"
	"log"
	"strconv"
)

type Alphavantage struct{}

func (provider Alphavantage) GetSymbols() []string {
	return []string{"EURUSD", "USDJPY", "GBPUSD", "AUDUSD", "USDCAD", "USDCHF", "NZDUSD", "EURGBP", "EURAUD", "EURCHF",
		"EURJPY", "EURNZD", "GBPEUR", "GBPJPY", "GBPAUD", "GBPCAD", "GBPCHF", "GBPNZD", "CADCHF", "CADJPY"}
}

func (provider Alphavantage) GetValues(symbol string, interval string, period int) finalResult {
	url := "https://www.alphavantage.co/query?function=EMA&symbol=" + symbol + "&interval=" + interval + "&time_period=" + strconv.Itoa(period) + "&series_type=close&apikey=" + helpers.Env("ALPHAVANTAGE_API_KEY")

	jsonResult := request(url)

	result := finalResult{}

	items, ok := jsonResult["Technical Analysis: EMA"]

	if !ok {
		log.Println("Bad Response from api")
		return result
	}

	for date, item := range items.(map[string]interface{}) {
		value := item.(map[string]interface{})["EMA"].(string)
		floatValue, _ := strconv.ParseFloat(value, 64)

		result = append(result, resultItem{
			date:  date,
			Value: floatValue,
		})
	}

	return result
}
