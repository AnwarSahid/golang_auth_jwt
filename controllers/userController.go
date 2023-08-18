package controllers

import (
	"golang_jwt/database"
	"golang_jwt/helpers"
	"golang_jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJson = "application/json"

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
			"error":   "bad request",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created succesfully",
		"data": gin.H{
			"Name":  User.FullName,
			"Email": User.Email,
		},
	})
}
