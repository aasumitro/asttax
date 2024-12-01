package keyboard

import (
	"fmt"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func LoadSettingKeyboardMarkup(user *model.User) tgbotapi.InlineKeyboardMarkup {
	if user == nil {
		return tgbotapi.InlineKeyboardMarkup{}
	}
	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("⬅️  Back", "back_to_start"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("🏳️%s  ➡️", func() string {
					if user.BotLang == "id" {
						return "Indonesia"
					}
					return "English"
				}()), "language_setting"),
			},
			{tgbotapi.NewInlineKeyboardButtonData("——— Trade Fees ———", "none")},
			{
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%sFast (0.0015 SOL)  🐴", func() string {
					if user.TradeFees != "fast" {
						return ""
					}
					return common.CheckmarkEmoticon
				}()), "none"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%sTurbo (0.0075 SOL)  🚀", func() string {
					if user.TradeFees != "turbo" {
						return ""
					}
					return common.CheckmarkEmoticon
				}()), "none"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%sCustom Fee (%.2f)  🔥", func() string {
					if user.TradeFees != "custom" {
						return ""
					}
					return common.CheckmarkEmoticon
				}(), func() float64 {
					if user.TradeFees != "custom" || user.CustomTradeFee == 0 {
						return 0
					}
					return user.CustomTradeFee
				}()), "none"),
			},
			{tgbotapi.NewInlineKeyboardButtonData("——— Trade Protections ———", "none")},
			{
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s  MEV Protect (Buys)", func() string {
					if user.MEVBuyProtection {
						return common.EnabledEmoticon
					}
					return common.DisabledEmoticon
				}()), "none"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s  MEV Protect (Sells)", func() string {
					if user.MEVSellProtection {
						return common.EnabledEmoticon
					}
					return common.DisabledEmoticon
				}()), "none"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s  Confirm Trade", func() string {
					if user.ConfirmTradeProtection {
						return common.EnabledEmoticon
					}
					return common.DisabledEmoticon
				}()), "none"),
			},
			{tgbotapi.NewInlineKeyboardButtonData("——— Buy Amounts ———", "none")},
			{
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%.2f SOL ✏️", user.BuyAmountP1), "none"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%.2f SOL ✏️", user.BuyAmountP2), "none"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%.2f SOL ✏️", user.BuyAmountP3), "none"),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%.2f SOL ✏️", user.BuyAmountP4), "none"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%.2f SOL ✏️", user.BuyAmountP5), "none"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%.2f SOL ✏️", user.BuyAmountP6), "none"),
			},
			{tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("Buy Slippage: %.2f%%  ✏️", user.BuySlippage), "none")},
			{tgbotapi.NewInlineKeyboardButtonData("——— Sell Amounts ———", "none")},
			{
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%.2f%%  ✏️", user.SellAmountP1), "none"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%.2f%%  ✏️", user.SellAmountP2), "none"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%.2f%%  ✏️", user.SellAmountP3), "none"),
			},
			{tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("Sell Slippage: %.2f%%  ✏️", user.SellSlippage), "none")},
			{tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("Sell Protection  %s", func() string {
				if user.SellProtection {
					return common.EnabledEmoticon
				}
				return common.DisabledEmoticon
			}()), "none")},
			{tgbotapi.NewInlineKeyboardButtonData("——— Account ———", "none")},
			{tgbotapi.NewInlineKeyboardButtonData("Wallet", "none")},
		},
	}
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
