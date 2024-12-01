package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/aasumitro/asttax/internal/model"
)

type IUserRepository interface {
	Find(ctx context.Context, telegramID int64) (*model.User, error)
	Insert(ctx context.Context, param *model.User) (*model.User, error)
	Update(ctx context.Context, param *model.User) (*model.User, error)
	// Delete(ctx context.Context, telegramID int64) error
}

type userRepository struct {
	db *sql.DB
}

func (repo *userRepository) Find(
	ctx context.Context,
	telegramID int64,
) (*model.User, error) {
	var user model.User
	query := `
    SELECT telegram_id, bot_language, accept_agreement, wallet_address, private_key,
           trade_fees, custom_trade_fee, mev_buy_protection, mev_sell_protection, confirm_trade_protection,
           buy_amount_p1, buy_amount_p2, buy_amount_p3, buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
           sell_amount_p1, sell_amount_p2, sell_amount_p3, sell_slippage, sell_protection
    FROM users
    WHERE telegram_id = ?`
	if err := repo.db.QueryRowContext(ctx, query, telegramID).Scan(
		&user.TelegramID, &user.BotLang, &user.AcceptAgreement,
		&user.WalletAddress, &user.PrivateKey,
		&user.TradeFees, &user.CustomTradeFee,
		&user.MEVBuyProtection, &user.MEVSellProtection, &user.ConfirmTradeProtection,
		&user.BuyAmountP1, &user.BuyAmountP2, &user.BuyAmountP3,
		&user.BuyAmountP4, &user.BuyAmountP5, &user.BuyAmountP6, &user.BuySlippage,
		&user.SellAmountP1, &user.SellAmountP2, &user.SellAmountP3,
		&user.SellSlippage, &user.SellProtection,
	); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) Update(
	_ context.Context,
	_ *model.User,
) (*model.User, error) {
	return nil, nil
}

func (repo *userRepository) Insert(
	ctx context.Context,
	param *model.User,
) (*model.User, error) {
	query := `
	INSERT INTO users (telegram_id, accept_agreement, wallet_address, private_key, created_at) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING telegram_id, bot_language, accept_agreement, wallet_address, private_key,
           trade_fees, custom_trade_fee, mev_buy_protection, mev_sell_protection, confirm_trade_protection,
           buy_amount_p1, buy_amount_p2, buy_amount_p3, buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
           sell_amount_p1, sell_amount_p2, sell_amount_p3, sell_slippage, sell_protection
	`
	row := repo.db.QueryRowContext(
		ctx, query, param.TelegramID, param.AcceptAgreement,
		param.WalletAddress, param.PrivateKey, time.Now().UnixMilli())
	var user model.User
	if err := row.Scan(
		&user.TelegramID, &user.BotLang, &user.AcceptAgreement,
		&user.WalletAddress, &user.PrivateKey,
		&user.TradeFees, &user.CustomTradeFee,
		&user.MEVBuyProtection, &user.MEVSellProtection, &user.ConfirmTradeProtection,
		&user.BuyAmountP1, &user.BuyAmountP2, &user.BuyAmountP3,
		&user.BuyAmountP4, &user.BuyAmountP5, &user.BuyAmountP6, &user.BuySlippage,
		&user.SellAmountP1, &user.SellAmountP2, &user.SellAmountP3,
		&user.SellSlippage, &user.SellProtection,
	); err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db: db}
}
