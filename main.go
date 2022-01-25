package main

import (
	"fmt"
	"godroid/dto"
	"godroid/utils"
	"time"
)

func main() {
	tokenList := initTokenPairs()
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				getTokenPrice(tokenList)
			}
		}
	}()

	time.Sleep(1 * time.Hour)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func initTokenPairs() []*dto.TokenPair{
	var tokenList []*dto.TokenPair
	for _, symbol := range dto.TokenList {
		tok := dto.TokenPair{Symbol: symbol}
		tokenList = append(tokenList, &tok)
	}
	return tokenList
}

func waitForerver() {
	select {}
}

func getTokenPrice(list []*dto.TokenPair) []*dto.TokenPair{
	for _, value := range list {
		price, err := utils.GetTokenPair(dto.GetPriceURL, value.Symbol)
		if err != nil {
			fmt.Println(err)
		}
		value.Mtx.Lock()
		defer value.Mtx.Unlock()
		value.Price = price
		value.PushPriceToSma(price)

		fmt.Println(value.Symbol)
		for _, val := range value.Sma10{
			fmt.Printf("%f\n", val)
		}
	}
	return list
}
