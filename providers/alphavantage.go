package providers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Alphavantage struct {

}

func (provider Alphavantage) GetValues() map[string]float64 {
	jsonResult := request()

	result := map[string]float64{}

	for date, item := range jsonResult["Technical Analysis: EMA"].(map[string]interface{}){
		value := item.(map[string]interface{})["EMA"].(string)
		result[date], _ = strconv.ParseFloat(value, 64)
	}

	return result
}

func request() map[string]interface{} {
	response, err := http.Get("https://www.alphavantage.co/query?function=EMA&symbol=EURUSD&interval=15min&time_period=12&series_type=close&apikey=MVS8TN4ZTQF1ONES");
	if err != nil {
		log.Panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Panic("bad response from alphavantage")
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)

	var jsonResult map[string]interface{}
	json.Unmarshal([]byte(bodyString), &jsonResult)

	return jsonResult
}