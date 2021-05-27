package controllers

import (
	"net/http"

	responses "github.com/Church-Management/api/response"
	"github.com/gin-gonic/gin"
)

func (server *Server) Home(c *gin.Context) {
	responses.JSON(c.Writer, http.StatusOK, "Welcome To This Awesome API")
}
