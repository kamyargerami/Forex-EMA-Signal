package main

import (
	"bed/providers"
	"fmt"
	"os"
)

func main() {
	provider := providers.Alphavantage{}

	for _, symbol := range provider.GetSymbols() {
		result := provider.GetValues(symbol, "15min", 12)
		fmt.Println(len(result))
		os.Exit(100)
	}

	//r := i.GetValues("EURUSD", "15min", 12)
	//fmt.Println("done", len(r))
}
