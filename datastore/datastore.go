package datastore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
	"awsplus/config"
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

	// create / populate tables if not already done...
	createSecurityGroupSQL := fmt.Sprintf(
		"CREATE TABLE security_group (" +
			"id varchar(20) CONSTRAINT sg_pkey PRIMARY KEY," +
			"vpc_id varchar(20) NOT NULL," +
			"description varchar(300)" +
			");")
	Db.Exec(createSecurityGroupSQL)

	createIngressSQL := fmt.Sprintf(
		"CREATE TABLE ingress (" +
			"sg_id varchar(20) REFERENCES security_group(id)," +
			"protocol varchar(10) NOT NULL," +
			"from_port integer NOT NULL," +
			"to_port integer NOT NULL," +
			"source varchar(40) NOT NULL," +
			"UNIQUE (sg_id, protocol, from_port, source)" +
			");")
	Db.Exec(createIngressSQL)

	return
}

