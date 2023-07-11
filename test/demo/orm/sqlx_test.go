package demo

import (
	"fmt"
	"log"
	"testing"
	"github.com/jmoiron/sqlx"
)


/***
Test sqlx 
Reference: https://pkg.go.dev/github.com/jmoiron/sqlx#section-readme
***/
func TestQuickStartSqlx(t *testing.T) {

    var db *sqlx.DB
	db, err := sqlx.Connect("sqlite3", "./../../resource/testSqlx.db")
    if err != nil {
        log.Fatalln(err)
    }
	fmt.Println(db)

	//test create a table 
	sql := `create table user(
		id varchar(50) primary key, 
		name varchar(500), 
		createTime datetime DEFAULT (datetime('now', 'localtime')),
		updateTime datetime
		)`

	db.MustExec(sql)

	defer db.Close()

}