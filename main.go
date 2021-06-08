package main

import (
	"log"
	"main/app/controllers"
	"main/config"
	"main/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}

