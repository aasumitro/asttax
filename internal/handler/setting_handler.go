package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aasumitro/asttax/internal/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) EditTradeFeeState(msg *tgbotapi.Message, fee string) {
	h.notifyProcessingUpdate(msg)

	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
	defer done()

	data, err := h.userSrv.SetTradeFee(ctxWT, msg, fee)
	if err != nil {
		log.Printf("FAILED_UPDATE_USER_DATA: %v", err)
		return
	}
	h.reply(data)
}

func (h *Handler) EditConfirmTradeState(msg *tgbotapi.Message) {
	h.notifyProcessingUpdate(msg)

	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
	defer done()

	data, err := h.userSrv.SetConfirmTrade(ctxWT, msg)
	if err != nil {
		log.Printf("FAILED_UPDATE_USER_DATA: %v", err)
		return
	}
	h.reply(data)
}

func (h *Handler) EditBuyAmountState(cmd string, msg *tgbotapi.Message, pos int) {
	stateKey := fmt.Sprintf("%d_state", msg.Chat.ID)
	messageKey := fmt.Sprintf("%d_msg", msg.Chat.ID)
	itemPosKey := fmt.Sprintf("%d_pos", msg.Chat.ID)
	if cmd != common.AwaitingBuyAmount {
		h.cache.Set(itemPosKey, pos, CacheDuration*time.Minute)
		h.promptForUpdate(stateKey, common.AwaitingBuyAmount, messageKey, msg)
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
	// get update pos
	var buyAmountPos int
	if cacheData, ok := h.cache.Get(itemPosKey); ok {
		if id, ok := cacheData.(int); ok {
			buyAmountPos = id
		}
	}

	// sent fallback for update
	npMsg := msg
	npMsg.MessageID = prevMessageID
	h.notifyProcessingUpdate(npMsg)

	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
	defer done()

	data, err := h.userSrv.SetBuyAmount(ctxWT, prevMessageID, buyAmountPos, msg)
	if err != nil {
		log.Printf("failed update user data: %v", err)
		return
	}

	// remove state
	h.cache.Delete(stateKey)
	h.cache.Delete(messageKey)
	h.cache.Delete(itemPosKey)
	h.reply(data)
}

func (h *Handler) EditBuySlippageState(cmd string, msg *tgbotapi.Message) {
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
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
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

func (h *Handler) EditSellAmountState(cmd string, msg *tgbotapi.Message, pos int) {
	stateKey := fmt.Sprintf("%d_state", msg.Chat.ID)
	messageKey := fmt.Sprintf("%d_msg", msg.Chat.ID)
	itemPosKey := fmt.Sprintf("%d_pos", msg.Chat.ID)
	if cmd != common.AwaitingSellAmount {
		h.cache.Set(itemPosKey, pos, CacheDuration*time.Minute)
		h.promptForUpdate(stateKey, common.AwaitingSellAmount, messageKey, msg)
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
	// get update pos
	var sellAmountPos int
	if cacheData, ok := h.cache.Get(itemPosKey); ok {
		if id, ok := cacheData.(int); ok {
			sellAmountPos = id
		}
	}

	// sent fallback for update
	npMsg := msg
	npMsg.MessageID = prevMessageID
	h.notifyProcessingUpdate(npMsg)

	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
	defer done()

	data, err := h.userSrv.SetSellAmount(ctxWT, prevMessageID, sellAmountPos, msg)
	if err != nil {
		log.Printf("failed update user data: %v", err)
		return
	}

	// remove state
	h.cache.Delete(stateKey)
	h.cache.Delete(messageKey)
	h.cache.Delete(itemPosKey)
	h.reply(data)
}

func (h *Handler) EditSellSlippageState(cmd string, msg *tgbotapi.Message) {
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
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
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

func (h *Handler) EditSellProtectionState(msg *tgbotapi.Message) {
	h.notifyProcessingUpdate(msg)

	ctx := context.Background()
	ctxWT, done := context.WithTimeout(ctx,
		ContextDuration*time.Second)
	defer done()

	data, err := h.userSrv.SetSellProtection(ctxWT, msg)
	if err != nil {
		log.Printf("failed update user data: %v", err)
		return
	}
	h.reply(data)
}
