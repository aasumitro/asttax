package handler

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/service"
	"github.com/aasumitro/asttax/internal/template/keyboard"
	"github.com/aasumitro/asttax/internal/template/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command struct {
	bot     *tgbotapi.BotAPI
	userSrv service.IUserService
}

func (h *Command) Start(msg *tgbotapi.Message) {
	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx, time.Second)
	defer done()

	data, err := h.userSrv.LoadUser(ctxWT,
		msg, true, false)
	if err != nil {
		log.Printf("failed to laod user: %v", err)
		return
	}

	h.reply(data)
}

func (h *Command) Buy(msg *tgbotapi.Message) {
	fmt.Println(msg.Chat.ID)
}

func (h *Command) Sell(msg *tgbotapi.Message) {
	fmt.Println(msg.Chat.ID)
}

func (h *Command) Positions(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, message.NoPositionTextBody())
	reply.ParseMode = common.MessageParseHTML
	reply.ReplyMarkup = keyboard.PositionKeyboardMarkup
	reply.ReplyToMessageID = msg.MessageID
	h.reply(&reply)
}

func (h *Command) Settings(msg *tgbotapi.Message) {
	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx, time.Second)
	defer done()
	data, err := h.userSrv.LoadUserSetting(ctxWT,
		msg, true)
	if err != nil {
		log.Printf("failed to laod user: %v", err)
		return
	}
	h.reply(data)
}

func (h *Command) Withdraw(msg *tgbotapi.Message) {
	fmt.Println(msg.Chat.ID)
}

func (h *Command) Help(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, message.HelpTextBody())
	reply.ParseMode = common.MessageParseHTML
	reply.ReplyMarkup = keyboard.BackToStartKeyboardMarkup
	h.reply(&reply)
}

func (h *Command) Backup(msg *tgbotapi.Message) {
	fmt.Println(msg.Chat.ID)
}

func (h *Command) reply(r interface{}) {
	fmt.Println(r != nil, reflect.TypeOf(r))

	switch msg := r.(type) {
	case *tgbotapi.MessageConfig: // Pointer type
		fmt.Println("sent *MessageConfig")
		if msg.Text == "" && msg.ChatID == 0 {
			log.Println(common.ErrNoMsg)
			return
		}
		if _, err := h.bot.Send(msg); err != nil {
			log.Printf("error sending reply: %v\n", err)
			return
		}

	case *tgbotapi.EditMessageTextConfig: // Pointer type
		fmt.Println("sent *EditMessageTextConfig")
		if msg.Text == "" && msg.ChatID == 0 {
			log.Println(common.ErrNoMsg)
			return
		}
		if _, err := h.bot.Send(msg); err != nil {
			log.Printf("error sending reply: %v\n", err)
			return
		}

	default:
		log.Printf("unexpected reply type: %T", r)
	}
}

func NewCommandHandler(
	bot *tgbotapi.BotAPI,
	userSrv service.IUserService,
) *Command {
	return &Command{bot: bot, userSrv: userSrv}
}
