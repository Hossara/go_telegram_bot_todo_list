package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram_todo_bot/services"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	services.SetTaskCallback(bot, update)
}
