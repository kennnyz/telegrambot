package main

import (
	"flag"
	tgClient "github.com/kennnyz/telegrambot/clients/telegram"
	event_consumer "github.com/kennnyz/telegrambot/consumer/event-consumer"
	"github.com/kennnyz/telegrambot/events/telegram"
	"github.com/kennnyz/telegrambot/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Print("service started...")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal()
	}

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token not specified")
	}
	return *token
}
