package database_test

import (
	"log"
	"fmt"
	"testing"
	"github.com/liujiage/restapi/common"
	"github.com/liujiage/restapi/service/dao"
	
)

func TestDBHelper_connect(t *testing.T) {
	dbHelper := common.DBHelper{}
	dbHelper.Builder()
	fmt.Println(dbHelper.DB)
	
}

func TestDBHelper_insert(t *testing.T) {
	dbHelper := common.DBHelper{}
	sql := `insert into user(id,name) values(?,?)`
	id := common.GetUUID()
	dbHelper.Builder().DB.MustExec(sql, id, "test")
}

func TestDBHelper_delete(t *testing.T) {
	dbHelper := common.DBHelper{}
	sql := `delete from user where id = ?`
	dbHelper.Builder().DB.MustExec(sql, 1)
}

func TestDBHelper_query(t *testing.T) {
	dbHelper := common.DBHelper{}
	sql := `select id as Id, name as Name from user limit 10 `
	var users []dao.UserModelDao
	rows, err:= dbHelper.Builder().DB.Queryx(sql)
	if err != nil {
		fmt.Println(err)
		return 
	}
	for rows.Next() {
		var id string
		var name string
		rows.Scan(&id,&name)
		fmt.Println(id, name)
		user := dao.UserModelDao{Id: id, Name: name}
		users = append(users, user)
	}
	common.PrintSlice(users)
 
}

func TestDBHelper_query_scan(t *testing.T){
	dbHelper := common.DBHelper{}
	sql := `select id, name from user limit 10 `
	users := make([]dao.UserModelDao,0)
	rows, err:= dbHelper.Builder().DB.Queryx(sql)
	if err != nil {
		fmt.Println(err)
		return 
	}
	for rows.Next() {
		var e dao.UserModelDao
		err = rows.StructScan(&e)
		if err != nil {
			log.Fatalf("Scan error: %s\n", err)
		}
		users = append(users, e)
	}
	common.PrintSlice(users)
}