package main

import (
	"flag"
	"log"

	telegram "github.com/I-Van-Radkov/TelegramBot/client/telegram"

	"github.com/I-Van-Radkov/TelegramBot/internal/bot"
)

func main() {
	token := mustFlag()

	client := telegram.NewClient(token)
	Worker := bot.NewWorker(client)
	log.Print("Telegram service is started")

	if err := Worker.Start(); err != nil {
		log.Fatal("Telegram service is stopped")
	}
}

func mustFlag() string {
	token := flag.String(
		"token",
		"",
		"for access to bot api",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("[ERR] Token is not be found")
	}

	return *token
}
