package common

import (
	"github.com/jmoiron/sqlx"
	 _ "github.com/mattn/go-sqlite3"
	"log"
)

/****
DBHelper interface, manipulate database table
****/
type DBHelper interface {
	GetConnect() *sqlx.DB
}

/****
DBModelHelper as a DBHelper instance. 
share the instance in container (IoC) by DBHelper interface
****/
type DBModelHelper struct{
	DB *sqlx.DB
}


/****
GetConnect return current connect
Public function 
****/
func (model *DBModelHelper) GetConnect() *sqlx.DB {
	if model.DB == nil{
		model.connect()
	}
	return model.DB
}

/****
Connect connect database
Private function
****/
func  (model *DBModelHelper) connect() *sqlx.DB {
	//todo change a const project root path
	db, err := sqlx.Connect("sqlite3", "./../../resource/app.db")
	if err != nil {
		log.Fatalln(err)
	}
	model.DB = db
	return db
}