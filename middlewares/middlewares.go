package middleware

import (
	"github.com/tasuke/go-auth/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	_, err = utils.ParseToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}
