package logistic_package

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/shlex"
	"github.com/ozonmp/omp-bot/internal/service/logistic/logistic_package"
)

func (c *LogisticPackageCommander) doGet(inputMessage *tgbotapi.Message) *tgbotapi.MessageConfig {
	args, err := shlex.Split(inputMessage.Text)
	if err != nil {
		log.Printf("LogisticPackage.Get: error parsing message - %v", err)
		return nil
	}

	if len(args) != 2 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Incorrect command format.\n"+
			"/get__logistic__package [package_index]")
		return &msg
	}

	idx, err := strconv.Atoi(args[1])
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Index must be a number")
		return &msg
	}

	p, err := c.packageService.Get(idx)
	if err == logistic_package.ErrIndexOutOfRange {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Index out of range")
		return &msg
	} else if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Internal error")
		return &msg
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, p.String())
	return &msg
}

func (c *LogisticPackageCommander) Get(inputMessage *tgbotapi.Message) {
	if msg := c.doGet(inputMessage); msg != nil {
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("LogisticPackage.Get: error sending reply message to chat - %v", err)
		}
	}
}
