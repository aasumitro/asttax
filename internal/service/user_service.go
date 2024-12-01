package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/model"
	sqlrepo "github.com/aasumitro/asttax/internal/repository/sql"
	"github.com/aasumitro/asttax/internal/template/keyboard"
	"github.com/aasumitro/asttax/internal/template/message"
	"github.com/aasumitro/asttax/internal/util"
	"github.com/blocto/solana-go-sdk/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mr-tron/base58"
)

type IUserService interface {
	LoadUser(ctx context.Context, msg *tgbotapi.Message, cmd, raw bool) (interface{}, error)
	LoadUserSetting(ctx context.Context, msg *tgbotapi.Message, cmd bool) (interface{}, error)
	CreateUser(ctx context.Context, msg *tgbotapi.Message) (*tgbotapi.EditMessageTextConfig, error)
}

type userService struct {
	userRepo  sqlrepo.IUserRepository
	secretKey string
}

func (srv *userService) LoadUser(
	ctx context.Context,
	msg *tgbotapi.Message,
	cmd, raw bool,
) (interface{}, error) {
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
	return srv.createStartMessage(msg, user, cmd), nil
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
	msg *tgbotapi.Message,
	user *model.User,
	cmd bool,
) interface{} {
	// if a user presses the menu button
	if !cmd {
		reply := tgbotapi.NewEditMessageTextAndMarkup(
			msg.Chat.ID, msg.MessageID,
			message.StartTextBody(user.WalletAddress),
			keyboard.StartKeyboardMarkup)
		reply.ParseMode = common.MessageParseMarkdown
		return &reply
	}
	// if a user uses the main command
	reply := tgbotapi.NewMessage(msg.Chat.ID, message.StartTextBody(user.WalletAddress))
	reply.ParseMode = common.MessageParseMarkdown
	reply.ReplyMarkup = keyboard.StartKeyboardMarkup
	reply.ReplyToMessageID = msg.MessageID
	return &reply
}

func (srv *userService) LoadUserSetting(
	ctx context.Context,
	msg *tgbotapi.Message,
	cmd bool,
) (interface{}, error) {
	user, err := srv.LoadUser(ctx, msg, true, true)
	if err != nil {
		return nil, err
	}

	u := user.(*model.User)
	replyKeyboard := keyboard.LoadSettingKeyboardMarkup(u)

	if cmd {
		reply := tgbotapi.NewMessage(msg.Chat.ID, message.SettingTextBody())
		reply.ParseMode = common.MessageParseHTML
		reply.ReplyMarkup = replyKeyboard
		reply.ReplyToMessageID = msg.MessageID
		return &reply, nil
	}

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

func NewUserService(
	userRepo sqlrepo.IUserRepository,
	secretKey string,
) IUserService {
	return &userService{
		userRepo:  userRepo,
		secretKey: secretKey,
	}
}
