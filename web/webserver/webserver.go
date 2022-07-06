package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/jessevdk/go-flags"
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"log"
	"net/url"
)

const (
	StaticPrefix = "/static/"
)

var (
	useFakes    *bool
	grpcBackend *string

	auth tgauth.TelegramAuth
)

type Options struct {
	GrpcBackend   string `short:"g" long:"grpc_backend" description:"GRPC Backend to use"`
	TelegramToken string `short:"t" long:"telegram_token" description:"Telegram token to use" required:"true"`
}

func main() {
	var opts Options
	_, err := flags.Parse(&opts)

	if err != nil {
		log.Fatal(err)
	}

	auth = tgauth.NewTelegramAuth(opts.TelegramToken, "/login", "/check-auth")

	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Static(StaticPrefix, "./templates/static")

	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/login" {
			log.Printf("Login page loading ...")
			return c.Next()
		}

		if c.Path() == "/check-auth" {
			log.Printf("/check-auth")

			url, err := url.Parse(c.Request().URI().String())
			if err != nil {
				log.Printf("Bad URL: %+v", err)
				return c.Redirect("/login")
			}
			params := url.Query()
			if ok, _ := auth.CheckAuth(params); !ok {
				log.Printf("Bad auth")
				return c.Redirect("/login")
			} else {
				log.Printf("Auth OK, proceeding...")
				_, err := auth.GetUserInfo(params)
				if err != nil {
					log.Printf("Bad user info")
					return c.Redirect("/login")
				} else {
					cookie, err := auth.CreateCookie(params)
					if err != nil {
						log.Printf("Cannot create cookie: %+v", err)
						return c.Redirect("/login")
					} else {
						c.Cookie(&fiber.Cookie{
							Name:    cookie.Name,
							Value:   cookie.Value,
							Path:    cookie.Path,
							Expires: cookie.Expires,
						})
						log.Printf("Cookie set: %+v", cookie)
						return c.Redirect("/")
					}
				}
			}
		}

		cookieValue := c.Cookies(tgauth.DefaultCookieName, "")
		if cookieValue == "" {
			log.Printf("No cookie, redirecting to /login")
			return c.Redirect("/login")
		}

		params, err := auth.GetParamsFromCookieValue(cookieValue)
		if err != nil {
			log.Printf("Cookie is invalid, redirecting to /login")
			return c.Redirect("/login")
		}
		_, err = auth.GetUserInfo(params)
		if err != nil {
			log.Printf("Can't get user info from params: %+v", params)
			return c.Redirect("/login")
		}

		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return nil
	})
	app.Get("/search/:term", func(ctx *fiber.Ctx) error {
		term := ctx.Params("term", "")

		log.Printf("/search/%s", term)

		return nil
	})
	app.Get("/track/:entityType/:id", func(ctx *fiber.Ctx) error {
		entityType, err := ctx.ParamsInt("entityType", -1)
		if err != nil {
			return fiber.ErrBadRequest
		}

		entityId, err := ctx.ParamsInt("id", -1)
		if err != nil {
			return fiber.ErrBadRequest
		}

		log.Printf("/track/type=%d/id=%d", entityType, entityId)

		return nil
	})
	app.Get("/login", func(ctx *fiber.Ctx) error {
		return ctx.Render("login", fiber.Map{})
	})

	log.Fatal(app.Listen(":8080"))
}

func fiberParamsToHttpParams(params map[string]string) map[string][]string {
	m := make(map[string][]string, len(params))
	for k, v := range params {
		m[k] = []string{v}
	}
	return m
}
