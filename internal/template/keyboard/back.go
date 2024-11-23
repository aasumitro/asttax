package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var BackToStartKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("⬅️ Back", "back_to_start")},
	},
}
