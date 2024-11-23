package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aasumitro/asttax/internal/common"
	sqlrepo "github.com/aasumitro/asttax/internal/repository/sql"
	"github.com/aasumitro/asttax/internal/template/keyboard"
	"github.com/aasumitro/asttax/internal/template/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IUserService interface {
	LoadUser(ctx context.Context, msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
}

type userService struct {
	userRepo sqlrepo.IUserRepository
}

func (srv *userService) LoadUser(
	ctx context.Context,
	msg *tgbotapi.Message,
) (*tgbotapi.MessageConfig, error) {
	user, err := srv.userRepo.Find(ctx, msg.Chat.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	var reply tgbotapi.MessageConfig

	if user == nil || errors.Is(err, sql.ErrNoRows) || !user.AcceptAgreement {
		reply = tgbotapi.NewMessage(msg.Chat.ID, message.AgreementTextBody())
		reply.ParseMode = common.MessageParseHTML
		reply.ReplyMarkup = keyboard.AgreementKeyboardMarkup
	}

	if user != nil && user.AcceptAgreement {
		reply = tgbotapi.NewMessage(msg.Chat.ID, message.StartTextBody())
		reply.ParseMode = common.MessageParseMarkdown
		reply.ReplyMarkup = keyboard.StartKeyboardMarkup
	}

	reply.ReplyToMessageID = msg.MessageID

	return &reply, nil
}

func NewUserService(userRepo sqlrepo.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}
