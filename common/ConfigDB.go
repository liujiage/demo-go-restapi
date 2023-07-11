package common

import (
	"github.com/jmoiron/sqlx"
	 _ "github.com/mattn/go-sqlite3"
	"log"
)


type DBHelper struct {
	DB *sqlx.DB
}

func (dbHelper *DBHelper) Builder() *DBHelper {
	dbConnect, err := sqlx.Connect("sqlite3", "./../../resource/app.db")
	if err != nil {
		log.Fatalln(err)
	}
	dbHelper.DB = dbConnect
	return dbHelper
}
