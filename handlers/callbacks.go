package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram_todo_bot/services"
	"telegram_todo_bot/utils"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd, taskId := utils.GetKeyValue(update.CallbackQuery.Data)
	switch {
	case cmd == "delete_task":
		services.DeleteTaskCallback(bot, update, taskId)
	}
}
