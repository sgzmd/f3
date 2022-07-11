package handlers

import (
	"github.com/sgzmd/f3/web/rpc"
	"github.com/sgzmd/go-telegram-auth/tgauth"
)

type ClientContext struct {
	RpcClient rpc.ClientInterface
	Auth      *tgauth.TelegramAuth
}
