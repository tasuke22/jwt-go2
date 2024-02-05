package controllers

import (
	"github.com/tasuke/go-auth/models"
	"github.com/tasuke/go-auth/pkg/utils"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *gorm.DB
}

func (handler *Handler) SignUpHandler(context *gin.Context) {
	var signUpInput models.SignUpInput
	err := context.ShouldBind(&signUpInput)
	if err != nil {
		// 本来ログ等でerrは出力した方がよいが今回は省略
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	newUser := &models.User{
		Name:     signUpInput.Name,
		Email:    signUpInput.Email,
		Password: signUpInput.Password,
	}

	err = newUser.Validate()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := newUser.Create(handler.DB)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	token, err := utils.GenerateToken(newUser.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to sign up",
		})
		return
	}

	context.SetCookie(
		"token",
		token,
		3600,
		"/",
		"localhost",
		false,
		true,
	)

	context.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
		"message": "Successfully created user",
	})
}

func (handler *Handler) LoginHandler(context *gin.Context) {
	var loginInput models.LoginInput
	if err := context.ShouldBind(&loginInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	user, err := models.FindUserByEmail(handler.DB, loginInput.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to find user",
		})
		return
	}

	if !user.VerifyPassword(loginInput.Password) {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password is invalid",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
	})
}
