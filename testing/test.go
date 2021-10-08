package main

import (
	"fmt"
	gamedb "gamePicker/drivers/database/game"
)

func main() {
	// fmt.Print(gamedb.GetOnSales())
	fmt.Print(gamedb.GetGameByName("ovi"))
}
