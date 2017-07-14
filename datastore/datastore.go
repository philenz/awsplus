package datastore

import (
	"database/sql"
	"fmt"
	"log"
	"sgmanager/config"
)

var Db *sql.DB

func init() {

	cfg := config.Config

	dataSourceName := fmt.Sprintf("dbname=%s user=%s sslmode=disable password=%s",
				cfg.DBName, cfg.DBUser, cfg.DBPass)

	var err error
	Db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return
}

