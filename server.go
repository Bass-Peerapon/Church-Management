package main

import (
	"github.com/Church-Management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userGroup := r.Group("users") 
	{
		userGroup.POST("register", routes.UsersRegister)
	}
	r.Run(":3000")
}
