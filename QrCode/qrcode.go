package qrcode

import (
	"image/color"

	"os"
	"strings"

	"github.com/skip2/go-qrcode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MakeNewQR(text []string) tgbotapi.FileBytes {
	textStr := strings.Join(text, " ")
	qrcode.WriteColorFile(textStr, qrcode.Medium, 512, color.CMYK{0, 35, 100, 0}, color.Black, "qrcode.png")
	photoBytes, _ := os.ReadFile("qrcode.png")
	return tgbotapi.FileBytes{
		Name:  "qrcode",
		Bytes: photoBytes,
	}
}
