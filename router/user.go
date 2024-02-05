package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuke/go-auth/controllers"
	middleware "github.com/tasuke/go-auth/middlewares"
)

type UserRouter struct {
	Handler *controllers.Handler
}

func userRouter(rg *gin.RouterGroup, h *controllers.Handler) {
	user := rg.Group("/user")
	user.Use(middleware.AuthMiddleware)
	{
		user.GET("/all", h.GetUsers)
		//user.GET("/:id", h.GetUser)
	}
}
