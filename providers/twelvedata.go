package providers

import "fmt"

type Twelvedata struct {

}

func (provider Twelvedata) GetValues(){
	fmt.Println("Hello from twelvedata")
}