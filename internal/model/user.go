package model

const (
	fastTrade  = "fast"
	turboTrade = "turbo"
)

type User struct {
	// user cred and bot setting
	TelegramID      int64  `sql:"telegram_id"`
	BotLang         string `sql:"bot_language"`
	AcceptAgreement bool   `sql:"accept_agreement"`
	// user wallet
	WalletAddress string `sql:"wallet_address"`
	PrivateKey    string `sql:"private_key"`
	// trade fee
	TradeFees string `sql:"trade_fees"`
	// trade protection
	ConfirmTradeProtection bool `sql:"confirm_trade_protection"`
	// buy amount
	BuyAmountP1 float64 `sql:"buy_amount_p1"`
	BuyAmountP2 float64 `sql:"buy_amount_p2"`
	BuyAmountP3 float64 `sql:"buy_amount_p3"`
	BuyAmountP4 float64 `sql:"buy_amount_p4"`
	BuySlippage float64 `sql:"buy_slippage"`
	// sell amount
	SellAmountP1   float64 `sql:"sell_amount_p1"`
	SellAmountP2   float64 `sql:"sell_amount_p2"`
	SellAmountP3   float64 `sql:"sell_amount_p3"`
	SellSlippage   float64 `sql:"sell_slippage"`
	SellProtection bool    `sql:"sell_protection"`
}

func (u *User) ToTradeFee() float64 {
	tradeFee := 0.0
	switch u.TradeFees {
	case fastTrade:
		tradeFee = 0.0015
	case turboTrade:
		tradeFee = 0.0075
	}
	return tradeFee
}
