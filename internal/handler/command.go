package handler

import (
	"context"
	"fmt"
	"log"
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
	reply, err := h.userSrv.LoadUser(ctxWT, msg)
	if err != nil {
		log.Println(err)
		return
	}
	h.reply(*reply)
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
	h.reply(reply)
}

func (h *Command) Settings(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, message.SettingTextBody())
	reply.ParseMode = common.MessageParseHTML
	reply.ReplyMarkup = keyboard.SettingKeyboardMarkup
	reply.ReplyToMessageID = msg.MessageID
	h.reply(reply)
}

func (h *Command) Withdraw(msg *tgbotapi.Message) {
	fmt.Println(msg.Chat.ID)
}

func (h *Command) Help(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, message.HelpTextBody())
	reply.ParseMode = common.MessageParseHTML
	reply.ReplyMarkup = keyboard.BackToStartKeyboardMarkup
	h.reply(reply)
}

func (h *Command) Backup(msg *tgbotapi.Message) {
	fmt.Println(msg.Chat.ID)
}

func (h *Command) reply(msg tgbotapi.MessageConfig) {
	if msg.Text == "" && msg.ChatID == 0 {
		log.Println(common.ErrNoMsg)
		return
	}
	if _, err := h.bot.Send(msg); err != nil {
		log.Printf("error sending reply: %v\n", err)
		return
	}
}

func NewCommandHandler(
	bot *tgbotapi.BotAPI,
	userSrv service.IUserService,
) *Command {
	return &Command{bot: bot, userSrv: userSrv}
}
