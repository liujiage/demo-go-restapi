package handler_test

import(
	"testing"
	"github.com/liujiage/restapi/handler"
)

func TestContainer_start(t *testing.T){
	//Inject db
	handler.ContainerStart()
	//todo got the db ....
}