package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var SettingKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("⬅️  Back", "back_to_start"),
			tgbotapi.NewInlineKeyboardButtonData("🏳️English  ➡️", "language_setting"),
		},
		{tgbotapi.NewInlineKeyboardButtonData("——— Trade Fees ———", "none")},
		{
			tgbotapi.NewInlineKeyboardButtonData("✅  Fast (0.0015 SOL)  🐴", "none"),
			tgbotapi.NewInlineKeyboardButtonData("Turbo (0.0075 SOL)  🚀", "none"),
			tgbotapi.NewInlineKeyboardButtonData("Custom Fee (0)  🔥", "none"),
		},
		{tgbotapi.NewInlineKeyboardButtonData("——— Trade Protections ———", "none")},
		{
			tgbotapi.NewInlineKeyboardButtonData("🔴  MEV Protect (Buys)", "none"),
			tgbotapi.NewInlineKeyboardButtonData("🔴  MEV Protect (Sells)", "none"),
			tgbotapi.NewInlineKeyboardButtonData("🔴  Confirm Trade", "none"),
		},
		{tgbotapi.NewInlineKeyboardButtonData("——— Buy Amounts ———", "none")},
		{
			tgbotapi.NewInlineKeyboardButtonData("0.25 SOL  ✏️", "none"),
			tgbotapi.NewInlineKeyboardButtonData("0.5 SOL ✏️", "none"),
			tgbotapi.NewInlineKeyboardButtonData("1 SOL ✏️", "none"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("2.5 SOL ✏️", "none"),
			tgbotapi.NewInlineKeyboardButtonData("5 SOL ✏️", "none"),
			tgbotapi.NewInlineKeyboardButtonData("10 SOL ✏️", "none"),
		},
		{tgbotapi.NewInlineKeyboardButtonData("Buy Slippage: 15%  ✏️", "none")},
		{tgbotapi.NewInlineKeyboardButtonData("——— Sell Amounts ———", "none")},
		{
			tgbotapi.NewInlineKeyboardButtonData("25%  ✏️", "none"),
			tgbotapi.NewInlineKeyboardButtonData("50%  ✏️", "none"),
			tgbotapi.NewInlineKeyboardButtonData("100%  ✏️", "none"),
		},
		{tgbotapi.NewInlineKeyboardButtonData("Sell Slippage: 15%  ✏️", "none")},
		{tgbotapi.NewInlineKeyboardButtonData("Sell Protection  🟢️", "none")},
		{tgbotapi.NewInlineKeyboardButtonData("——— Account ———", "none")},
		{
			tgbotapi.NewInlineKeyboardButtonData("Wallet", "none"),
		},
	},
}

var LanguageSettingKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("✅  English", "none"),
			tgbotapi.NewInlineKeyboardButtonData("Bahasa Indonesia", "none"),
		},
		{tgbotapi.NewInlineKeyboardButtonData("⬅️ Back", "back_to_setting")},
	},
}
