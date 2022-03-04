package services

import (
	"bed/helpers"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Telegram struct{}

func (service Telegram) Send(text string) map[string]interface{} {
	postBody, _ := json.Marshal(map[string]string{
		"text":    text,
		"chat_id": helpers.Env("TELEGRAM_CHAT_ID"),
	})

	response, err := http.Post("https://api.telegram.org/bot"+helpers.Env("TELEGRAM_BOT_TOKEN")+"/sendMessage", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		log.Panic(err)
	}

	if response.StatusCode != http.StatusOK {
		log.Panic("cannot post telegram")
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
