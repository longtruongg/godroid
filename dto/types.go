package dto

import "strconv"

type (
	TokenPair struct {
		Symbol		string	`json:"symbol"`
		Price		string	`json:"price"`
		Sma5		[]float64
		Sma10		[]float64
		Sma5Avg		float64
		Sma10Avg	float64
	}
)

func (pair *TokenPair) PushPriceToSma(price string) (err error) {
	var priceFloat float64
	if priceFloat, err = strconv.ParseFloat(price, 64); err != nil {
		return err
	}
	if len(pair.Sma5) == 5 {
		pair.Sma5 = pair.Sma5[1:]
	}
	if len(pair.Sma10) == 10 {
		pair.Sma10 = pair.Sma10[1:]
	}
	pair.Sma5 = append(pair.Sma5, priceFloat)
	pair.Sma10 = append(pair.Sma10, priceFloat)
	return nil
}
