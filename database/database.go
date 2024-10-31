package database

import (
	"fmt"
	"log"
	"telegram_todo_bot/config"
	"telegram_todo_bot/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// Define database connection info
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s, dbname=%s",
		config.Config("POSTGRES_HOST"),
		config.Config("POSTGRES_USER"),
		config.Config("POSTGRES_PASSWORD"),
		config.Config("POSTGRES_PORT"),
		"postgres")

	// Connect to database
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Remove old database
	//if err := DB.Exec("DROP DATABASE IF EXISTS todolist;").Error; err != nil {
	//	panic(err)
	//}

	// Create database
	/*if err := DB.Exec("CREATE DATABASE todolist").Error; err != nil {
		log.Fatalf("error executing query %v", err)
	}*/

	// Migrate tables
	DB.AutoMigrate(&models.Task{})

	return DB
}
