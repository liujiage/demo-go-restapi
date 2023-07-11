package database_test

import (
	"fmt"
	"testing"
	"github.com/liujiage/restapi/common"
)

func TestDBHelper_test(t *testing.T) {
	//dbHelper := database.GetDBHelper()
	//fmt.Println(dbHelper.DB)
	dbHelper2 := common.DBHelper{}
	dbHelper2.Builder()
	fmt.Println(dbHelper2.DB)
	
}
