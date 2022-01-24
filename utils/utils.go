package utils

import (
	"encoding/json"
	"godroid/dto"
	"io/ioutil"
	"net/http"
)

func GetTokenPair(url string, token string) (pair *dto.TokenPair, err error){
	result, err := http.Get(url + token)
	if err != nil {
		return  nil, err
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &pair); err != nil {
		return nil, err
	}

	return pair, nil
}

func CalculateSmaAvg(list []float64, num int) (smaAvg float64) {
	for i := 0; i < len(list); i++ {
		smaAvg += list[i]
	}
	switch num {
	case 5:
		smaAvg = smaAvg/5
	case 10:
		smaAvg = smaAvg/10
	}

	return smaAvg
}
