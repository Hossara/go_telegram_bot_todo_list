package repositories

import (
	"gorm.io/gorm"
	"telegram_todo_bot/database"
)

var DB *gorm.DB = database.Init()
