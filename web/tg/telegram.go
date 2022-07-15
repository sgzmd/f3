package main

import (
	"log"

	"github.com/jessevdk/go-flags"
	"github.com/sgzmd/f3/web/common"
	"github.com/sgzmd/f3/web/rpc"
	"github.com/sgzmd/f3/web/telegrambot"
)

func main() {

	var opts common.Options
	_, err := flags.Parse(&opts)

	if err != nil {
		return
	}

	pclient, err := rpc.NewClient(&opts.GrpcBackend)
	client := pclient
	if err != nil {
		log.Fatal(err)
	}

	telegramToken := opts.TelegramToken
	telegrambot.BotFunc(telegramToken, client)
}
