package service

import (
	"fmt"
	"github.com/liujiage/restapi/service/dao"
)

/***
User CRUD Interface. 
Service level
***/
type UserService interface{
	UserAdd() int
	UserMutiManage() int
}

/****
User struct. 
Service level
****/
type UserModelService struct{
	User dao.UserModelDao
}

/****
Add a new user
****/
func (user UserModelService) UserAdd() int{
	fmt.Println("UserAdd")
	user.User.UserAdd()
	return 0
}

/****
Simulate Service level call more than one Dao services
****/
func (user UserModelService) UserMutiManage() int{
	fmt.Println("UserMutiManage")
	//1. query user by id 
    user.User.UserQueryById()
	//2. update user by id 
	user.User.UserUpdateById()
	//3. delete the user by id 
	user.User.UserDeleteById()
	return 0
}


/****
Only for testing 
****/
func Say(msg string) string {
	return "Hello " + msg
}
