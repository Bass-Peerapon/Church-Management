package main

import (
	"github.com/Church-Management/models"
	"github.com/Church-Management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()
	userGroup := r.Group("users") 
	{
		userGroup.POST("register", routes.UsersRegister)
	}
	r.Run(":3000")
}
