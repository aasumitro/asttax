package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/aasumitro/asttax/internal/model"
)

type IUserRepository interface {
	Find(ctx context.Context, telegramID int64) (*model.User, error)
	Insert(ctx context.Context, param *model.User) (*model.User, error)
	Update(ctx context.Context, key string, value interface{}, telegramID int64) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (repo *userRepository) Find(
	ctx context.Context,
	telegramID int64,
) (*model.User, error) {
	query := `SELECT telegram_id, bot_language, accept_agreement, wallet_address, private_key,
	trade_fees, confirm_trade_protection,
	buy_amount_p1, buy_amount_p2, buy_amount_p3, buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
	sell_amount_p1, sell_amount_p2, sell_amount_p3, sell_slippage, sell_protection
    FROM users WHERE telegram_id = ?`
	row := repo.db.QueryRowContext(ctx, query, telegramID)
	return scan(row)
}

func (repo *userRepository) Insert(
	ctx context.Context,
	param *model.User,
) (*model.User, error) {
	query := `INSERT INTO users (telegram_id, accept_agreement, wallet_address, private_key, created_at) 
	VALUES ($1, $2, $3, $4, $5) RETURNING telegram_id, bot_language, accept_agreement, wallet_address, private_key,
    trade_fees, confirm_trade_protection,
    buy_amount_p1, buy_amount_p2, buy_amount_p3, buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
    sell_amount_p1, sell_amount_p2, sell_amount_p3, sell_slippage, sell_protection`
	row := repo.db.QueryRowContext(
		ctx, query, param.TelegramID, param.AcceptAgreement,
		param.WalletAddress, param.PrivateKey, time.Now().UnixMilli())
	return scan(row)
}

func (repo *userRepository) Update(
	ctx context.Context,
	key string,
	value interface{},
	telegramID int64,
) (*model.User, error) {
	validColumns := map[string]bool{
		"bot_language":             true,
		"trade_fees":               true,
		"confirm_trade_protection": true,
		"buy_amount_p1":            true,
		"buy_amount_p2":            true,
		"buy_amount_p3":            true,
		"buy_amount_p4":            true,
		"buy_amount_p5":            true,
		"buy_amount_p6":            true,
		"buy_slippage":             true,
		"sell_amount_p1":           true,
		"sell_amount_p2":           true,
		"sell_amount_p3":           true,
		"sell_slippage":            true,
		"sell_protection":          true,
	}
	if !validColumns[key] {
		return nil, fmt.Errorf("invalid column name: %s", key)
	}
	query := func() string {
		if v, ok := value.(float64); ok {
			return fmt.Sprintf(`UPDATE users SET %s = %f,`, key, v)
		}
		if v, ok := value.(string); ok && strings.Contains(v, "NOT") {
			return fmt.Sprintf(`UPDATE users SET %s = %s,`, key, v)
		}
		return fmt.Sprintf("UPDATE users SET %s = '%s',", key, value)
	}()
	query += ` updated_at = $1 WHERE telegram_id = $2
	RETURNING telegram_id, bot_language, accept_agreement, wallet_address, private_key,
		trade_fees, confirm_trade_protection, buy_amount_p1, buy_amount_p2, buy_amount_p3,
		buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage, sell_amount_p1, sell_amount_p2, 
		sell_amount_p3, sell_slippage, sell_protection`
	row := repo.db.QueryRowContext(ctx, query, time.Now().UnixMilli(), telegramID)
	return scan(row)
}

func scan(row *sql.Row) (*model.User, error) {
	var user model.User
	if err := row.Scan(
		&user.TelegramID, &user.BotLang, &user.AcceptAgreement,
		&user.WalletAddress, &user.PrivateKey,
		&user.TradeFees, &user.ConfirmTradeProtection,
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
