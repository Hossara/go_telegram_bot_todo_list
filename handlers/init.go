package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type ConfigurationType struct {
	NewUpdateOffset int
	Timeout         int
}

var Configuration = &ConfigurationType{
	NewUpdateOffset: 0,
	Timeout:         60,
}

func Init(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(Configuration.NewUpdateOffset)
	u.Timeout = Configuration.Timeout

	updates, _ := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		// Handle what is the message type
		if update.CallbackQuery != nil {
			Callbacks(bot, update)
		} else if update.Message.IsCommand() {
			Commands(bot, update)
		} else {
			Messages(bot, update)
		}
	}
}
