package main

import (
	"fmt"
	"strings"

	quote "github.com/losevs/TgBot/Quote"
	weather "github.com/losevs/TgBot/Weather"
	"github.com/losevs/TgBot/currency"
	"github.com/losevs/TgBot/help"
	"github.com/losevs/TgBot/random"
	"github.com/losevs/TgBot/token"

	highlight "github.com/losevs/TgBot/Highlight"
	qrcode "github.com/losevs/TgBot/QrCode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(token.TG_API)
	if err != nil {
		fmt.Println(err)
	}
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "help":
			msg.Text = help.Help()

		// HIGHLIGHT
		case "light":
			words := strings.Split(update.Message.Text, " ")
			words = append(words[:0], words[1:]...)
			msg.Text = highlight.Light(words)

		// QR-CODE
		case "qr":
			words := strings.Split(update.Message.Text, " ")
			words = append(words[:0], words[1:]...)
			if len(words) == 0 {
				msg.Text = "Bad request: message text is empty"
				bot.Send(msg)
				continue
			}
			bot.Send(tgbotapi.NewPhoto(update.Message.Chat.ID, qrcode.MakeNewQR(words)))
			continue

		// WEATHER
		case "weather":
			words := strings.Split(update.Message.Text, " ")

			if len(words) < 2 {
				msg.Text = "Bad request: city text is empty"
				bot.Send(msg)
				continue
			}
			City := words[1]
			response, err := weather.KnowWeather(City)
			if err != nil {
				msg.Text = err.Error()
				bot.Send(msg)
				continue
			}
			msg.Text = response

		// RANDOM
		case "slot":
			bot.Send(tgbotapi.NewDiceWithEmoji(update.Message.Chat.ID, "🎰"))
			continue
		case "dice":
			bot.Send(tgbotapi.NewDice(update.Message.Chat.ID))
			continue
		case "roll":
			words := strings.Split(update.Message.Text, " ")
			if len(words) == 1 {
				num, err := random.Roll("100")
				if err != nil {
					msg.Text = err.Error()
					bot.Send(msg)
					continue
				}
				msg.Text = fmt.Sprint(num)
			} else {
				num, err := random.Roll(words[1])
				if err != nil {
					msg.Text = err.Error()
					bot.Send(msg)
					continue
				}
				msg.Text = fmt.Sprint(num)
			}

		// CURRENCY
		case "allcur":
			msg.Text = currency.AllCurr()
		case "curr":
			words := strings.Split(update.Message.Text, " ")
			if len(words) != 4 {
				msg.Text = "Bad request"
				bot.Send(msg)
				continue
			}
			result, err := currency.Convert(words[1], words[2], words[3])
			if err != nil {
				msg.Text = err.Error()
				bot.Send(msg)
				continue
			}
			msg.Text = result

		// QUOTE
		case "quote":
			resp, err := quote.NewQuote()
			if err != nil {
				msg.Text = err.Error()
				bot.Send(msg)
				continue
			}
			msg.Text = resp
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know this command"))
			continue
		}
		msg.ParseMode = tgbotapi.ModeMarkdown
		//msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err)))
		}
	}
}
