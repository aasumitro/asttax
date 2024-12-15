package keyboard

import (
	"fmt"

	"github.com/aasumitro/asttax/internal/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TrenchesKeyboardMarkup(state string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(
					"%s%s", func() string {
						if state == common.TrenchesNewPairs {
							return common.CheckmarkEmoticon + " "
						}
						return ""
					}(), common.PlantEmoticon), common.TrenchesNewPairs),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(
					"%s%s", func() string {
						if state == common.TrenchesIgnitingEngines {
							return common.CheckmarkEmoticon + " "
						}
						return ""
					}(), common.FireEmoticon), common.TrenchesIgnitingEngines),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(
					"%s%s", func() string {
						if state == common.TrenchesGraduated {
							return common.CheckmarkEmoticon + " "
						}
						return ""
					}(), common.RocketEmoticon), common.TrenchesGraduated),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s Back",
					common.BackwardEmoticon), "back_to_start"),
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s Refresh",
					common.RefreshEmoticon), "refresh_trances"),
			},
		},
	}
}
