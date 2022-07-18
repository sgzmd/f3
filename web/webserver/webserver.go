package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/jessevdk/go-flags"
	"github.com/sgzmd/f3/web/common"
	"github.com/sgzmd/f3/web/rpc"
	"github.com/sgzmd/f3/web/telegrambot"
	"github.com/sgzmd/f3/web/webserver/handlers"
	"github.com/sgzmd/f3/web/webserver/updates"
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"log"
)

var (
	auth tgauth.TelegramAuth
)

func main() {
	var opts common.Options

	_, err := flags.Parse(&opts)

	log.Printf("Startng with options: %+v", opts)

	if err != nil {
		log.Fatal(err)
	}

	auth = tgauth.NewTelegramAuth(opts.TelegramToken, "/login", "/check-auth")
	client, err := rpc.NewClient(&opts.GrpcBackend)
	if err != nil {
		log.Fatal(err)
	}

	clientContext := handlers.ClientContext{
		RpcClient: client,
		Auth:      &auth,
		Opts:      &opts,
	}

	// Starting bot in a parallel thread
	go telegrambot.BotFunc(opts.TelegramToken, client)

	engine := html.New("./templates/web", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Static(handlers.StaticPrefix, "./templates/static")

	app.Use(handlers.Auth(clientContext))

	app.Get("/favicon.ico", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("templates/static/favicon.ico", false)
	})

	app.Get("/", handlers.IndexHandler(clientContext))
	app.Get("/search/:searchTerm", handlers.SearchHandler(clientContext))
	app.Get("/track/:entityType/:id", handlers.TrackUntrackHandler(clientContext, handlers.Track))
	app.Get("/untrack/:entityType/:id", handlers.TrackUntrackHandler(clientContext, handlers.Untrack))
	app.Get("/check-updates-r2d2", func(ctx *fiber.Ctx) error {
		_, e := updates.CheckAndSendUpdates(clientContext, opts.TelegramToken)
		return e
	})
	app.Get(handlers.Login, LoginHandler(opts))

	// Scheduling a forever loop which checks for updates every hour
	go updates.CheckUpdatesLoop(clientContext, opts.TelegramToken)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", opts.WebPort)))
}

func LoginHandler(opts common.Options) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.Render("login", fiber.Map{
			"DomainName": opts.DomainName,
			"BotName":    opts.BotName,
		})
	}
}
