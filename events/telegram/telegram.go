package telegram

import "github.com/kennnyz/telegrambot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}
