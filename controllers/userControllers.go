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

func GetAllUser(c *gin.Context) {
	var (
		result gin.H
	)
	user, err := repository.GetAllUser(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"user": user,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetAllCustomer(c *gin.Context) {
	var (
		result gin.H
	)
	customer, err := repository.GetAllCustomer(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"user": customer,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetAllUserCustomer(c *gin.Context) {
	var (
		result            gin.H
		userCustomer      model.UserCustomer
		userCustomerSlice []model.UserCustomer
	)
	user, err := repository.GetAllUser(database.DbConnection)
	for i := 0; i < len(user); i++ {
		customer, _ := repository.FindCustomerById(database.DbConnection, user[i].CustomerID)
		userCustomer.CustomerID = customer.CustomerID
		userCustomer.Email = customer.Email
		userCustomer.MobileNumber = customer.MobileNumber
		userCustomer.Username = user[i].Username
		userCustomer.ID = user[i].ID
		userCustomerSlice = append(userCustomerSlice, userCustomer)
	}

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"user": userCustomerSlice,
		}
	}
	c.JSON(http.StatusOK, result)
}

func UpdateProfile(c *gin.Context) {
	var cusTemp model.Customer
	var cus model.Customer
	err := c.ShouldBindJSON(&cus)
	user, err := utils.CurrentUser(c)
	if err != nil {
		panic(err)
	}
	cusTemp, err = repository.FindCustomerById(database.DbConnection, user.CustomerID)
	fmt.Println(cusTemp.CustomerID)
	if cusTemp.CustomerID != 0 {
		fmt.Println("Update")
		if cusTemp.Email != "" {
			cus.Email = cusTemp.Email
		}
		if cusTemp.MobileNumber != "" {
			cus.MobileNumber = cusTemp.MobileNumber
		}
		fmt.Println(cusTemp.Email)
		fmt.Println(cusTemp.MobileNumber)
		err = repository.UpdateProfile(database.DbConnection, cus, cusTemp.CustomerID)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": "Success Update CustomerID " + c.Param("id"),
		})
	} else {
		fmt.Println("Save")
		fmt.Println(cus.Email)
		fmt.Println(cus.MobileNumber)
		max, err := repository.MaxcustomerID(database.DbConnection)
		cus.CustomerID = max + 1
		err = repository.SaveProfile(database.DbConnection, cus, cus.CustomerID)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": "Success Save CustomerID " + c.Param("id"),
		})
	}

}
