package logistic_package

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *LogisticPackageCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	callbackData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &callbackData)
	if err != nil {
		log.Printf("LogisticPackage.CallbackList: error json.Unmarshal  - %v", err)
		return
	}

	textMsg := ""

	list, hasNext := c.packageService.List(callbackData.Offset)

	if len(list) == 0 {
		textMsg = "Packages not found"
	} else {
		var sb strings.Builder
		for key, val := range list {
			sb.WriteString(fmt.Sprintf("Id: %d. %s\n", callbackData.Offset+key, val.String()))
		}
		textMsg = sb.String()
	}

	msg := tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, textMsg)

	var prevNextInlineKeyboard []tgbotapi.InlineKeyboardButton

	if callbackData.Offset > 0 {
		serializedData, _ := json.Marshal(&CallbackListData{
			Offset: callbackData.Offset - c.packageService.CountOfRow,
		})

		prevNextInlineKeyboard = append(prevNextInlineKeyboard, tgbotapi.NewInlineKeyboardButtonData("◀️️ Prev", path.CallbackPath{
			Domain:       "logistic",
			Subdomain:    "package",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}.String()))
	}
	if hasNext {
		serializedData, _ := json.Marshal(&CallbackListData{
			Offset: callbackData.Offset + c.packageService.CountOfRow,
		})

		prevNextInlineKeyboard = append(prevNextInlineKeyboard, tgbotapi.NewInlineKeyboardButtonData("Next ▶️", path.CallbackPath{
			Domain:       "logistic",
			Subdomain:    "package",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}.String()))
	}

	if prevNextInlineKeyboard != nil {
		markup := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(prevNextInlineKeyboard...))
		msg.ReplyMarkup = &markup
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackage.CallbackList: error sending reply message to chat - %v", err)
	}
}
