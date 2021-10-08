package gamedb

type gameConsume struct {
	InternalName       string  `json:"internalName"`
	Title              string  `json:"title"`
	MetacriticLink     string  `json:"metacriticLink"`
	DealID             string  `json:"dealID"`
	StoreID            string  `json:"storeID"`
	GameID             string  `json:"gameID"`
	SalePrice          float64 `json:"salePrice"`
	normalPrice        float64 `json:"normalPrice"`
	IsOnSale           float64 `json:"isOnSale"`
	MetacriticScore    string  `json:"metacriticScore"`
	steamRatingText    string  `json:"steamRatingText"`
	steamRatingCount   string  `json:"steamRatingPercent"`
	steamRatingPercent string  `json:"steamRatingCount"`
	steamAppID         string  `json:"steamAppID"`
	releaseDate        string  `json:"releaseDate"`
	lastChange         string  `json:"lastChange"`
	dealRating         string  `json:"dealRating"`
	thumb              string  `json:"thumb"`
}
