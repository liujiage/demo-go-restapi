package database_test

import (
	"log"
	"fmt"
	"testing"
	"github.com/liujiage/restapi/common"
	"github.com/liujiage/restapi/service/dao"
	"github.com/golobby/container/v3"
	"github.com/stretchr/testify/assert"
)

var instance = container.New()
var dbHelper = common.DBModelHelper{}
func TestDBHelper_connect(t *testing.T) {
	fmt.Println(dbHelper.GetConnect())
}

func TestDBHelper_connect_container(t *testing.T){
    instance.NamedSingleton("db", func() common.DBHelper {
		//inital connect to database
		dbHelper.GetConnect();
		return &dbHelper
	})
	var helper common.DBHelper
    err := instance.NamedResolve(&helper, "db")
	assert.NoError(t, err)
	assert.Equal(t, helper.GetConnect(), dbHelper.GetConnect())
}

func TestDBHelper_insert(t *testing.T) {
	sql := `insert into user(id,name) values(?,?)`
	id := common.GetUUID()
	dbHelper.GetConnect().MustExec(sql, id, "test")
}

func TestDBHelper_delete(t *testing.T) {
	sql := `delete from user where id = ?`
	dbHelper.GetConnect().MustExec(sql, 1)
}

func TestDBHelper_update(t *testing.T){
	sql := `update user set name = ? where id = ?`
	dbHelper.GetConnect().MustExec(sql, "test_update", "2")
}

func TestDBHelper_query(t *testing.T) {
	sql := `select id as Id, name as Name from user limit 10 `
	var users []dao.UserModelDao
	rows, err:= dbHelper.GetConnect().Queryx(sql)
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
	sql := `select id, name from user limit 10 `
	users := make([]dao.UserModelDao,0)
	rows, err:= dbHelper.GetConnect().Queryx(sql)
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