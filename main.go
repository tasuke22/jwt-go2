package main

import (
	"fmt"
	"github.com/tasuke/go-auth/models"
	"github.com/tasuke/go-auth/router"
)

func main() {
	err := models.SetUpDB()
	if err != nil {
		fmt.Println("Cannot connect to database ", "mysql")
	}

	router.Run()
}
