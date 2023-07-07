package main

import (
	"fmt"
	"github.com/liujiage/restapi/service"
)

func main() {
	reply := service.Say("jiage liu")
	fmt.Println(reply)
    var userService service.UserService = service.User{"1", ""}
	userService.UserAdd()
	userService.UserDeleteById()
	userService.UserUpdateById()
	userService.UserQueryById()
	 
}
