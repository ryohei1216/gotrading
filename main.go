package main

import (
	"main/app/controllers"
	"main/config"
	"main/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
