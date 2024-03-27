# Tiny telegream bot router 

## Example telegram bot with standart handler
This code will be print standart error. Use NewTRouterEx for create with own error handler

```
package main

// Hello world handler
func TDownloadYTAudioHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello world")
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := bot.Send(msg)
	return err
}

main() {
    token := "XXXX"
    bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed create telegram bot: %v", err)
	}
    router := trouter.NewTRouter(ser.bot, set)
	downloadYTAudioHandler := trouter.CreateHandlerKit(TDownloadYTAudioHandler)
	router.AddHandler("^\\/hello$", downloadYTAudioHandler.Handler)
	router.AddDefaultHandler(downloadYTAudioHandler.Handler)
	router.Run()
}
```

## Simple example telegram bot with hanler
```
package main

// Hello world handler
func HelloWorldHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello world")
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

main() {
    token := "XXXX"
    bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed create telegram bot: %v", err)
	}
    set := trouter.Settings{Timeout: 100}
    router := trouter.NewTRouter(bot, set)
	router.AddHandler("^\\/hello$", HelloWorldHandler)
	router.AddDefaultHandler(HelloWorldHandler)
	router.Run()
}
```