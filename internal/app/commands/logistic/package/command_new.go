package logistic_package

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/shlex"
	"github.com/ozonmp/omp-bot/internal/service/logistic/logistic_package"
)

func (c *LogisticPackageCommander) doNew(inputMessage *tgbotapi.Message) *tgbotapi.MessageConfig {
	args, err := shlex.Split(inputMessage.Text)
	if err != nil {
		log.Printf("LogisticPackage.New: error parsing message - %v", err)
		return nil
	}

	if len(args) != 5 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Incorrect command format.\n"+
			"/new__logistic__package \"[package_name]\" [package_long] [package_width] [package_height]")
		return &msg
	}

	var errors []string
	title := args[1]
	if len(title) == 0 {
		errors = append(errors, "Title must be filled in")
	}

	long, err := strconv.Atoi(args[2])
	if err != nil {
		errors = append(errors, "Long must be a number")
	} else if long <= 0 {
		errors = append(errors, "Long must be greater 0")
	}

	width, err := strconv.Atoi(args[3])
	if err != nil {
		errors = append(errors, "Width must be a number")
	} else if width <= 0 {
		errors = append(errors, "Width must be greater 0")
	}

	height, err := strconv.Atoi(args[4])
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
	c.packageService.New(&p)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "New package has been added.\n"+p.String())
	return &msg
}

func (c *LogisticPackageCommander) New(inputMessage *tgbotapi.Message) {
	if msg := c.doNew(inputMessage); msg != nil {
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("LogisticPackage.New: error sending reply message to chat - %v", err)
		}
	}
}
