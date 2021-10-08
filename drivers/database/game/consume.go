package gamedb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetOnSales() []GameConsume {
	response, _ := http.Get("https://www.cheapshark.com/api/1.0/deals?storeID=1&onSale=1")
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var result []GameConsume
	json.Unmarshal(responseData, &result)
	return result
}

func GetGameByName_deals(name string) []GameConsume {
	url := fmt.Sprintf("https://www.cheapshark.com/api/1.0/deals?title=%s", name)
	response, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var result []GameConsume
	json.Unmarshal(responseData, &result)
	return result
}

func GetGameByMaxPrice_deals(price float64) []GameConsume {
	priceStr := fmt.Sprintf("%f", price)
	url := fmt.Sprintf("https://www.cheapshark.com/api/1.0/deals?upperPrice=%s", priceStr)
	response, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var result []GameConsume
	json.Unmarshal(responseData, &result)
	return result
}

func GetGameByDealsID_deals(dealsID string) GameConsume {
	url := fmt.Sprintf("https://www.cheapshark.com/api/1.0/deals?id=%s", dealsID)
	response, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var result GameConsume
	json.Unmarshal(responseData, &result)
	return result
}

func GetGameByName_GameInfo(name string) []GameInfo {
	url := fmt.Sprintf("https://www.cheapshark.com/api/1.0/games?title=%s&limit=60&exact=0", name)
	response, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var result []GameInfo
	json.Unmarshal(responseData, &result)
	return result
}
