package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"
)

func GenerateToken(userId string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY") // 暗号化、復号化するためのキー
	tokenLifeTime, err := strconv.Atoi(os.Getenv("TOKEN_LIFETIME"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * time.Duration(tokenLifeTime)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
