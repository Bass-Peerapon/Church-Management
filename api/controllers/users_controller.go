package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Church-Management/api/models"
	responses "github.com/Church-Management/api/response"
	"github.com/Church-Management/api/utils/formaterror"
	"github.com/gin-gonic/gin"
)

func (server *Server) CreateUser(c *gin.Context){
	body , err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(c.Writer, http.StatusInternalServerError, formattedError)
		return
	}
	c.JSON(http.StatusCreated , userCreated)
}
