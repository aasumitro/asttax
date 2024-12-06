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

	lang := "English"
	if user.BotLang == "id" {
		lang = "Indonesia"
	}

	tradeFeeFast := ""
	if user.TradeFees == "fast" {
		tradeFeeFast = common.CheckmarkEmoticon
	}

	tradeFeeTurbo := ""
	if user.TradeFees == "turbo" {
		tradeFeeTurbo = common.CheckmarkEmoticon
	}

	confirmTrade := common.DisabledEmoticon
	if user.ConfirmTradeProtection {
		confirmTrade = common.EnabledEmoticon
	}

	sellProtection := common.DisabledEmoticon
	if user.SellProtection {
		sellProtection = common.EnabledEmoticon
	}

	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%s  Back", common.BackwardEmoticon), "back_to_start"),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("🏳️%s  %s", lang, common.ForwardEmoticon), "language_setting"),
			},
			{tgbotapi.NewInlineKeyboardButtonData("——— Trade Fees ———", "none")},
			{tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%sFast (0.0015 SOL)  %s",
					tradeFeeFast, common.HorseEmoticon), common.FastTradeFee)},
			{tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%sTurbo (0.0075 SOL)  %s",
					tradeFeeTurbo, common.RocketEmoticon), common.TurboTradeFee)},
			{tgbotapi.NewInlineKeyboardButtonData("——— Trade Protections ———", "none")},
			{tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s  Confirm Trade", confirmTrade), common.ConfirmTrade)},
			{tgbotapi.NewInlineKeyboardButtonData("——— Buy Amounts ———", "none")},
			{
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%.2f SOL %s", user.BuyAmountP1,
						common.EditEmoticon), common.BuyAmountP1),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%.2f SOL %s", user.BuyAmountP2,
						common.EditEmoticon), common.BuyAmountP2),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%.2f SOL %s", user.BuyAmountP3,
						common.EditEmoticon), common.BuyAmountP3),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%.2f SOL %s", user.BuyAmountP4,
						common.EditEmoticon), common.BuyAmountP4),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%.2f SOL %s", user.BuyAmountP5,
						common.EditEmoticon), common.BuyAmountP5),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%.2f SOL %s", user.BuyAmountP6,
						common.EditEmoticon), common.BuyAmountP6),
			},
			{tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("Buy Slippage: %.2f%%  %s", user.BuySlippage,
					common.EditEmoticon), common.BuySlippage)},
			{tgbotapi.NewInlineKeyboardButtonData("——— Sell Amounts ———", "none")},
			{
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%.2f%%  %s", user.SellAmountP1,
						common.EditEmoticon), common.SellAmountP1),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%.2f%%  %s", user.SellAmountP2,
						common.EditEmoticon), common.SellAmountP2),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%.2f%%  %s", user.SellAmountP3,
						common.EditEmoticon), common.SellAmountP3),
			},
			{tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("Sell Slippage: %.2f%%  %s", user.SellSlippage,
					common.EditEmoticon), common.SellSlippage)},
			{tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("Sell Protection  %s", sellProtection), common.SellProtection)},
			{tgbotapi.NewInlineKeyboardButtonData("——— Account ———", "none")},
			{tgbotapi.NewInlineKeyboardButtonData("Wallet", "none")},
		},
	}
}

var LanguageSettingKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s English",
				common.CheckmarkEmoticon), "none"),
			// tgbotapi.NewInlineKeyboardButtonData("Bahasa Indonesia", "none"),
		},
		{tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s Back",
			common.BackwardEmoticon), "back_to_setting")},
	},
}

var AfterUpdateSettingKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s Back",
			common.BackwardEmoticon), "back_to_setting")},
	},
}
