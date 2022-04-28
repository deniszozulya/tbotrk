package telegram

import (
	"fmt"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	bot tgbotapi.BotAPI
	CertFilepath string
}

func (t *Telegram) New(token, WHAddr string, debug bool) error {	
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = debug

	wh, _ := tgbotapi.NewWebhookWithCert(WHAddr+bot.Token, tgbotapi.FilePath(t.CertFilepath))

	_, err = bot.Request(wh)
	if err != nil {
		return err
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		return err
	}

	if info.LastErrorDate != 0 {
		return fmt.Errorf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	t.bot = *bot

	return nil
}



func (t *Telegram) Start() tgbotapi.UpdatesChannel {

	log.Printf("Authorized on account %s", t.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

	for update := range updates {

		fmt.Println(update)

		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		orderBtn := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("Хочу кушать!", "http://1.com"),
			),
		)

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "start":
			msg.Text = "Что умеет бот?\nЯ помогу быстро заказать еду и организую доставку"

			msg.ReplyMarkup = orderBtn
		case "web_app_setup_main_button":
			msg.Text = "Серёжа молодец!"
		case "message":
			msg.Text = "message"
		default:
			msg.Text = "Я еще не обучен этой команде ;("
		}

		if _, err := t.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	return t.bot.ListenForWebhook("/" + t.bot.Token)
}