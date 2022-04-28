package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wakecold/my-pelmeni-bot/internal/constants"
	"github.com/wakecold/my-pelmeni-bot/internal/keyboards"
)

// {
//	userid: {
// 			itemid: amount
// 		}
// }
var todaysOrder = make(map[string]map[int]int)
var orderCreator int64
var isOrdering = false

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}
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
		// This block is for work with what user clicked on a keyboard
		if update.CallbackQuery != nil {
			data := update.CallbackQuery.Data
			from := update.CallbackQuery.From.UserName
			chatID := update.CallbackQuery.Message.Chat.ID

			itemId, _ := strconv.Atoi(data)
			// if id = 99 user asked for his order
			if itemId != 99 {
				onUserClick(bot, data, from, chatID)
			}

			replyMessage := "Thank you! Your order is: \n"

			if itemId != 0 && len(todaysOrder[from]) != 0 {
				for item, amount := range todaysOrder[from] {
					replyMessage += strconv.Itoa(amount) + "x " + constants.Goods[item] + " \n"
				}
			} else {
				replyMessage += "empty"
			}

			callbackReply := tgbotapi.NewCallbackWithAlert(update.CallbackQuery.ID, replyMessage)
			bot.Send(callbackReply)

		}
		if update.Message == nil {
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case constants.Print:
			// collect all items and output items and count
			itemsAndCount := make(map[int]int)
			for _, items := range todaysOrder {
				for itemId, amount := range items {
					itemsAndCount[itemId] += amount
				}
			}

			if len(itemsAndCount) == 0 {
				msg.Text = "Current order is empty"
			} else {
				countResult := "Current order is \n"
				for itemId, amount := range itemsAndCount {
					countResult += strconv.Itoa(amount) + "x " + constants.Goods[itemId]
					countResult += "\n"
				}
				msg.Text = countResult
			}
		case constants.Finish:
			if update.Message.From.ID != orderCreator {
				msg.Text = "Sorry, you are not order starter"
			} else {
				isOrdering = false
				orderResult := "Thank you for your order. Your order is \n"
				for user, val := range todaysOrder {

					orderResult += user + " ordered: "
					for itemId, amount := range val {
						orderResult += strconv.Itoa(amount) + "x " + constants.Goods[itemId]
					}

					orderResult += "\n"
				}
				msg.Text = orderResult
				// TODO: clear todaysOrder
			}
		case constants.Create:
			if isOrdering {
				msg.Text = "Sorry, I am already taking order"
			} else {
				isOrdering = true
				orderCreator = update.Message.From.ID
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

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}

				msg.Text = "USEFUL STUFF"
				msg.ReplyMarkup = keyboards.ResetOrderKeyboard
			}
		case constants.Yura:
			msg.Text = "pidor"
		case constants.Help:
			msg.Text = `Пельмень бот
			1. Один человек может выбрать одни большие одни маленькие и одну сметану
			2. Салаты пока отдельно
			3. Если передумали по какой то из позиций  жмите "Дать заднюю"
			4. Один человек создает заказ и один его закрывает
			5. Бот агрегирует заказ`
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

	}

}

func onUserClick(bot *tgbotapi.BotAPI, data string, from string, chatID int64) {
	itemId, _ := strconv.Atoi(data)
	if itemId == 0 {
		// empty users order
		todaysOrder[from] = make(map[int]int)
	} else {
		// check if user ordered before
		if _, ok := todaysOrder[from]; ok {
			todaysOrder[from][itemId] += 1
		} else {
			todaysOrder[from] = make(map[int]int)
			todaysOrder[from][itemId] = 1
		}
	}
	fmt.Printf("----------%v\n", todaysOrder)
}
