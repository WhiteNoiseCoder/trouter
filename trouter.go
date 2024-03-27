package trouter

import (
	"regexp"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// THandler is type for handler telegram request
type THandler = func(bot *tgbotapi.BotAPI, update *tgbotapi.Update)

// TRouter is telegram router
type TRouter struct {
	handlers       map[string]THandler
	defaultHandler THandler

	bot      *tgbotapi.BotAPI
	settings *Settings
}

// NewTRouter is router constructor
func NewTRouter(bot *tgbotapi.BotAPI, settings *Settings) *TRouter {
	return &TRouter{bot: bot, settings: settings, handlers: make(map[string]THandler)}

}

// AddHandler is func for add telegram query handler
func (r *TRouter) AddHandler(match string, h THandler) {
	r.handlers[match] = h
}

// AddDefaultHandler is func for add telegram handler which handle query without scpecial command
func (r *TRouter) AddDefaultHandler(h THandler) {
	r.defaultHandler = h
}

// Run is func for start routing
func (r TRouter) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = r.settings.Timeout

	updates := r.bot.GetUpdatesChan(u)
	for update := range updates {
		go r.handle(update)
	}
	return nil
}

func (r TRouter) handle(update tgbotapi.Update) {
	for m, h := range r.handlers {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("panic on regexp %s in router: %v", m, err)
			}
		}()
		matched := regexp.MustCompile(m).MatchString(update.Message.Text)
		if matched {
			h(r.bot, &update)
			return
		}
	}
	if r.defaultHandler != nil {
		r.defaultHandler(r.bot, &update)
	}
}
