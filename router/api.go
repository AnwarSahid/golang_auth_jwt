package router

import (
	"golang_jwt/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.Login)
	}
	// middlewareRouter := router.Group("/testing")
	// {
	// 	middlewareRouter.Use(middleware.Authorization())
	// 	middlewareRouter.GET("/test", controllers.VerifyToken)
	// }

	return router
}
