package message

import "fmt"

func AgreementTextBody() string {
	return `
<b><u>AsttaX Bot User Agreement</u></b>

<b>1. Introduction</b>

Welcome to AsttaX, a Telegram-based trading bot designed for Solana blockchain transactions. By using the AsttaX bot (“the Bot”), you agree to comply with and be bound by the following terms and conditions (“User Agreement”). Please read this agreement carefully before using the Bot.

<b>2. Acceptance of Terms</b>

By interacting with the Bot, you confirm that:
	•	You are at least 18 years old or have reached the legal age of majority in your jurisdiction.
	•	You have reviewed and agreed to these terms.
	•	You comply with all applicable laws and regulations related to blockchain technology, cryptocurrency, and trading in your jurisdiction.

<b>3. Services</b>

The Bot provides the following functionalities:
	•	Creation of Solana-compatible cryptocurrency wallets.
	•	Trading functionality through integrations with third-party platforms such as Jupiter and Raydium.

<b>4. Disclaimer of Responsibility</b>

	•	<b>No Financial Advice:</b> The Bot does not provide financial, investment, or trading advice. Use of the Bot is at your own risk.
	•	<b>Third-Party Integrations:</b> The Bot relies on third-party platforms (e.g., Jupiter and Raydium) for trading. Astatx is not responsible for any failures, delays, or losses incurred due to issues with these platforms.

<b>5. Risks</b>

By using the Bot, you acknowledge and agree that:
	•	Cryptocurrency trading involves significant financial risks, including market volatility and the potential for loss.
	•	You are solely responsible for understanding and managing these risks.

<b>6. Data and Privacy</b>

	•	The Bot may collect and process limited personal information as required to provide its services (e.g., Telegram usernames, id, etc.).
	•	Wallets created via the Bot are private, and the Bot does not store private keys. Users are solely responsible for safeguarding their private keys and recovery phrases.

<b>7. Limitation of Liability</b>

AsttaX and its developers are not liable for any:
	•	Direct, indirect, or incidental damages arising from the use of the Bot.
	•	Loss of funds due to user errors, technical failures, or external factors such as hacks.

<b>8. Termination</b>

The Bot’s services may be suspended or terminated at any time, without prior notice, if you violate this agreement or if operational changes require discontinuation of services.

<b>9. Amendments</b>

AsttaX reserves the right to update this User Agreement at any time. Continued use of the Bot after changes are published constitutes acceptance of the updated terms.

<b>10. Contact</b>

For support or inquiries, please contact us at hello@astta.xyz.
`
}

func ConfirmAgreementCallbackTextBody() string {
	return `
Welcome to AsttaX. AsttaX enables you to quickly buy or sell tokens and set automations like Limit Orders, DCA, Copy-trading and Sniping.

✅ You've accept our terms, you can now use AsttaX

🟠 <i>Creating AsttaX account and solana wallet please wait . . . . .</i>
`
}

func AccountCreatedTextBody(walletAddress, secretKey string) string {
	return fmt.Sprintf(`
🟢 <i>Generated your account & wallet</i>

SOL: %s
<u><b>Secret Key:</b></u>
%s

BE SURE TO RETAIN THE INFORMATION ABOVE IN A SAFE PLACE.
THIS MESSAGE WILL AUTO-DELETE AND NOT BE AVAILABLE IN YOUR CHAT HISTORY AFTER 5 MINUTES.
`, walletAddress, secretKey)
}
