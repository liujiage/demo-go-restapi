package service

import (
	"fmt"
	"github.com/liujiage/restapi/service/dao"
	"github.com/liujiage/restapi/common"
)

/***
User CRUD Interface. 
Service level
***/
type UserService interface{
	UserAdd() string
	UserMutiManage() string
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
func (user UserModelService) UserAdd() string{
	fmt.Println("UserAdd")
	user.User.UserAdd()
	return common.SUCCESS
}

/****
Simulate Service level call more than one Dao services
****/
func (user UserModelService) UserMutiManage() string{
	fmt.Println("UserMutiManage")
	//1. query user by id 
    user.User.UserQueryById()
	//2. update user by id 
	user.User.UserUpdateById()
	//3. delete the user by id 
	user.User.UserDeleteById()
	return common.SUCCESS
}


/****
Only for testing 
****/
func Say(msg string) string {
	return "Hello " + msg
}
