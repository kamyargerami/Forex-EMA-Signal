package main

import (
	"bed/providers"
	"bed/services"
	"fmt"
	"log"
	"time"
)

func main() {
	provider := providers.Twelvedata{}

	for _, symbol := range provider.GetSymbols() {
		result12Days := provider.GetValues(symbol, "15min", 12)
		result32Days := provider.GetValues(symbol, "15min", 32)

		if len(result12Days) == 0 || len(result32Days) == 0 {
			log.Println("No result from api")
			continue
		}

		sell := 0
		buy := 0

		for i := 0; i < 8; i++ {
			if result12Days[i].Value > result32Days[i].Value {
				buy++
			} else {
				sell++
			}
		}

		if sell != 0 && sell <= 2 {
			fmt.Println(symbol, sell, buy, "Sell Signal")
			services.Telegram{}.Send("Sell " + symbol)
		} else if buy != 0 && buy <= 2 {
			fmt.Println(symbol, sell, buy, "Buy Signal")
			services.Telegram{}.Send("Buy " + symbol)
		} else {
			fmt.Println(symbol, sell, buy)
		}

		time.Sleep(25 * time.Second)
	}
}
