package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var PositionKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Back", "back_to_start"),
			tgbotapi.NewInlineKeyboardButtonData("🔄 Refresh", "refresh_position"),
		},
	},
}
