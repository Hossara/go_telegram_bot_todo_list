package services

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram_todo_bot/keyboards"
	"telegram_todo_bot/repositories"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Greet
	text := "Hi, here you can create todos for your todolist."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	// Show keyboard menu
	msg.ReplyMarkup = keyboards.CmdKeyboard()

	// Send welcome message
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Ask for task
	text := "Please, write todo."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Todo successfully created"

	err := repositories.SetTask(update)
	if err != nil {
		text = "Couldn't set task"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func DeleteTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Prompt user which task user wants to delete
	data, _ := repositories.GetAllTasks(update.Message.Chat.ID)

	// Show each task as 2*X grid filled by buttons with (text=task_name, content=task_id+action) settings
	var btns []tgbotapi.InlineKeyboardButton

	for i := 0; i < len(data); i++ {
		btn := tgbotapi.NewInlineKeyboardButtonData(data[i].Task, "delete_task="+data[i].ID.String())
		btns = append(btns, btn)
	}

	var rows [][]tgbotapi.InlineKeyboardButton

	for i := 0; i < len(btns); i += 2 {
		if i < len(btns) && i+1 < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i], btns[i+1])
			rows = append(rows, row)
		} else if i < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i])
			rows = append(rows, row)
		}
	}

	fmt.Println(len(rows))
	var keyboard = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
	//keyboard.InlineKeyboard = rows

	// Show keyboard and send message
	text := "Please, select todo you want to delete"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func DeleteTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, taskId string) {
	text := "Task successfully deleted"

	err := repositories.DeleteTask(taskId)
	if err != nil {
		text = "Couldn't delete task"
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func ShowAllTasks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Tasks: \n"

	tasks, err := repositories.GetAllTasks(update.Message.Chat.ID)
	if err != nil {
		text = "Couldn't get tasks"
	}

	for i := 0; i < len(tasks); i++ {
		text += tasks[i].Task + " \n"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
