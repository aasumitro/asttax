package common

/**
* Command
* start - Trade on Solana with AsttaX
* buy - Buy a Token
* sell - Sell a Token
* positions - View Detailed information about your hodls
* settings - Configure your settings
* withdraw - Withdraw tokens (SOL)
* help - FAQ and Telegram Channel
* backup - Backup bots in case of log and issues
 */

const (
	AcceptAgreement  = "accept_agreement"
	Start            = "start"
	Buy              = "buy"
	Sell             = "sell"
	Positions        = "positions"
	Settings         = "settings"
	LanguageSettings = "language_setting"
	Withdraw         = "withdraw"
	Help             = "help"
	Backup           = "backup"
	Refresh          = "refresh"
	NewPairs         = "new_pairs"
	BackToStart      = "back_to_start"
	BackToSetting    = "back_to_setting"
)

const (
	MessageParseMarkdown = "Markdown"
	MessageParseHTML     = "HTML"
)
