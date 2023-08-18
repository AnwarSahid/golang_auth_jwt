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
func Login(c *gin.Context) {

	db := database.GetDB()

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.User{}
	password := ""

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "unauthorized",
			"message": " Invalid username, password ",
		})
		return
	}

	comparePass := helpers.ComparePassword([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "unautorize",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"token":  token,
	})
}
