package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/jessevdk/go-flags"
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"log"
	"net/url"
	"time"
)

const (
	StaticPrefix = "/static/"
	Login        = "/login"
	CheckAuth    = "/check-auth"
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

	engine := html.New("./templates/web", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Static(StaticPrefix, "./templates/static")

	app.Use(AuthMiddleware())

	app.Get("/", IndexHandler())
	app.Get("/search/:searchTerm", SearchHandler())
	app.Get("/track/:entityType/:id", TrackHandler())
	app.Get(Login, LoginHandler())

	log.Fatal(app.Listen(":8080"))
}

func LoginHandler() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.Render("login", fiber.Map{})
	}
}

func TrackHandler() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
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
			log.Printf(CheckAuth)

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
