package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"grocery/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	gin.SetMode(viper.GetString("mode"))
	user := r.Group("/api/user")
	{
		user.GET("/list", UserBaseHandler.UserListHandler)
		user.GET("/:id", UserBaseHandler.GetUserInfo)
		user.POST("/", UserBaseHandler.AddUserHandler)
		user.PUT("/", UserBaseHandler.EditUserHandler)
		user.DELETE("/", UserBaseHandler.DeleteUserHandler)

	}
	port:=viper.GetString("port")
	r.Run(port)
}
