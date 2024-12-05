package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Setting struct {
	bot     *tgbotapi.BotAPI
	userSrv service.IUserService
}

func (h *Setting) EditTradeFee(msg *tgbotapi.Message, fee string) {
	h.fallback(msg)

	ctx := context.Background()
	ctxDur := 5
	ctxWT, done := context.WithTimeout(ctx,
		time.Duration(ctxDur)*time.Second)
	defer done()

	data, err := h.userSrv.SetTradeFee(ctxWT, msg, fee)
	if err != nil {
		log.Printf("FAILED_UPDATE_USER_DATA: %v", err)
		return
	}
	h.reply(data)
}

func (h *Setting) EditConfirmTrade(msg *tgbotapi.Message) {
	h.fallback(msg)

	ctx := context.Background()
	ctxDur := 5
	ctxWT, done := context.WithTimeout(ctx,
		time.Duration(ctxDur)*time.Second)
	defer done()

	data, err := h.userSrv.SetConfirmTrade(ctxWT, msg)
	if err != nil {
		log.Printf("FAILED_UPDATE_USER_DATA: %v", err)
		return
	}
	h.reply(data)
}

func (h *Setting) EditBuyAmount(msg *tgbotapi.Message, pos int) {
	fmt.Print(msg, pos)
}

func (h *Setting) EditBuySlippage(msg *tgbotapi.Message) {
	fmt.Print(msg)
}

func (h *Setting) EditSellAmount(msg *tgbotapi.Message, pos int) {
	fmt.Print(msg, pos)
}

func (h *Setting) EditSellSlippage(msg *tgbotapi.Message) {
	fmt.Print(msg)
}

func (h *Setting) EditSellProtection(msg *tgbotapi.Message) {
	h.fallback(msg)

	ctx := context.Background()
	ctxDur := 5
	ctxWT, done := context.WithTimeout(ctx,
		time.Duration(ctxDur)*time.Second)
	defer done()

	data, err := h.userSrv.SetSellProtection(ctxWT, msg)
	if err != nil {
		log.Printf("FAILED_UPDATE_USER_DATA: %v", err)
		return
	}
	h.reply(data)
}

func (h *Setting) fallback(msg *tgbotapi.Message) {
	txt := "Your update is being processed. Please wait a moment while we apply the changes..."
	reply := tgbotapi.NewEditMessageText(msg.Chat.ID, msg.MessageID, txt)
	reply.ParseMode = common.MessageParseMarkdown
	h.reply(&reply)
}

func (h *Setting) reply(r interface{}) {
	switch msg := r.(type) {
	case *tgbotapi.MessageConfig:
		if msg.Text == "" && msg.ChatID == 0 {
			log.Println(common.ErrNoMsg)
			return
		}
		if _, err := h.bot.Send(msg); err != nil {
			log.Printf("error sending reply: %v\n", err)
			return
		}
	case *tgbotapi.EditMessageTextConfig:
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

func NewSettingHandler(
	bot *tgbotapi.BotAPI,
	userSrv service.IUserService,
) *Setting {
	return &Setting{bot: bot, userSrv: userSrv}
}
