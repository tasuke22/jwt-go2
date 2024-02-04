package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func SetUpDB() *gorm.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := "localhost" // または Docker Compose で定義されたサービス名
	dbPort := "3306"      // デフォルトのMySQLポート
	dbName := os.Getenv("MYSQL_DATABASE")

	// データベースへの接続文字列を構築
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ", "mysql")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", "mysql")
	}

	return db
}
