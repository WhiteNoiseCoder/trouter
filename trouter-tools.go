package trouter

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type TErrorHandler = func(bot *tgbotapi.BotAPI, update *tgbotapi.Update, err error)
type TStandardHandler = func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (err error)

func StandardErrorHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update, err error) {
	errormessage := fmt.Sprintf("Error on query: %v", err)
	log.Errorf(errormessage)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, errormessage)
	bot.Send(msg)
}

type THandlerKit struct {
	TStandardHandler
	TErrorHandler
}

func CreateHandlerKit(handler TStandardHandler) *THandlerKit {
	return &THandlerKit{TStandardHandler: handler}
}

func CreateHandlerKitEx(handler TStandardHandler, errorHandler TErrorHandler) *THandlerKit {
	return &THandlerKit{TStandardHandler: handler, TErrorHandler: errorHandler}
}

func (qh THandlerKit) Handler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	err := qh.TStandardHandler(bot, update)
	if err != nil {
		if qh.TErrorHandler != nil {
			qh.TErrorHandler(bot, update, err)
		} else {
			StandardErrorHandler(bot, update, err)
		}
	}
}
