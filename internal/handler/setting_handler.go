package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/service"
	"github.com/aasumitro/asttax/internal/util/cache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Setting struct {
	bot     *tgbotapi.BotAPI
	userSrv service.IUserService
	cache   *cache.Cache
}

func (h *Setting) EditTradeFee(msg *tgbotapi.Message, fee string) {
	h.notifyProcessingUpdate(msg)

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
	h.notifyProcessingUpdate(msg)

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

func (h *Setting) EditBuySlippage(cmd string, msg *tgbotapi.Message) {
	// state to retrieve user update
	stateKey := fmt.Sprintf("%d_state", msg.Chat.ID)
	messageKey := fmt.Sprintf("%d_msg", msg.Chat.ID)
	if cmd == common.BuySlippage {
		h.promptForUpdate(stateKey, common.AwaitingBuySlippage,
			messageKey, msg)
		return
	}

	// removing user message
	h.remove(msg.Chat.ID, msg.MessageID)

	// get prev msg id
	var prevMessageID int
	if cacheData, ok := h.cache.Get(messageKey); ok {
		if id, ok := cacheData.(int); ok {
			prevMessageID = id
		}
	}

	// sent fallback for update
	npMsg := msg
	npMsg.MessageID = prevMessageID
	h.notifyProcessingUpdate(npMsg)

	ctx := context.Background()
	ctxDur := 5
	ctxWT, done := context.WithTimeout(ctx,
		time.Duration(ctxDur)*time.Second)
	defer done()

	data, err := h.userSrv.SetBuySlippage(ctxWT, prevMessageID, msg)
	if err != nil {
		log.Printf("failed update user data: %v", err)
		return
	}

	// remove state
	h.cache.Delete(stateKey)
	h.cache.Delete(messageKey)
	h.reply(data)
}

func (h *Setting) EditSellAmount(msg *tgbotapi.Message, pos int) {
	fmt.Print(msg, pos)
}

func (h *Setting) EditSellSlippage(cmd string, msg *tgbotapi.Message) {
	// state to retrieve user update
	stateKey := fmt.Sprintf("%d_state", msg.Chat.ID)
	messageKey := fmt.Sprintf("%d_msg", msg.Chat.ID)
	if cmd == common.SellSlippage {
		h.promptForUpdate(stateKey, common.AwaitingSellSlippage,
			messageKey, msg)
		return
	}

	// removing user message
	h.remove(msg.Chat.ID, msg.MessageID)

	// get prev msg id
	var prevMessageID int
	if cacheData, ok := h.cache.Get(messageKey); ok {
		if id, ok := cacheData.(int); ok {
			prevMessageID = id
		}
	}

	// sent fallback for update
	npMsg := msg
	npMsg.MessageID = prevMessageID
	h.notifyProcessingUpdate(npMsg)

	ctx := context.Background()
	ctxDur := 5
	ctxWT, done := context.WithTimeout(ctx,
		time.Duration(ctxDur)*time.Second)
	defer done()

	data, err := h.userSrv.SetSellSlippage(ctxWT, prevMessageID, msg)
	if err != nil {
		log.Printf("failed update user data: %v", err)
		return
	}

	// remove state
	h.cache.Delete(stateKey)
	h.cache.Delete(messageKey)
	h.reply(data)
}

func (h *Setting) EditSellProtection(msg *tgbotapi.Message) {
	h.notifyProcessingUpdate(msg)

	ctx := context.Background()
	ctxDur := 5
	ctxWT, done := context.WithTimeout(ctx,
		time.Duration(ctxDur)*time.Second)
	defer done()

	data, err := h.userSrv.SetSellProtection(ctxWT, msg)
	if err != nil {
		log.Printf("failed update user data: %v", err)
		return
	}
	h.reply(data)
}

func (h *Setting) promptForUpdate(stateKey, stateValue, msgKey string, msg *tgbotapi.Message) {
	cacheDur := 1
	h.cache.Set(stateKey, stateValue, time.Duration(cacheDur)*time.Minute)
	h.cache.Set(msgKey, msg.MessageID, time.Duration(cacheDur)*time.Minute)
	txt := `
*Note*

Please send us a valid decimal value:
- Examples: 1, 2, 3, 0.0, 1.5, 2.75
- Only whole numbers and decimals are allowed.

Please send us your update:
`
	reply := tgbotapi.NewEditMessageText(msg.Chat.ID, msg.MessageID, txt)
	reply.ParseMode = common.MessageParseMarkdown
	h.reply(&reply)
}

func (h *Setting) notifyProcessingUpdate(msg *tgbotapi.Message) {
	txt := "Your update is being processed. Please wait a moment while we apply the changes . . ."
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
		return
	}
}

func (h *Setting) remove(chatID int64, msgID int) {
	deleteMsg := tgbotapi.NewDeleteMessage(chatID, msgID)
	if _, err := h.bot.Request(deleteMsg); err != nil {
		log.Printf("failed delete message: %v", err)
		return
	}
}

func NewSettingHandler(
	bot *tgbotapi.BotAPI,
	userSrv service.IUserService,
	cachePool *cache.Cache,
) *Setting {
	return &Setting{
		bot:     bot,
		userSrv: userSrv,
		cache:   cachePool,
	}
}
