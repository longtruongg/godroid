package main

import (
	"fmt"
	"godroid/dto"
	"godroid/utils"
)

func main() {
	//var tokenPairList []dto.TokenPair
	for _, value := range dto.TokenList {
		pair, err := utils.GetTokenPair(dto.GetPriceURL, value)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(pair.Price)
	}
}
