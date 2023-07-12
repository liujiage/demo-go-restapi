package handler

import(
	"github.com/golobby/container/v3"
	"github.com/liujiage/restapi/common"
)


func ContainerStart(){
	//init and inject it as a singleton 
	dbHelper := common.DBHelper{}
	dbHelper.Builder()
	container.Fill(&dbHelper)
}