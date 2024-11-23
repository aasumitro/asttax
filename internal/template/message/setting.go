package message

func SettingTextBody() string {
	return `
<b>FAQ:</b>

🚀 <b>Fast/Turbo:</b> Set your preferred priority fee to decrease likelihood of failed transactions.

🛡️ <b>MEV Protection:</b>
Enable this setting to send transactions privately and avoid getting frontrun or sandwiched.
<b><u>Important Note</u></b>: If you enable MEV Protection your transactions may take longer to get confirmed.

🔴 <b>Confirm Trades: Red = off</b>, clicking on the amount of SOL to purchase or setting a custom amount will instantly initiate the transaction.

🟢 <b>Confirm Trades: Green = on</b>, you will need to confirm your intention to swap by clicking the Buy or Sell buttons.

🟢 <b>Sell Protection: Green = on</b>, you will need to confirm your intention when selling more than 75% of your token balance.
`
}
