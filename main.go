package main

import (
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	//tgClient := telegram.New(tgBotHost, mustToken())

	//consumer.Start(fetcher, processor) Получает событие и обрабатывает их
	//Fetcher хранит все данные
	//Processor обрабатывает событие
	//token = flags.Get(token)
	//tgClient = telegram.New(token)
	//fetcher := fetcher.New()
	//processor := processor.New()
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token not specified")
	}
	return *token
}
