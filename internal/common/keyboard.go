package common

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var StartReplyMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Buy", "buy"),
			tgbotapi.NewInlineKeyboardButtonData("Sell", "sell"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("Positions", "positions"),
			tgbotapi.NewInlineKeyboardButtonData("Limit Orders  🔜", "limit_orders"),
			tgbotapi.NewInlineKeyboardButtonData("DCA Orders  🔜", "dca_orders"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("New Pairs", "new_pairs"),
			tgbotapi.NewInlineKeyboardButtonData("Copy Trade  🔜", "copy_trade"),
			tgbotapi.NewInlineKeyboardButtonData("Sniper  🔜", "sniper"),
		},
		{tgbotapi.NewInlineKeyboardButtonData("🔄  Refresh", "refresh")},
		{
			tgbotapi.NewInlineKeyboardButtonData("⚙️  Settings", "settings"),
			tgbotapi.NewInlineKeyboardButtonData("📖  Help", "help"),
		},
	},
}

var BackReplyMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("⬅️ Back", "back_to_start")},
	},
}
