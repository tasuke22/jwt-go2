package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuke/go-auth/controllers"
)

func addAuthRouter(rg *gin.RouterGroup, h *controllers.Handler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/signup", h.SignUpHandler)
		auth.POST("/login", h.LoginHandler)
	}
}
