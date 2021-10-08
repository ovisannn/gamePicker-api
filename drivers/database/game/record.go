package gamedb

type GameConsume struct {
	InternalName       string  `json:"internalName"`
	Title              string  `json:"title"`
	MetacriticLink     string  `json:"metacriticLink"`
	DealID             string  `json:"dealID"`
	StoreID            string  `json:"storeID"`
	GameID             string  `json:"gameID"`
	SalePrice          float64 `json:"salePrice"`
	NormalPrice        float64 `json:"normalPrice"`
	IsOnSale           float64 `json:"isOnSale"`
	MetacriticScore    string  `json:"metacriticScore"`
	SteamRatingText    string  `json:"steamRatingText"`
	SteamRatingCount   string  `json:"steamRatingPercent"`
	SteamRatingPercent string  `json:"steamRatingCount"`
	SteamAppID         string  `json:"steamAppID"`
	ReleaseDate        string  `json:"releaseDate"`
	LastChange         string  `json:"lastChange"`
	DealRating         string  `json:"dealRating"`
	Thumb              string  `json:"thumb"`
}

type GameInfo struct {
	GameID         string `json:"gameID"`
	SteamAppID     string `json:"steamAppID"`
	Cheapest       string `json:"cheapest"`
	CheapestDealID string `json:"cheapestDealID"`
	External       string `json:"external"`
	InternalName   string `json:"internalName"`
	Thumb          string `json:"thumb"`
}
