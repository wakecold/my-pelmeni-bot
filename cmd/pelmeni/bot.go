package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/wakecold/my-pelmeni-bot/internal/constants"
	"github.com/wakecold/my-pelmeni-bot/internal/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
var bot *tgbotapi.BotAPI

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}
	order := make(map[string]map[string]int)


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
			onUserClick(data, from, update)
		}
		if update.Message == nil {
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case constants.Print:
			msg.Text = fmt.Sprintf("Current order is %v", order)
		case constants.Finish:

			orderResult := "Thank you for your order. Your order is \n"
			for user, val := range order {

				orderResult += user + " ordered: "
				for itemId, amount := range val {
					orderResult += strconv.Itoa(amount) + "x " + constants.Goods[itemId]
				}

				orderResult += "\n"
			}
			msg.Text = orderResult
		case constants.Create:
			msg.Text = "SMALL PELMENI KEYBOARD"
			msg.ReplyMarkup = keyboards.SmallPelmeniKeyboard

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

			msg.Text = "BIG PELMENI KEYBOARD"
			msg.ReplyMarkup = keyboards.BigPelmeniKeyboard

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

			msg.Text = "SAUCE KEYBOARD"
			msg.ReplyMarkup = keyboards.SauceKeyboard
		case constants.Yura:
			msg.Text = "pidor"
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

	}

}

func onUserClick(data string, from string, update tgbotapi.Update){
	userID := update.CallbackQuery.Message.Chat.ID
	msg := fmt.Sprintf("Thank you forr your order %s", from)
	bot.Send(tgbotapi.NewMessage(userID, msg))
}
