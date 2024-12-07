package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var ConfirmTradeKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("No, go back!", "back_to_start"),
			tgbotapi.NewInlineKeyboardButtonData("Yes, proceed!", "none"),
		},
	},
}
