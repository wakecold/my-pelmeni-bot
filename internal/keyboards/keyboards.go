package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var SmallPelmeniKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("бык", "1"),
		tgbotapi.NewInlineKeyboardButtonData("мясные", "2"),
		tgbotapi.NewInlineKeyboardButtonData("сырные", "3"),
		tgbotapi.NewInlineKeyboardButtonData("итальянцы", "4"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("креветка", "5"),
		tgbotapi.NewInlineKeyboardButtonData("бульба", "6"),
		tgbotapi.NewInlineKeyboardButtonData("мамин сибиряк", "7"),
		// tgbotapi.NewInlineKeyboardButtonData("дать заднюю", "8"),
	),
)

var BigPelmeniKeyboard = tgbotapi.NewInlineKeyboardMarkup(

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("бык", "11"),
		tgbotapi.NewInlineKeyboardButtonData("мясные", "12"),
		tgbotapi.NewInlineKeyboardButtonData("сырные", "13"),
		tgbotapi.NewInlineKeyboardButtonData("итальянцы", "14"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("креветка", "15"),
		tgbotapi.NewInlineKeyboardButtonData("бульба", "16"),
		tgbotapi.NewInlineKeyboardButtonData("мамин сибиряк", "17"),
	),
)

var SauceKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("сметана", "21"),
		tgbotapi.NewInlineKeyboardButtonData("чесночный", "22"),
		tgbotapi.NewInlineKeyboardButtonData("мазик", "23"),
		tgbotapi.NewInlineKeyboardButtonData("острая сметана", "24"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("сметана с песто", "25"),
		tgbotapi.NewInlineKeyboardButtonData("сацбели", "26"),
		tgbotapi.NewInlineKeyboardButtonData("томатный с базиликом", "27"),
	),
)

var ResetOrderKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("reset my order", "0"),
		tgbotapi.NewInlineKeyboardButtonData("whats my order", "99"),
	),
)
