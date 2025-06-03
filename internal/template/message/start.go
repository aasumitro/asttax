package message

import (
	"fmt"

	"github.com/aasumitro/asttax/internal/common"
)

func StartTextBody(solanaAddress string, balanceSOL, balanceUSD float64) string {
	solanaURL := "https://solscan.io/account/" + solanaAddress
	solanaBalance := fmt.Sprintf("%.2f SOL (%.2f USD)", balanceSOL, balanceUSD)
	return "Welcome to AsttaX on Solana! \n\n" +
		common.BagEmoticon + " *Solana* • [🅴](" + solanaURL + ")\n" +
		"`" + solanaAddress + "` _(Tap to copy)_\n" +
		common.CoinEmoticon + " Balance: `" + solanaBalance + "`\n\n" +
		"Click on the buttons below to navigate:\n" +
		"• Use the Buy/Sell buttons to trade.\n" +
		"• Check your positions, and New available Pairs!\n" +
		"• Use the Refresh button to update your current balance.\n\n"
}
