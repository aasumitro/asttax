package keyboard

import (
	"fmt"

	"github.com/aasumitro/asttax/internal/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var AgreementKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s Accept Agreement",
			common.CheckmarkEmoticon), "accept_agreement")},
	},
}

var AccountCreatedKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("Start Trading", "start")},
	},
}
