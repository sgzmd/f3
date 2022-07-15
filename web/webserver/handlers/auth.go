package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"log"
	"net/url"
	"time"
)

func Auth(ctx ClientContext) func(c *fiber.Ctx) error {
	auth := *ctx.Auth
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
				ui, err := auth.GetUserInfo(params)
				if err != nil {
					log.Printf("Bad user info")
					return c.Redirect(Login)
				} else {
					cookieValue, err := auth.GetCookieValue(params)
					if err != nil {
						log.Printf("Cannot create cookie: %+v", err)
						return c.Redirect(Login)
					} else {
						action := proto.UserInfoAction_USER_INFO_ACTION_CREATE
						resp, err := ctx.RpcClient.GetUserInfo(&proto.GetUserInfoRequest{
							UserId:         MakeUserKeyFromUserNameAndId(ui.UserName, ui.Id),
							UserTelegramId: ui.Id,
							Action:         &action,
						})

						if err != nil {
							log.Printf("Cannot create user: %+v", err)
							return c.Redirect(Login)
						} else {
							if resp.UserCreated == proto.UserCreated_USER_CREATED_YES {
								log.Printf("New user created: %+v", resp)
							} else {
								log.Printf("User already exists: %+v", resp)
							}
						}

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
