package service

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IUserService interface {
	LoadUser(ctx context.Context, msg *tgbotapi.Message, cmd, raw bool) (interface{}, error)
	LoadUserSetting(ctx context.Context, msg *tgbotapi.Message, cmd bool) (interface{}, error)
	CreateUser(ctx context.Context, msg *tgbotapi.Message) (*tgbotapi.EditMessageTextConfig, error)
	SetTradeFee(ctx context.Context, msg *tgbotapi.Message, item string) (interface{}, error)
	SetConfirmTrade(ctx context.Context, msg *tgbotapi.Message) (interface{}, error)
	SetBuyAmount(ctx context.Context, prevMsgID, itemPos int, msg *tgbotapi.Message) (interface{}, error)
	SetBuySlippage(ctx context.Context, prevMsgID int, msg *tgbotapi.Message) (interface{}, error)
	SetSellAmount(ctx context.Context, prevMsgID, itemPos int, msg *tgbotapi.Message) (interface{}, error)
	SetSellSlippage(ctx context.Context, prevMsgID int, msg *tgbotapi.Message) (interface{}, error)
	SetSellProtection(ctx context.Context, msg *tgbotapi.Message) (interface{}, error)
}

type ITransactionService interface {
	// Sell
	// Buy
	// Pairs
	// Position
	// Withdraw
}
