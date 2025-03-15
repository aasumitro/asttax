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

func (h *Handler) StartCommand(msg *tgbotapi.Message) {
	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
	defer done()

	data, err := h.userSrv.LoadUser(ctxWT,
		msg, true, false)
	if err != nil {
		log.Printf("failed to laod user: %v", err)
		return
	}

	h.reply(data)
}

func (h *Handler) BuyCommand(msg *tgbotapi.Message) {
	// TODO
	fmt.Println(msg.Chat.ID)
}

func (h *Handler) SellCommand(msg *tgbotapi.Message) {
	// TODO
	fmt.Println(msg.Chat.ID)
}

func (h *Handler) PositionsCommand(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, message.NoPositionTextBody)
	reply.ParseMode = common.MessageParseHTML
	reply.ReplyMarkup = keyboard.PositionKeyboardMarkup
	reply.ReplyToMessageID = msg.MessageID
	h.reply(&reply)
}

func (h *Handler) SettingsCommand(msg *tgbotapi.Message) {
	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
	defer done()
	data, err := h.userSrv.LoadUserSetting(ctxWT,
		msg, true)
	if err != nil {
		log.Printf("failed to laod user: %v", err)
		return
	}
	h.reply(data)
}

func (h *Handler) WithdrawCommand(msg *tgbotapi.Message) {
	// TODO
	fmt.Println(msg.Chat.ID)
	// display list of token
	// -> after select display
	// how much to withdraw (50%, 100%, X% [edit]) - first button
	// full button X {TOKEN_NAME} [edit] - by token size
	// full button set withdraw address
	// if total token selected && wallet address set display withdraw button (CAPS)
	// display notify -
	//
	//
	// Withdraw $SOL — (Solana) 🅴
	//
	// Balance: 0 SOL
	//
	// 🟠 Sending withdrawal transaction...                  ---- process
	// 🔴 Withdrawal transaction failed   View on Solscan    ---- error
	// 🟢 Withdrawal transaction success   View on Solscan   ---- success

}

func (h *Handler) HelpCommand(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, message.HelpTextBody())
	reply.ParseMode = common.MessageParseHTML
	reply.ReplyMarkup = keyboard.BackToStartKeyboardMarkup
	h.reply(&reply)
}

func (h *Handler) BackupCommand(msg *tgbotapi.Message) {
	// TODO
	fmt.Println(msg.Chat.ID)
}
