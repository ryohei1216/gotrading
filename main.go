package main

import (
	"fmt"
	"main/bitflyer"
	"main/config"
	"main/utils"
	"time"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	ticker, _ := apiClient.GetTicker("BTC_USD")
	fmt.Println(ticker.GetMidPrice())
	fmt.Println(ticker.DateTime())
	fmt.Println(ticker.TruncateDateTime(time.Hour))	
}
