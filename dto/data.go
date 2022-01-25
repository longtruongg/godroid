package dto

var (
	TokenList []string
		GetPriceURL = "https://api.binance.com/api/v3/ticker/price?symbol="
)

func init() {
	TokenList = append(TokenList, "ETHUSDT")
	TokenList = append(TokenList, "NEARUSDT")
	TokenList = append(TokenList, "SOLUSDT")
	TokenList = append(TokenList, "DOTUSDT")
}
