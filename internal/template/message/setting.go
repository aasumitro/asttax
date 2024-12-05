package message

import (
	"fmt"

	"github.com/aasumitro/asttax/internal/common"
)

func SettingTextBody() string {
	return fmt.Sprintf(`
<b>FAQ:</b>

%s <b>Fast/Turbo:</b> Set your preferred priority fee to decrease likelihood of failed transactions.

%s <b>Confirm Trades: Red = off</b>, clicking on the amount of SOL to purchase or setting a custom amount will instantly initiate the transaction.

%s <b>Confirm Trades: Green = on</b>, you will need to confirm your intention to swap by clicking the Buy or Sell buttons.

%[3]s <b>Sell Protection: Green = on</b>, you will need to confirm your intention when selling more than 75%% of your token balance.
`, common.RocketEmoticon, common.DisabledEmoticon, common.EnabledEmoticon)
}
