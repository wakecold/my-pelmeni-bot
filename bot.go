package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	goods := map[string]string{
		"1":  "бык small",
		"2":  "мясные small",
		"3":  "сырные small",
		"4":  "итальянцы small",
		"5":  "креветка small",
		"6":  "бульба small",
		"7":  "мамин сибиряк small",
		"8":  "классика жанра small",
		"11": "бык big",
		"12": "мясные big",
		"13": "сырные big",
		"14": "итальянцы big",
		"15": "креветка big",
		"16": "бульба big",
		"17": "мамин сибиряк big",
		"18": "классика жанра big",
		"21": "сметана",
		"22": "чесночный",
		"23": "мазик",
		"24": "острая сметана",
		"25": "сметана с песто",
		"26": "сацбели",
		"27": "томатный с базиликом",
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	order := make(map[string]map[string]int)

	var smallPelmeniKeyboard = tgbotapi.NewInlineKeyboardMarkup(
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
			tgbotapi.NewInlineKeyboardButtonData("классика жанра", "8"),
		),
	)

	var bigPelmeniKeyboard = tgbotapi.NewInlineKeyboardMarkup(

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
			tgbotapi.NewInlineKeyboardButtonData("классика жанра", "18"),
		),
	)

	var sauceKeyboard = tgbotapi.NewInlineKeyboardMarkup(
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

	// TODO: remove
	bot.Debug = true
	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {

		// Used to work with what user clicked on a keyboard
		if update.CallbackQuery != nil {
			data := update.CallbackQuery.Data
			from := update.CallbackQuery.From.UserName
			// bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))

			if len(order[from]) == 0 {
				order[from] = make(map[string]int)
				order[from][data] = 1
			} else {
				order[from][data] += 1
			}

			bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("Thank you for your order @%v", from)))
		}
		if update.Message == nil {
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "print_current_order":
			msg.Text = fmt.Sprintf("Current order is %v", order)
		case "finish_order":

			orderResult := "Thank you for your order. Your order is \n"
			for user, val := range order {

				orderResult += user + " ordered: "
				for itemId, amount := range val {
					orderResult += strconv.Itoa(amount) + "x " + goods[itemId]
				}

				orderResult += "\n"
			}
			msg.Text = orderResult
		case "create_order":
			msg.Text = "SMALL PELMENI KEYBOARD"
			msg.ReplyMarkup = smallPelmeniKeyboard

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

			msg.Text = "BIG PELMENI KEYBOARD"
			msg.ReplyMarkup = bigPelmeniKeyboard

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

			msg.Text = "SAUCE KEYBOARD"
			msg.ReplyMarkup = sauceKeyboard
		case "yura":
			msg.Text = "pidor"
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

	}

}
