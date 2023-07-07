package handler
import (
	"github.com/liujiage/restapi/config"
	"github.com/gin-gonic/gin"
)


func Start(){
	router := gin.Default()
	router.POST(config.ADD_USER, UserAdd)
	router.DELETE("/user/delete/:id", UserDeleteById)
	router.PUT("/user/update", UserUpdateById)
	router.GET("/user/:id", UserQueryById)

	router.Run("localhost:8080")
}