package logistic_package

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/shlex"
	"github.com/ozonmp/omp-bot/internal/service/logistic/logistic_package"
)

func (c *LogisticPackageCommander) doEdit(inputMessage *tgbotapi.Message) *tgbotapi.MessageConfig {
	args, err := shlex.Split(inputMessage.Text)
	if err != nil {
		log.Printf("LogisticPackage.Edit: error parsing message - %v", err)
		return nil
	}

	if len(args) != 6 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Incorrect command format.\n"+
			"/edit__logistic__package [package_idx] \"[package_name]\" [package_long] [package_width] [package_height]")
		return &msg
	}

	idx, err := strconv.Atoi(args[1])
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Index must be a number")
		return &msg
	}

	var errors []string
	title := args[2]
	if len(title) == 0 {
		errors = append(errors, "Title must be filled in")
	}

	long, err := strconv.Atoi(args[3])
	if err != nil {
		errors = append(errors, "Long must be a number")
	} else if long <= 0 {
		errors = append(errors, "Long must be greater 0")
	}

	width, err := strconv.Atoi(args[4])
	if err != nil {
		errors = append(errors, "Width must be a number")
	} else if width <= 0 {
		errors = append(errors, "Width must be greater 0")
	}

	height, err := strconv.Atoi(args[5])
	if err != nil {
		errors = append(errors, "Height must be a number")
	} else if height <= 0 {
		errors = append(errors, "Height must be greater 0")
	}

	if len(errors) != 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, strings.Join(errors, "\n"))
		return &msg
	}

	p := logistic_package.Package{
		Title: title,
		Sizes: logistic_package.Sizes{
			Width:  width,
			Height: height,
			Long:   long,
		},
	}

	err = c.packageService.Edit(idx, &p)
	if err == logistic_package.ErrIndexOutOfRange {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Index out of range")
		return &msg
	} else if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Internal error")
		return &msg
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Package info has been updated")
	return &msg
}

func (c *LogisticPackageCommander) Edit(inputMessage *tgbotapi.Message) {
	if msg := c.doEdit(inputMessage); msg != nil {
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("LogisticPackage.Edit: error sending reply message to chat - %v", err)
		}
	}
}
