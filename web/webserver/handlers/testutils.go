package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/sgzmd/f3/web/common"
	"github.com/sgzmd/f3/web/rpc"
	"github.com/sgzmd/go-telegram-auth/testing"
)

func CreateTestingApp() (*fiber.App, ClientContext) {
	engine := html.New("../../templates/web", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	auth := testing.NewFakeTelegramAuth(true, "testuser")
	client, _ := rpc.NewClient(nil)

	clientContext := ClientContext{
		RpcClient: client,
		Auth:      &auth,
		Opts: &common.Options{
			GrpcBackend:   "",
			TelegramToken: "",
			WebPort:       0,
			BotName:       "TestBot",
			DomainName:    "test.com",
		},
	}

	app.Use(Auth(clientContext))
	return app, clientContext
}
