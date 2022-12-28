package database

import (
	"log"

	"github.com/MatheusLasserre/go-fiber-docker-sqloback/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDBConnection() {
	db, err := sqlx.Open("mysql", config.GetEnv("DSN"))
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	log.Println("Successfully connected to database")

	Db = db
}
