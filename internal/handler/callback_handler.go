package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/template/keyboard"
	"github.com/aasumitro/asttax/internal/template/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) AcceptAgreementCallback(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageText(msg.Chat.ID,
		msg.MessageID, message.ConfirmAgreementCallbackTextBody())
	reply.ParseMode = common.MessageParseHTML
	h.reply(&reply)

	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
	defer done()

	data, err := h.userSrv.CreateUser(ctxWT, msg)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return
	}

	h.reply(data)
}

func (h *Handler) StartCallback(msg *tgbotapi.Message) {
	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
	defer done()

	data, err := h.userSrv.LoadUser(ctxWT, msg, false, false)
	if err != nil {
		log.Printf("failed to laod user: %v", err)
		return
	}

	h.reply(data)
}

func (h *Handler) BuyCallback(msg *tgbotapi.Message) {
	// TODO
	fmt.Println(msg.Chat.ID)
}

func (h *Handler) SellCallback(msg *tgbotapi.Message) {
	// TODO
	fmt.Println(msg.Chat.ID)
}

func (h *Handler) TrenchesCallback(msg *tgbotapi.Message, state string) {
	// TODO: apply from pump fun
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.TrenchesTextBody(state),
		keyboard.TrenchesKeyboardMarkup(state))
	reply.ParseMode = common.MessageParseHTML
	h.reply(&reply)
}

func (h *Handler) PositionsCallback(msg *tgbotapi.Message) {
	// TODO: apply from main net
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.NoPositionTextBody,
		keyboard.PositionKeyboardMarkup)
	reply.ParseMode = common.MessageParseHTML
	h.reply(&reply)
}

func (h *Handler) SettingCallback(msg *tgbotapi.Message) {
	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
	defer done()
	data, err := h.userSrv.LoadUserSetting(ctxWT,
		msg, false)
	if err != nil {
		log.Printf("failed to laod user: %v", err)
		return
	}
	h.reply(data)
}

func (h *Handler) LanguageSettingCallback(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.SettingTextBody(), keyboard.LanguageSettingKeyboardMarkup)
	reply.ParseMode = common.MessageParseHTML
	h.reply(&reply)
}

func (h *Handler) HelpCallback(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.HelpTextBody(), keyboard.BackToStartKeyboardMarkup)
	reply.ParseMode = common.MessageParseHTML
	h.reply(&reply)
}

func (h *Handler) BackToStartCallback(msg *tgbotapi.Message) {
	h.StartCallback(msg)
}

func (h *Handler) BackToSettingCallback(msg *tgbotapi.Message) {
	h.SettingCallback(msg)
}

func (h *Handler) RefreshCallback(msg *tgbotapi.Message) {
	h.StartCallback(msg)
}
