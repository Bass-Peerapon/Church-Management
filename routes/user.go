package routes

import (
	"net/http"
	"time"

	"github.com/Church-Management/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UsersRegister(c *gin.Context) {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		panic(err)
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	newUser := models.User{
		Name:         user.Name,
		SurName:      user.SurName,
		Birthday:     time.Time{},
		Email:        user.Email,
		PasswordHash: string(passwordHash),
		CreateTime:   time.Time{},
		UnixTime:     time.Time{},
	}
	models.DB.Create(&newUser)
	c.JSON(http.StatusOK, newUser)
}
