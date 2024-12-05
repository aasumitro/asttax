package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/model"
	restRepo "github.com/aasumitro/asttax/internal/repository/rest"
	rpcRepo "github.com/aasumitro/asttax/internal/repository/rpc"
	sqlrepo "github.com/aasumitro/asttax/internal/repository/sql"
	"github.com/aasumitro/asttax/internal/template/keyboard"
	"github.com/aasumitro/asttax/internal/template/message"
	"github.com/aasumitro/asttax/internal/util"
	"github.com/blocto/solana-go-sdk/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mr-tron/base58"
	"golang.org/x/sync/errgroup"
)

type IUserService interface {
	LoadUser(ctx context.Context, msg *tgbotapi.Message, cmd, raw bool) (interface{}, error)
	LoadUserSetting(ctx context.Context, msg *tgbotapi.Message, cmd bool) (interface{}, error)
	CreateUser(ctx context.Context, msg *tgbotapi.Message) (*tgbotapi.EditMessageTextConfig, error)
	SetTradeFee(ctx context.Context, msg *tgbotapi.Message, item string) (interface{}, error)
	SetConfirmTrade(ctx context.Context, msg *tgbotapi.Message) (interface{}, error)
	SetSellProtection(ctx context.Context, msg *tgbotapi.Message) (interface{}, error)
}

type userService struct {
	userRepo  sqlrepo.IUserRepository
	restRepo  restRepo.ICoingeckoRepository
	rpcRepo   rpcRepo.ISolanaRPCRepository
	secretKey string
}

func (srv *userService) LoadUser(
	ctx context.Context,
	msg *tgbotapi.Message,
	cmd, raw bool,
) (interface{}, error) {
	// get user data form database
	user, err := srv.userRepo.Find(ctx, msg.Chat.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	// If user does not exist or hasn't accepted the agreement
	if user == nil || errors.Is(err, sql.ErrNoRows) || !user.AcceptAgreement {
		return srv.createAgreementMessage(msg), nil
	}
	// if only requested account data
	if raw {
		return user, nil
	}
	// If user exists and has accepted the agreement
	return srv.createStartMessage(ctx, msg, user, cmd)
}

func (srv *userService) createAgreementMessage(
	msg *tgbotapi.Message,
) interface{} {
	reply := tgbotapi.NewMessage(msg.Chat.ID, message.AgreementTextBody())
	reply.ParseMode = common.MessageParseHTML
	reply.ReplyMarkup = keyboard.AgreementKeyboardMarkup
	reply.ReplyToMessageID = msg.MessageID
	return &reply
}

func (srv *userService) createStartMessage(
	ctx context.Context,
	msg *tgbotapi.Message,
	user *model.User,
	cmd bool,
) (interface{}, error) {
	// get user balance
	balanceSOL, balanceUSD, err := srv.getUserBalance(ctx, user.WalletAddress)
	if err != nil {
		return nil, err
	}
	// if a user presses the menu button
	msgTxt := message.StartTextBody(user.WalletAddress, balanceSOL, balanceUSD)
	if !cmd {
		reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
			msg.MessageID, msgTxt, keyboard.StartKeyboardMarkup)
		reply.ParseMode = common.MessageParseMarkdown
		return &reply, nil
	}
	// if a user uses the main command
	reply := tgbotapi.NewMessage(msg.Chat.ID, msgTxt)
	reply.ParseMode = common.MessageParseMarkdown
	reply.ReplyMarkup = keyboard.StartKeyboardMarkup
	reply.ReplyToMessageID = msg.MessageID
	return &reply, nil
}

func (srv *userService) LoadUserSetting(
	ctx context.Context,
	msg *tgbotapi.Message,
	cmd bool,
) (interface{}, error) {
	// load user data from database
	user, err := srv.LoadUser(ctx, msg, true, true)
	if err != nil {
		return nil, err
	}
	// build keyboard menus
	u := user.(*model.User)
	replyKeyboard := keyboard.LoadSettingKeyboardMarkup(u)
	// if requested from command, reply with this item
	if cmd {
		reply := tgbotapi.NewMessage(msg.Chat.ID, message.SettingTextBody())
		reply.ParseMode = common.MessageParseHTML
		reply.ReplyMarkup = replyKeyboard
		reply.ReplyToMessageID = msg.MessageID
		return &reply, nil
	}
	// if request not from command (from a menu) replies with this item
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.SettingTextBody(), replyKeyboard)
	reply.ParseMode = common.MessageParseHTML
	return &reply, nil
}

func (srv *userService) CreateUser(
	ctx context.Context,
	msg *tgbotapi.Message,
) (*tgbotapi.EditMessageTextConfig, error) {
	// init base user data
	user := &model.User{AcceptAgreement: true}
	user.TelegramID = msg.Chat.ID
	// create solana wallet
	account := types.NewAccount()
	privateKey := base58.Encode(account.PrivateKey)
	secretKey := string(util.NormalizeKey(srv.secretKey))
	encryptedPrivateKey, err := util.Encrypt(privateKey, secretKey)
	if err != nil {
		return nil, err
	}
	user.WalletAddress = account.PublicKey.ToBase58()
	user.PrivateKey = encryptedPrivateKey
	// store user data
	if _, err := srv.userRepo.Insert(ctx, user); err != nil {
		return nil, err
	}
	// build a reply message
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID, msg.MessageID,
		message.AccountCreatedTextBody(user.WalletAddress, privateKey),
		keyboard.AccountCreatedKeyboardMarkup)
	reply.ParseMode = common.MessageParseHTML
	return &reply, nil
}

func (srv *userService) getUserBalance(
	ctx context.Context,
	walletAddress string,
) (float64, float64, error) {
	var err error
	var balance uint64
	var solPrice float64
	// get user balance and current solana price
	eg, ctxEG := errgroup.WithContext(ctx)
	eg.Go(func() error {
		balance, err = srv.rpcRepo.GetBalance(ctxEG, walletAddress)
		return err
	})
	eg.Go(func() error {
		solPrice, err = srv.restRepo.GetSolanaPrice(ctxEG)
		return err
	})
	if err := eg.Wait(); err != nil {
		return 0, 0, err
	}
	// calculate and convert solana balance into usd
	currUnit := 1e9
	balanceSol := float64(balance) / currUnit
	balanceUsd := balanceSol * solPrice
	return balanceSol, balanceUsd, nil
}

func (srv *userService) SetTradeFee(
	ctx context.Context,
	msg *tgbotapi.Message,
	item string,
) (interface{}, error) {
	data, err := srv.userRepo.Update(ctx, "trade_fees",
		item, msg.Chat.ID)
	if err != nil {
		return nil, err
	}
	return srv.settingMessageReply(data, msg)
}

func (srv *userService) SetConfirmTrade(
	ctx context.Context,
	msg *tgbotapi.Message,
) (interface{}, error) {
	data, err := srv.userRepo.Update(ctx, "confirm_trade_protection",
		"NOT confirm_trade_protection", msg.Chat.ID)
	if err != nil {
		return nil, err
	}
	return srv.settingMessageReply(data, msg)
}

func (srv *userService) SetSellProtection(
	ctx context.Context,
	msg *tgbotapi.Message,
) (interface{}, error) {
	data, err := srv.userRepo.Update(ctx, "sell_protection",
		"NOT sell_protection", msg.Chat.ID)
	if err != nil {
		return nil, err
	}
	return srv.settingMessageReply(data, msg)
}

func (srv *userService) settingMessageReply(
	user *model.User,
	msg *tgbotapi.Message,
) (interface{}, error) {
	replyKeyboard := keyboard.LoadSettingKeyboardMarkup(user)
	reply := tgbotapi.NewEditMessageTextAndMarkup(msg.Chat.ID,
		msg.MessageID, message.SettingTextBody(), replyKeyboard)
	reply.ParseMode = common.MessageParseHTML
	return &reply, nil
}

func NewUserService(
	userRepo sqlrepo.IUserRepository,
	restRepo restRepo.ICoingeckoRepository,
	rpcRepo rpcRepo.ISolanaRPCRepository,
	secretKey string,
) IUserService {
	return &userService{
		userRepo:  userRepo,
		restRepo:  restRepo,
		rpcRepo:   rpcRepo,
		secretKey: secretKey,
	}
}
