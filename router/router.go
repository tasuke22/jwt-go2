package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuke/go-auth/controllers"
	"github.com/tasuke/go-auth/models"
)

func Run() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	handler := &controllers.Handler{
		DB: models.SetUpDB(),
	}

	api := router.Group("/api")
	v1 := api.Group("/v1")
	authRouter(v1, handler)
	userRouter(v1, handler)

	return router
}
