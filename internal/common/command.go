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
	Trenches         = "trenches"
	BackToStart      = "back_to_start"
	BackToSetting    = "back_to_setting"

	TrenchesNewPairs        = "trenches_new_pairs"
	TrenchesIgnitingEngines = "trenches_igniting_engines"
	TrenchesGraduated       = "trenches_graduated"

	FastTradeFee   = "fast_trade_fee"
	TurboTradeFee  = "turbo_trade_fee"
	ConfirmTrade   = "confirm_trade"
	BuyAmountP1    = "buy_amount_p1"
	BuyAmountP2    = "buy_amount_p2"
	BuyAmountP3    = "buy_amount_p3"
	BuyAmountP4    = "buy_amount_p4"
	BuySlippage    = "buy_slippage"
	SellAmountP1   = "sell_amount_p1"
	SellAmountP2   = "sell_amount_p2"
	SellAmountP3   = "sell_amount_p3"
	SellSlippage   = "sell_slippage"
	SellProtection = "sell_protection"

	AwaitingBuySlippage  = "awaiting_buy_slippage"
	AwaitingBuyAmount    = "awaiting_buy_amount"
	AwaitingSellSlippage = "awaiting_sell_slippage"
	AwaitingSellAmount   = "awaiting_sell_amount"
)
