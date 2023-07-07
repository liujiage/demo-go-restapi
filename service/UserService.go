package service
import(
	"fmt"
)

/***
User CRUD Interface
***/
type UserService interface{
	UserAdd() int
	UserDeleteById() int
	UserUpdateById() int
	UserQueryById()  User
}

/****
User struct
****/
type User struct{
	Id string
	Name string
}

/****
Add a new user
****/
func (user User) UserAdd() int{
	fmt.Println("UserAdd")
	//TODO: save user into database 
	return 0
}

/****
Delete a user by id
****/
func (user User) UserDeleteById() int{
	fmt.Println("UserDeleteById")
	//TODO: delete user by id from database
	return 0
}

/****
Update a user by id
****/
func (user User) UserUpdateById() int{
	fmt.Println("UserUpdateById")
	//TODO: update a user by id from database
	return 0
}

/****
Query a user by id
****/
func (user User) UserQueryById() User{
    fmt.Println("UserQuery")
	//TODO: query user by id from database
	return User{}

}

/****
Only for testing 
****/
func Say(msg string) string {
	return "Hello " + msg
}
