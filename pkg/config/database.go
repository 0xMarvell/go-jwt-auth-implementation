package config

import (
	"fmt"
	"log"
	"os"

	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/models"
	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var dbConnectErr error

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		user,
		password,
		dbname,
		port)

	DB, dbConnectErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.CheckErr("error connecting to database: ", dbConnectErr)
	log.Println("Database Connection Successful 🚀")
}

// RunMigrations runs migrations for the database.
func RunMigrations() {
	migrationErr := DB.AutoMigrate(&models.User{})
	utils.CheckErr("Migration failed: ", migrationErr)
	log.Println("Migration Successful 🚀")
	fmt.Println()
}
