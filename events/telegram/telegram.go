package telegram

import "teleg_bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}
