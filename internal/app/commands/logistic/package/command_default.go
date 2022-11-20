package logistic_package

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticPackageCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Unknown command.\nSend /help__logistic__package for help.")

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackage.Default: error sending reply message to chat - %v", err)
	}
}
