package gamedb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetOnSales() []gameConsume {
	response, _ := http.Get("https://www.cheapshark.com/api/1.0/deals?storeID=1&onSale=1")
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var result []gameConsume
	json.Unmarshal(responseData, &result)
	return result
}

func GetGameByName(name string) []gameConsume {
	url := fmt.Sprintf("https://www.cheapshark.com/api/1.0/deals?title=%s", name)
	response, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var result []gameConsume
	json.Unmarshal(responseData, &result)
	return result
}
