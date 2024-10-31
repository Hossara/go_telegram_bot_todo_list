package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func CmdKeyboard() tgbotapi.ReplyKeyboardMarkup {
	// Define keyboard commands
	var cmdKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/set_todo"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/delete_todo"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/show_all_todos"),
		),
	)
	return cmdKeyboard
}
