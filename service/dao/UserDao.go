package dao

import (
	"fmt"
	"github.com/liujiage/restapi/common"
)

/***
User CRUD Interface. 
Dao level
***/
type UserDao interface{
	UserAdd() string
	UserDeleteById() string
	UserUpdateById() string
	UserQueryById()  UserModelDao
}

/****
User struct. 
Dao level
****/
type UserModelDao struct{
    Id string   `db:"id"`
	Name string `db:"name"`
}


/****
Add a new user
****/
func (user *UserModelDao) UserAdd() string{
	fmt.Println("UserAdd from dao")
	sql := `insert into user(id,name) values(?,?)`
	//id := common.GetUUID()
    common.GetDBConnect().MustExec(sql, user.Id, user.Name) 
	return common.SUCCESS
}

/****
Delete a user by id
****/
func (user *UserModelDao) UserDeleteById() string{
	fmt.Println("UserDeleteById")
	sql := `delete from user where id=?`
	common.GetDBConnect().MustExec(sql,user.Id)
	return common.SUCCESS
}

/****
Update a user by id
****/
func (user *UserModelDao) UserUpdateById() string{
	fmt.Println("UserUpdateById")
	sql := "update user set name=? where id=?"
	common.GetDBConnect().MustExec(sql,user.Name, user.Id)
	return common.SUCCESS
}

/****
Query a user by id
****/
func (user *UserModelDao) UserQueryById() UserModelDao{
    fmt.Println("UserQuery")
	//TODO: query user by id from database
	return UserModelDao{}

}