package common

type Options struct {
	GrpcBackend   string `short:"g" long:"grpc_backend" description:"GRPC Backend to use"`
	TelegramToken string `short:"t" long:"telegram_token" description:"Telegram token to use" required:"true"`
	WebPort       int    `short:"p" long:"web_port" description:"Web server port" default:"8080"`
	BotName       string `short:"b" long:"bot_name" description:"Bot name" default:"F3"`
	DomainName    string `short:"d" long:"domain_name" description:"Domain name"`
}
