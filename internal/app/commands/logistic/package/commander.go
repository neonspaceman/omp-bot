package logistic_package

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/logistic/logistic_package"
)

type LogisticPackageCommander struct {
	bot            *tgbotapi.BotAPI
	packageService *logistic_package.Service
}

func NewLogisticPackageCommander(
	bot *tgbotapi.BotAPI,
) *LogisticPackageCommander {
	packageService := logistic_package.NewService(1)

	return &LogisticPackageCommander{
		bot:            bot,
		packageService: packageService,
	}
}

func (c *LogisticPackageCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *LogisticPackageCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
