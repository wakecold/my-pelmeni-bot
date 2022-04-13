package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	order := []string{}
	isOrdering := false

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
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "create_order":
			msg.Text = "I am ready to collect your order"
			isOrdering = true
		case "add_item":
			if isOrdering {
				arg := update.Message.CommandArguments()
				from := update.Message.From
				order = append(order, arg)
				msg.Text = fmt.Sprintf("Added %v from %v", arg, from)
			} else {
				msg.Text = "Currently not taking orders. /create_order to start ordering"
			}
		case "remove_item":
			msg.Text = "TODO:"
		case "print_current_order":
			msg.Text = fmt.Sprintf("Current order is %v", order)
		case "finish_order":
			isOrdering = false
			msg.Text = fmt.Sprintf("Thank you for your order. Your order is %v", order)
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
