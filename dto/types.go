package dto

import (
	"strconv"
	"sync"
)

type (
	TokenPair struct {
		Symbol		string	`json:"symbol"`
		Price		string	`json:"price"`
		Sma5		[]float64
		Sma10		[]float64
		Sma5Avg		float64
		Sma10Avg	float64
		Mtx			sync.Mutex
	}
)

func (pair *TokenPair) PushPriceToSma(price string) {
	var priceFloat float64

	priceFloat, _ = strconv.ParseFloat(price, 64)

	if len(pair.Sma5) == 5 {
		pair.Sma5 = pair.Sma5[1:]
	}
	if len(pair.Sma10) == 10 {
		pair.Sma10 = pair.Sma10[1:]
	}
	pair.Sma5 = append(pair.Sma5, priceFloat)
	pair.Sma10 = append(pair.Sma10, priceFloat)
}

func (pair TokenPair)UpdatePrice(price string) error {
	pair.Price = price
	return nil
}
