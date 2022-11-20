package logistic_package

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticPackageCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Help from Logistic.Packages:\n"+
			"/help__logistic__package — print list of commands\n"+
			"/list__logistic__package — get a list of your entity\n"+
			"/get__logistic__package [package_index] — get a entity - get a entity\n"+
			"/delete__logistic__package [package_index] — delete an existing entity\n"+
			"/new__logistic__package \"[package_name]\" [package_long] [package_width] [package_height] — create a new entity\n"+
			"/edit__logistic__package [package_index] \"[package_name]\" [package_long] [package_width] [package_height] — edit a entity")

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackage.Help: error sending reply message to chat - %v", err)
	}
}
