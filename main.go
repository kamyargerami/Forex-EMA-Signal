package main

import (
	"bed/providers"
	"fmt"
)

func main(){
	i := providers.Alphavantage{}
	r := i.GetValues()
	fmt.Println(r)
}
