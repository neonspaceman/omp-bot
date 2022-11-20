package logistic_package

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *LogisticPackageCommander) List(inputMessage *tgbotapi.Message) {
	textMsg := ""

	list, hasNext := c.packageService.List(0)

	if len(list) == 0 {
		textMsg = "Packages not found"
	} else {
		var sb strings.Builder
		for key, val := range list {
			sb.WriteString(fmt.Sprintf("Id: %d. %s\n", key, val.String()))
		}
		textMsg = sb.String()
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, textMsg)

	if hasNext {
		serializedData, _ := json.Marshal(&CallbackListData{
			Offset: c.packageService.CountOfRow,
		})

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Next ▶️", path.CallbackPath{
				Domain:       "logistic",
				Subdomain:    "package",
				CallbackName: "list",
				CallbackData: string(serializedData),
			}.String())))
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackage.List: error sending reply message to chat - %v", err)
	}
}
