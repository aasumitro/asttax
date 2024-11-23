package message

func StartTextBody() string {
	solanaAddress := "3x6QDiKyZR4vDjtJGXyXcEsQyh4CX2QoxyLjjhVFkqCG"
	solanaURL := "https://solscan.io/account/" + solanaAddress
	solanaValue := "0 SOL ($0.00)"
	return "Welcome to AsttaX on Solana! \n\n" +
		" *Solana* • [🅴](" + solanaURL + ")\n" +
		"💰`" + solanaAddress + "` _(Tap to copy)_\n" +
		"🪙Balance: `" + solanaValue + "`\n\n" +
		"Click on the buttons below to navigate:\n" +
		"• Use the Buy/Sell buttons to trade.\n" +
		"• Check your positions, and New available Pairs!\n" +
		"• Use the Refresh button to update your current balance.\n\n"
}
