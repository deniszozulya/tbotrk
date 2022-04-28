package main

import (
	"log"
	"net/http"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"main/cmd/telegram-bot/config"
	"main/internal/server"
)

func main() {
	cfg := config.New()

	bot, err := tgbotapi.NewBotAPI(cfg.BotSecret)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	wh, _ := tgbotapi.NewWebhookWithCert(cfg.WHAddr+bot.Token, tgbotapi.FilePath(cfg.CertFilepath))

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	
	go http.ListenAndServeTLS("0.0.0.0:8443", cfg.CertFilepath, cfg.KeyFilepath, nil)

	for update := range updates {
		log.Printf("%+v\n", update)


		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
	
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "start":
			msg.Text = "Что умеет бот?\nЯ помогу быстро заказать еду и организую доставку"
		case "web_app_setup_main_button":
			msg.Text = "Серёжа молодец!"
		case "message":
			msg.Text = "message"
		default:
			msg.Text = "Я еще не обучен этой команде ;("
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	s := server.New(cfg.ServerAddress, cfg.CertFilepath, cfg.KeyFilepath)

	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}