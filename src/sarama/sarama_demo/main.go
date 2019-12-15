package main

import (
	"fmt"
	_ "github.com/Shopify/sarama"
)

func returnWhole(){
	for {
		return
	}
	fmt.Println("whole")
}
func main() {
	returnWhole()
}

