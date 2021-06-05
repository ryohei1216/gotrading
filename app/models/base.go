package models

import (
	"database/sql"
	"fmt"
	"log"
	"main/config"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const(
	tableNameSignalEvents = "signal_events"
)

var DbConnection *sql.DB

func GetCanleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}

func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmd := fmt.Sprintf(`
				CREATE TABLE IF NOT EXISTS %s (
				time DATETIME PRIMARY KEY NOT NULL,
				product_code STRING,
				side STRING,
				price FLOAT,
				size FLOAT)`, tableNameSignalEvents)
	DbConnection.Exec(cmd)

	for _, duration := range config.Config.Durations {
		tableName := GetCanleTableName(config.Config.ProductCode, duration)
		c := fmt.Sprintf(`
				CREATE TABLE IF NOT EXISTS %s (
				time DATETIME PRIMARY KEY NOT NULL,
				open FLOAT,
				close FLOAT,
				high FLOAT,
				low open FLOAT,
				volume FLOAT)`, tableName)
		DbConnection.Exec(c)
	} 
}