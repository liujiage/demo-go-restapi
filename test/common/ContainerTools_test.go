package common_test

import (
	"fmt"
	"os"
	"runtime"
	"path/filepath"
	"testing"
	"github.com/liujiage/restapi/common"
)


func TestTools_GetDBConnect(t *testing.T){
	db := common.GetDBConnect()
	fmt.Println(db)
}
 


func TestGetWorkPath(t *testing.T){
	f, _ := os.Getwd()  
	fmt.Println(filepath.Base(f))
	fmt.Println(filepath.Dir(f))
	fmt.Println(filepath.Clean(f))
	_,b,_,_ := runtime.Caller(0)
    basepath:= filepath.Dir(b)
	fmt.Println(basepath)
	
}