package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func UserAdd(c *gin.Context) {
	fmt.Println("add user handler")
}

func UserDeleteById(c *gin.Context) {
	fmt.Println("delete user handler")
}

func UserUpdateById(c *gin.Context){
	fmt.Println("update user handler")
}

func UserQueryById(c *gin.Context){
	fmt.Println("query user handler")
}

