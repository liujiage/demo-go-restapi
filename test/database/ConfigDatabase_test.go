package database_test

import (
	"fmt"
	"testing"
	"github.com/liujiage/restapi/database"
)

func TestStart_test(t *testing.T) {
	//dbHelper := database.GetDBHelper()
	//fmt.Println(dbHelper.DB)
	dbHelper2 := database.DBHelper{}
	dbHelper2.Builder()
	fmt.Println(dbHelper2.DB)
	
}
