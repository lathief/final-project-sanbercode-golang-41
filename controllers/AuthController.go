package controllers

import (
	"bioskop/database"
	"bioskop/model"
	"bioskop/repository"
	"bioskop/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input model.AuthenticationInput
	var err error
	if err := context.ShouldBindJSON(&input); err != nil {
		fmt.Println("ERROR1")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	cusID, err := repository.MaxcustomerID(database.DbConnection)
	if err != nil {
		fmt.Println("ERROR2")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.CustomerID = cusID + 1
	err = repository.SaveUser(database.DbConnection, user)
	if err != nil {
		fmt.Println("ERROR2")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": "Success"})
}

func Login(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input.Password)
	user, err := repository.FindUserByUsername(database.DbConnection, input.Username)
	fmt.Println(user.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		fmt.Println("okee")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
