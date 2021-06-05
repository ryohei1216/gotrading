package main

import (
	"fmt"
	"main/app/models"
	"main/config"
	"main/utils"
)



func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
}
