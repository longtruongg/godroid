package dto

type (
	TokenPair struct {
		Symbol	string	`json:"symbol"`
		Price	string	`json:"price"`
		Sma5	[5]float32
		Sma10	[10]float32
	}
)
