package providers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type provider interface {
	GetSymbols()
	GetValues()
}

type finalResult []resultItem

type resultItem struct {
	date  string
	Value float64
}

func request(url string) map[string]interface{} {
	response, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}

	if response.StatusCode != http.StatusOK {
		log.Panic("cannot get " + url)
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)

	var jsonResult map[string]interface{}
	json.Unmarshal([]byte(bodyString), &jsonResult)

	return jsonResult
}
