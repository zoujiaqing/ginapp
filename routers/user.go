package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zoujiaqing/ginapp/controllers"
)

func LoadUserRouters(r *gin.Engine) {
	userRouter := r.Group("user")
	{
		userRouter.GET("/register", controllers.Register)
		userRouter.GET("/login", controllers.Login)
	}
}
