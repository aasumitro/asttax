package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var AgreementKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("✅ Accept Agreement", "accept_agreement")},
	},
}

var AccountCreatedKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("Start Trading", "start")},
	},
}
