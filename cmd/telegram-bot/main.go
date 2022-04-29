package main

import (
        "log"
        "net/http"

        "github.com/go-telegram-bot-api/telegram-bot-api/v5"

        "main/cmd/telegram-bot/config"
)

func main() {
        cfg := config.New()

        bot, err := tgbotapi.NewBotAPI(cfg.BotSecret)
        if err != nil {
                log.Fatal(err)
        }

        bot.Debug = true

        log.Printf("Authorized on account %s", bot.Self.UserName)

        wh, _ := tgbotapi.NewWebhook(cfg.WHAddr + "/" + bot.Token)

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

        go http.ListenAndServe("0.0.0.0:80", nil)

        for update := range updates {
                log.Printf("%+v\n", update)
		
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

                if update.PreCheckoutQuery != nil {
                        log.Println("Start precheck processing")

			handlePreCheckoutQuery(bot, update)

                        log.Println("Finish precheck processing")
			
			msg = "Поздравляю! Заказ успешно создан"
		}

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
                        log.Fatal(err)
                }

                log.Println("Update finished")
        }


}

func handlePreCheckoutQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	pca := tgbotapi.PreCheckoutConfig{
		OK:                 true,
		PreCheckoutQueryID: update.PreCheckoutQuery.ID,
	}
	_, err := bot.Request(pca)
	if err != nil {
                log.Panic(err)
        }
}
