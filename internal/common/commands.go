package common

type Command string

const (
	Start     Command = "start"     // start - Trade on Solana with AsttaX
	Buy       Command = "buy"       // buy - Buy a Token
	Sell      Command = "sell"      // sell - Sell a Token
	Positions Command = "positions" // positions - View Detailed information about your hodls
	Settings  Command = "settings"  // settings - Configure your settings
	Withdraw  Command = "withdraw"  // withdraw - Withdraw tokens (SOL)
	Help      Command = "help"      // help - FAQ and Telegram Channel
	Backup    Command = "backup"    // backup - Backup bots in case of log and issues
)
