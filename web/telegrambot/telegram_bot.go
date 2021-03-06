package telegrambot

import (
	"fmt"
	"github.com/sgzmd/f3/web/telegrambot/intf"
	"github.com/sgzmd/f3/web/webserver/handlers"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/rpc"
)

// Broadly speaking, do we want to have a proper object here with a state and some lifecycle
// rather than dragging every single thing from one method call to another?

type TelegramBotHandler struct {
	bot    intf.IBotApiWrapper
	client rpc.ClientInterface
}

// creates new TelegramBotHandler
func NewTelegramBotHandler(bot intf.IBotApiWrapper, client rpc.ClientInterface) *TelegramBotHandler {
	return &TelegramBotHandler{
		bot:    bot,
		client: client,
	}
}

func (tbh *TelegramBotHandler) ListHandler(update tgbotapi.Update) ([]tgbotapi.Chattable, error) {
	resp, err := tbh.client.ListTrackedEntries(&pb.ListTrackedEntriesRequest{
		UserId: MakeUserKey(update)})
	if err != nil {
		tbh.reportError(update, err)
		return nil, err
	}

	if len(resp.Entry) == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сперва надо на что-нибудь подписаться!")
		tbh.bot.Send(msg)
		return nil, fmt.Errorf("no entries found")
	}

	messages := make([]tgbotapi.Chattable, 0, len(resp.Entry))
	for _, entry := range resp.Entry {
		entryText := formatEntry(
			entry.Key.EntityType,
			entry.EntryName,
			entry.EntryAuthor,
			entry.NumEntries,
			entry.Key.EntityId)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, entryText)

		msg.ParseMode = tgbotapi.ModeHTML
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			"❌ Удалить", fmt.Sprintf("untrack|%s|%d", entry.Key.EntityType, int(entry.Key.EntityId)))))

		messages = append(messages, msg)
	}

	return messages, nil
}

func (tbh *TelegramBotHandler) reportError(update tgbotapi.Update, err error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error: %+v")
	msg.Text = fmt.Sprintf("Error: %+v", err)
	log.Print(msg.Text)
	tbh.bot.Send(msg)
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

func errorToTg(update tgbotapi.Update, text string, err error, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.Text = fmt.Sprintf("Error: %+v", err)
	log.Print(msg.Text)
	bot.Send(msg)
}

func ListCommandHandler(update tgbotapi.Update, client rpc.ClientInterface, bot intf.IBotApiWrapper) ([]tgbotapi.Chattable, error) {
	resp, err := client.ListTrackedEntries(&pb.ListTrackedEntriesRequest{
		UserId: MakeUserKey(update)})
	if err != nil {
		return nil, err
	}

	messages := make([]tgbotapi.Chattable, 0, len(resp.Entry))

	for _, entry := range resp.Entry {
		entryText := formatEntry(entry.Key.EntityType, entry.EntryName, "", entry.NumEntries, entry.Key.EntityId)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, entryText)
		msg.ParseMode = tgbotapi.ModeHTML
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			"❌ Удалить", fmt.Sprintf("untrack|%s|%d", entry.Key.EntityType, int(entry.Key.EntityId)))))

		messages = append(messages, msg)

		bot.Send(msg)
	}

	return messages, nil
}

func MakeUserKey(update tgbotapi.Update) string {
	return handlers.MakeUserKeyFromUserNameAndId(
		update.Message.From.UserName, update.Message.From.ID)
}

func HandleCallbackQuery(update tgbotapi.Update, bot *tgbotapi.BotAPI, client rpc.ClientInterface) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		errorToTg(update, "Couldn't request callback data: %+v", err, bot)
		return
	}

	data := update.CallbackQuery.Data
	req := strings.SplitN(data, "|", 3)
	entryId, err := strconv.Atoi(req[2])
	entryType, ok := pb.EntryType_value[req[1]]

	if len(req) != 3 || err != nil || !ok {
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Received bad callback: "+update.CallbackQuery.Data)
		bot.Send(msg)
	} else {
		from := update.CallbackQuery.From
		if req[0] == "track" {
			resp, err := client.TrackEntry(&pb.TrackEntryRequest{Key: &pb.TrackedEntryKey{
				EntityId:   int64(entryId),
				EntityType: pb.EntryType(entryType),
				UserId:     handlers.MakeUserKeyFromUserNameAndId(from.UserName, from.ID)}})
			if err != nil {
				errorText := fmt.Sprintf("Failed to track story: %+v", err)
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, errorText)
				bot.Send(msg)
				log.Printf(errorText)
			} else if resp.Result == pb.TrackEntryResult_TRACK_ENTRY_RESULT_ALREADY_TRACKED {
				text := "✔️ Уже добавлено"
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
				bot.Send(msg)
			} else if resp.Result == pb.TrackEntryResult_TRACK_ENTRY_RESULT_OK {
				text := "✅️ Добавлено!"
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
				bot.Send(msg)
			}
		} else if req[0] == "untrack" {
			_, err := client.UntrackEntry(&pb.UntrackEntryRequest{Key: &pb.TrackedEntryKey{
				EntityId:   int64(entryId),
				EntityType: pb.EntryType(entryType),
				UserId:     handlers.MakeUserKeyFromUserNameAndId(from.UserName, from.ID)}})
			if err != nil {
				errorText := fmt.Sprintf("Failed to untrack story: %+v", err)
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, errorText)
				bot.Send(msg)
				log.Printf(errorText)
			} else {
				text := "✔️ Удалили"
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
				bot.Send(msg)

			}
		}
	}
}

func SearchCommandHandler(update tgbotapi.Update, client rpc.ClientInterface, bot *tgbotapi.BotAPI) {
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
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
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

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, entryText)
		msg.ParseMode = tgbotapi.ModeHTML
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			"➕ Добавить", fmt.Sprintf("track|%s|%d", entry.EntryType, int(entry.EntryId)))))

		bot.Send(msg)
		numSent++
	}

	if numSent == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error formatting response, check log for details")
		bot.Send(msg)
	}
}

const (
	StartCommand     = "start"
	SearchCommand    = "search"
	ListCommand      = "list"
	CheckUpdates     = "updates"
	MaxSearchEntries = 50
)
