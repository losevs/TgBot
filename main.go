package main

import (
	"fmt"
	"strings"
	highlight "tgBot/Highlight"
	qrcode "tgBot/QrCode"
	weather "tgBot/Weather"
	"tgBot/token"

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
			msg.Text = "/light + Text\n/qr + Text\n/weather + City/Country"
		case "light":
			words := strings.Split(update.Message.Text, " ")
			words = append(words[:0], words[1:]...)
			msg.Text = highlight.Light(words)
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
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know this command"))
		}

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, resultText)
		msg.ParseMode = tgbotapi.ModeMarkdown
		//msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err)))
		}
	}
}
