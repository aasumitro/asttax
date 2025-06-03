package keyboard

import (
	"fmt"

	"github.com/aasumitro/asttax/internal/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var BackToStartKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s Back",
			common.BackwardEmoticon), "back_to_start")},
	},
}
