package main

import (
	"log"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/jessevdk/go-flags"
	"github.com/sgzmd/f3/web/common"
	"github.com/sgzmd/f3/web/rpc"
	"github.com/sgzmd/f3/web/webserver/handlers"
	"github.com/sgzmd/go-telegram-auth/tgauth"
)

const (
	StaticPrefix = "/static/"
	Login        = "/login"
	CheckAuth    = "/check-auth"
)

var (
	auth tgauth.TelegramAuth
)

func main() {
	var opts common.Options

	_, err := flags.Parse(&opts)

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
	}

	engine := html.New("./templates/web", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Static(StaticPrefix, "./templates/static")

	app.Use(AuthMiddleware())

	app.Get("/", handlers.IndexHandler(clientContext))
	app.Get("/search/:searchTerm", handlers.SearchHandler(clientContext))
	app.Get("/track/:entityType/:id", handlers.TrackUntrackHandler(clientContext, handlers.Track))
	app.Get("/untrack/:entityType/:id", handlers.TrackUntrackHandler(clientContext, handlers.Untrack))
	app.Get(Login, LoginHandler())

	log.Fatal(app.Listen(":8080"))
}

func LoginHandler() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.Render("login", fiber.Map{})
	}
}

type FakeSearchResult struct {
	Name string
}

func SearchHandler() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		_ = ctx.Params("searchTerm", "")
		userInfo := ctx.Locals("user").(*tgauth.UserInfo)
		sr := []FakeSearchResult{
			{Name: "Some Name"},
		}
		return ctx.Render("index", fiber.Map{"Name": userInfo.FirstName, "HasSearchResults": true, "SearchResults": sr})
	}
}

func IndexHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		ui := c.Locals("user")
		userInfo := ui.(*tgauth.UserInfo)

		return c.Render("index", fiber.Map{"Name": userInfo.FirstName})
	}
}

func AuthMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Path() == Login {
			log.Printf("Login page loading ...")
			return c.Next()
		}

		if c.Path() == CheckAuth {
			url, err := url.Parse(c.Request().URI().String())
			if err != nil {
				log.Printf("Bad URL: %+v", err)
				return c.Redirect(Login)
			}
			params := valuesToParams(url.Query())

			if ok, _ := auth.CheckAuth(params); !ok {
				log.Printf("Bad auth")
				return c.Redirect(Login)
			} else {
				log.Printf("Auth OK, proceeding...")
				_, err := auth.GetUserInfo(params)
				if err != nil {
					log.Printf("Bad user info")
					return c.Redirect(Login)
				} else {
					cookieValue, err := auth.GetCookieValue(params)
					if err != nil {
						log.Printf("Cannot create cookie: %+v", err)
						return c.Redirect(Login)
					} else {
						cookie := &fiber.Cookie{
							Name:    tgauth.DefaultCookieName,
							Value:   cookieValue,
							Path:    "/",
							Expires: time.Now().Add(time.Hour + 24),
						}
						c.Cookie(cookie)
						log.Printf("Cookie set: %+v", cookie)
						return c.Redirect("/")
					}
				}
			}
		}

		cookieValue := c.Cookies(tgauth.DefaultCookieName, "")
		if cookieValue == "" {
			log.Printf("No cookie, redirecting to /login")
			return c.Redirect(Login)
		}

		params, err := auth.GetParamsFromCookieValue(cookieValue)
		if err != nil {
			log.Printf("Cookie is invalid, redirecting to /login")
			return c.Redirect(Login)
		}
		ui, err := auth.GetUserInfo(params)
		if err != nil {
			log.Printf("Can't get user info from params: %+v", params)
			return c.Redirect(Login)
		}

		c.Locals("user", ui)
		return c.Next()
	}
}

func valuesToParams(values map[string][]string) tgauth.Params {
	p := tgauth.Params{}
	for k, v := range values {
		p[k] = v[0]
	}
	return p
}
