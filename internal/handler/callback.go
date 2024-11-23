package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/service"
	"github.com/aasumitro/asttax/internal/template/keyboard"
	"github.com/aasumitro/asttax/internal/template/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Callback struct {
	bot     *tgbotapi.BotAPI
	userSrv service.IUserService
}

func (h *Callback) AcceptAgreement(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageText(msg.Chat.ID,
		msg.MessageID, message.ConfirmAgreementTextBody())
	reply.ParseMode = common.MessageParseHTML
	h.reply(reply)

	time.Sleep(time.Second)

	reply = tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.AccountCreatedTextBody(),
		keyboard.AccountCreatedKeyboardMarkup)
	reply.ParseMode = common.MessageParseHTML
	h.reply(reply)
}

func (h *Callback) Start(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID, msg.MessageID,
		message.StartTextBody(), keyboard.StartKeyboardMarkup)
	reply.ParseMode = common.MessageParseMarkdown
	h.reply(reply)
}

func (h *Callback) Buy(msg *tgbotapi.Message) {
	fmt.Println(msg.Chat.ID)
}

func (h *Callback) Sell(msg *tgbotapi.Message) {
	fmt.Println(msg.Chat.ID)
}

func (h *Callback) Positions(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.NoPositionTextBody(), keyboard.PositionKeyboardMarkup)
	reply.ParseMode = common.MessageParseHTML
	h.reply(reply)
}

func (h *Callback) NewPair(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.ComingSoonTextBody("New Pair"), keyboard.BackToStartKeyboardMarkup)
	reply.ParseMode = common.MessageParseHTML
	h.reply(reply)
}

func (h *Callback) Setting(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.SettingTextBody(), keyboard.SettingKeyboardMarkup)
	reply.ParseMode = common.MessageParseHTML
	h.reply(reply)
}

func (h *Callback) LanguageSetting(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.SettingTextBody(), keyboard.LanguageSettingKeyboardMarkup)
	reply.ParseMode = common.MessageParseHTML
	h.reply(reply)
}

func (h *Callback) Help(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.HelpTextBody(), keyboard.BackToStartKeyboardMarkup)
	reply.ParseMode = common.MessageParseHTML
	h.reply(reply)
}

func (h *Callback) BackToStart(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.StartTextBody(), keyboard.StartKeyboardMarkup)
	reply.ParseMode = common.MessageParseMarkdown
	h.reply(reply)
}

func (h *Callback) BackToSetting(msg *tgbotapi.Message) {
	h.Setting(msg)
}

func (h *Callback) Refresh(msg *tgbotapi.Message) {
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.StartTextBody(), keyboard.StartKeyboardMarkup)
	reply.ParseMode = common.MessageParseMarkdown
	h.reply(reply)
}

func (h *Callback) reply(msg tgbotapi.EditMessageTextConfig) {
	if msg.Text == "" && msg.ChatID == 0 {
		log.Println(common.ErrNoMsg)
		return
	}
	if _, err := h.bot.Send(msg); err != nil {
		log.Printf("error sending reply: %v\n", err)
		return
	}
}

func NewCallbackHandler(
	bot *tgbotapi.BotAPI,
	userSrv service.IUserService,
) *Callback {
	return &Callback{bot: bot, userSrv: userSrv}
}
