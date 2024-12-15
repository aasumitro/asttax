package keyboard

import (
	"fmt"

	"github.com/aasumitro/asttax/internal/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var StartKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Buy", "buy"),
			tgbotapi.NewInlineKeyboardButtonData("Sell", "sell"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("Trenches", "trenches"),
			tgbotapi.NewInlineKeyboardButtonData("Positions", "positions"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s  Settings",
				common.SettingEmoticon), "settings"),
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s  Help",
				common.HelpEmoticon), "help"),
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s  Refresh",
				common.RefreshEmoticon), "refresh"),
		},
	},
}
