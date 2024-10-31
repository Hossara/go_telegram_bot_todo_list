package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram_todo_bot/services"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Define commands
	switch update.Message.Command() {
	case "start":
		services.Start(bot, update)
	case "set_todo":
		services.SetTask(bot, update)
	case "delete_todo":
		services.DeleteTask(bot, update)
	case "show_all_todos":
		services.ShowAllTasks(bot, update)
	}
}
