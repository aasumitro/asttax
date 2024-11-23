package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var StartKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Buy", "buy"),
			tgbotapi.NewInlineKeyboardButtonData("Sell", "sell"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("New Pairs", "new_pairs"),
			tgbotapi.NewInlineKeyboardButtonData("Positions", "positions"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("⚙️  Settings", "settings"),
			tgbotapi.NewInlineKeyboardButtonData("📖  Help", "help"),
			tgbotapi.NewInlineKeyboardButtonData("🔄  Refresh", "refresh"),
		},
	},
}
