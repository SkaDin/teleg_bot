package main

import (
	"flag"
	"log"
	"teleg_bot/clients/telegram"
)

func main() {
	// token = flags.Get(token)
	tgClient := telegram.New(mustHost(), mustToken())

	_ = tgClient
	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}

func mustHost() string {
	host := flag.String(
		"host-bot-telegram",
		"api.telegram.org",
		"host telegram",
	)
	flag.Parse()
	if *host == "" {
		log.Fatal("host is not specified")
	}
	return *host
}
