package main

import (
	"fmt"
	"godroid/dto"
	"godroid/utils"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				getTokenPrice()
			}
		}
	}()

	time.Sleep(1 * time.Hour)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func waitForerver() {
	select {}
}

func getTokenPrice(){
	for _, value := range dto.TokenList {
		pair, err := utils.GetTokenPair(dto.GetPriceURL, value)
		if err != nil {
			fmt.Println(err)
		}
		if err := pair.PushPriceToSma(pair.Price); err != nil {
			fmt.Println("cannot push price into sma")
		}
	}
}
