package dao

import (
	"fmt"
)

/***
User CRUD Interface. 
Dao level
***/
type UserDao interface{
	UserAdd() int
	UserDeleteById() int
	UserUpdateById() int
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
func (user *UserModelDao) UserAdd() int{
	fmt.Println("UserAdd from dao")
	//TODO: save user into database 
	return 0
}

/****
Delete a user by id
****/
func (user *UserModelDao) UserDeleteById() int{
	fmt.Println("UserDeleteById")
	//TODO: delete user by id from database
	return 0
}

/****
Update a user by id
****/
func (user *UserModelDao) UserUpdateById() int{
	fmt.Println("UserUpdateById")
	//TODO: update a user by id from database
	return 0
}

/****
Query a user by id
****/
func (user *UserModelDao) UserQueryById() UserModelDao{
    fmt.Println("UserQuery")
	//TODO: query user by id from database
	return UserModelDao{}

}