package router

import (
	"golang_jwt/controllers"
	"golang_jwt/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.Login)
	}

	productRouter := router.Group("/products")
	{
		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", middleware.ProductAuthorization(), controllers.UpdateProduct)
	}
	return router
}
