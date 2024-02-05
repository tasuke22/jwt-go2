package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuke/go-auth/models"
	"net/http"
)

func (handler *Handler) GetUsers(c *gin.Context) {
	// データベースからすべてのユーザーを取得
	users, err := models.GetAllUsers(handler.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

//func (handler *Handler) GetUser(c *gin.Context) {
//	// The authenticated user's ID is already set on the gin Context in the JWT middleware
//	userId, exists := c.Get("user_id")
//	fmt.Println(userId)
//	if !exists {
//		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve user id from context"})
//		return
//	}
//
//	// The userId is fetched from the database
//	user, err := models.GetUserById(handler.DB, userId.(string))
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve user"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"user": user})
//}
