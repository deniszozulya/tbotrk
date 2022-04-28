package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type config struct {
	BotSecret string `env:"BOT_SECRET" envDefault:"5166381858:AAH-cz97Iz-zpdSJoBGXQcZCupXPC7rGHms"`
	APISecret string `env:"API_SECRET" envDefault:"ce939b6c-611d-47a1-90a6-5054c8492733"`
	WHAddr string `env:"WEBHOOK_ADDRESS" envDefault:"https://bokovski.ru/work/tg-bot/index.html"`
	Debug bool `env:"DEBUG"`
	ServerAddress string `env:"SERVER_ADDRESS"`
	CertFilepath string `env:"CERT_FILE" envDefault:"cert.pem"`
	KeyFilepath string `env:"KEY_FILE" envDefault:"key.pem"`
}

func New() *config {
	var c config

	err := env.Parse(&c)
	if err != nil {
		log.Fatal(err)
	}

	return &c
}