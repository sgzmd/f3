package main

import (
	"fmt"
	tb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jessevdk/go-flags"
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/rpc"
	"log"
	"strconv"
	"strings"
)

type Options struct {
	GrpcBackend   string `short:"g" long:"grpc_backend" description:"GRPC Backend to use"`
	TelegramToken string `short:"t" long:"telegram_token" description:"Telegram token to use" required:"true"`
}

const (
	StartCommand     = "start"
	SearchCommand    = "search"
	ListCommand      = "list"
	MaxSearchEntries = 50
)

func main() {

	var opts Options
	_, err := flags.Parse(&opts)

	if err != nil {
		return
	}

	bot, err := tb.NewBotAPI(opts.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	pclient, err := rpc.NewClient(&opts.GrpcBackend)
	client := *pclient
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tb.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case StartCommand:
					msg := tb.NewMessage(update.Message.Chat.ID, "Hello world!")
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
				case SearchCommand:
					searchCommandHandler(update, client, bot)
				case ListCommand:
					listCommandHandler(update, client, bot)
				}

			}
		} else if update.CallbackQuery != nil {
			handleCallbackQuery(update, bot, client)
		}
	}
}

func listCommandHandler(update tb.Update, client rpc.ClientInterface, bot *tb.BotAPI) {
	resp, err := client.ListTrackedEntries(&pb.ListTrackedEntriesRequest{UserId: update.Message.From.UserName})
	if err != nil {
		errorToTg(update, "Error listing entries: %+v", err, bot)
		return
	}

	for _, entry := range resp.Entry {
		entryText := formatEntry(entry.Key.EntityType, entry.EntryName, "", entry.NumEntries, entry.Key.EntityId)
		msg := tb.NewMessage(update.Message.Chat.ID, entryText)
		msg.ParseMode = tb.ModeHTML
		bot.Send(msg)
	}

}

func handleCallbackQuery(update tb.Update, bot *tb.BotAPI, client rpc.ClientInterface) {
	callback := tb.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		errorToTg(update, "Couldn't request callback data: %+v", err, bot)
		return
	}

	data := update.CallbackQuery.Data
	req := strings.SplitN(data, "|", 2)
	entryId, err := strconv.Atoi(req[1])
	entryType, ok := pb.EntryType_value[req[0]]
	if len(req) != 2 || err != nil || !ok {
		msg := tb.NewMessage(update.CallbackQuery.Message.Chat.ID, "Received bad callback: "+update.CallbackQuery.Data)
		bot.Send(msg)
	} else {
		resp, err := client.TrackEntry(&pb.TrackEntryRequest{Key: &pb.TrackedEntryKey{
			EntityId: int64(entryId), EntityType: pb.EntryType(entryType), UserId: update.CallbackQuery.From.UserName}})
		if err != nil {
			errorText := fmt.Sprintf("Failed to track story: %+v", err)
			msg := tb.NewMessage(update.CallbackQuery.Message.Chat.ID, errorText)
			bot.Send(msg)
			log.Printf(errorText)
		} else if resp.Result == pb.TrackEntryResult_TRACK_ENTRY_RESULT_ALREADY_TRACKED {
			text := "✔️ Уже добавлено"
			msg := tb.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
			bot.Send(msg)
		} else if resp.Result == pb.TrackEntryResult_TRACK_ENTRY_RESULT_OK {
			text := "✅️ Добавлено!"
			msg := tb.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
			bot.Send(msg)
		}
	}
}

func searchCommandHandler(update tb.Update, client rpc.ClientInterface, bot *tb.BotAPI) {
	text := strings.Replace(update.Message.Text, "/"+SearchCommand, "", -1)

	log.Printf("Searching for %s", text)

	resp, err := client.GlobalSearch(&pb.GlobalSearchRequest{
		SearchTerm: text,
	})
	if err != nil {
		errorToTg(update, text, err, bot)
		return
	}

	if len(resp.Entry) > MaxSearchEntries {
		msg := tb.NewMessage(update.Message.Chat.ID, text)
		msg.Text = fmt.Sprintf("Too many search results: %d", len(resp.Entry))
		bot.Send(msg)
		return
	}

	numSent := 0
	for _, entry := range resp.Entry {

		entryText := formatEntry(entry.EntryType, entry.EntryName, entry.Author, entry.NumEntities, entry.EntryId)

		if entryText == "" {
			break
		}

		msg := tb.NewMessage(update.Message.Chat.ID, entryText)
		msg.ParseMode = tb.ModeHTML
		msg.ReplyMarkup = tb.NewInlineKeyboardMarkup(tb.NewInlineKeyboardRow(tb.NewInlineKeyboardButtonData(
			"➕ Добавить", fmt.Sprintf("%s|%d", entry.EntryType, int(entry.EntryId)))))

		bot.Send(msg)
		numSent++
	}

	if numSent == 0 {
		msg := tb.NewMessage(update.Message.Chat.ID, "Error formatting response, check log for details")
		bot.Send(msg)
	}
}

func formatEntry(entryType pb.EntryType, entryName string, entryAuthor string, numEntities int32, entryId int64) string {
	var entryText string
	switch entryType {
	case pb.EntryType_ENTRY_TYPE_SERIES:
		entryText = fmt.Sprintf("📚 <b>%s</b> - %s (%d книг) \n\n<a href='http://flibusta.is/s/%d'>Открыть</>", entryName, entryAuthor, numEntities, entryId)
	case pb.EntryType_ENTRY_TYPE_AUTHOR:
		entryText = fmt.Sprintf("🧑 <b>%s</b>  (%d книг) \n\n<a href='http://flibusta.is/a/%d'>Открыть</a>", entryAuthor, numEntities, entryId)
	default:
		entryText = ""
	}
	return entryText
}

func errorToTg(update tb.Update, text string, err error, bot *tb.BotAPI) {
	msg := tb.NewMessage(update.Message.Chat.ID, text)
	msg.Text = fmt.Sprintf("Error: %+v", err)
	log.Print(msg.Text)
	bot.Send(msg)
}
