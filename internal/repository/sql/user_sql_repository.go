package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/aasumitro/asttax/internal/model"
)

type IUserRepository interface {
	Find(ctx context.Context, telegramID int64) (*model.User, error)
	Insert(ctx context.Context, param *model.User) (*model.User, error)
	Update(ctx context.Context, updates map[string]interface{}, telegramID int64) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (repo *userRepository) Find(
	ctx context.Context,
	telegramID int64,
) (*model.User, error) {
	query := `SELECT telegram_id, bot_language, accept_agreement, wallet_address, 
	trade_fees, confirm_trade_protection, buy_amount_p1, buy_amount_p2, buy_amount_p3, 
	buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
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
	VALUES ($1, $2, $3, $4, $5) RETURNING telegram_id, bot_language, accept_agreement, wallet_address, 
    trade_fees, confirm_trade_protection,  buy_amount_p1, buy_amount_p2, buy_amount_p3, 
	buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
    sell_amount_p1, sell_amount_p2, sell_amount_p3, sell_slippage, sell_protection`
	row := repo.db.QueryRowContext(
		ctx, query, param.TelegramID, param.AcceptAgreement,
		param.WalletAddress, param.PrivateKey, time.Now().UnixMilli())
	return scan(row)
}

func (repo *userRepository) Update(
	ctx context.Context,
	updates map[string]interface{},
	telegramID int64,
) (*model.User, error) {
	// Validate column names
	validColumns := map[string]bool{
		"bot_language":             true, // string
		"trade_fees":               true, // string
		"confirm_trade_protection": true, // boolean
		"buy_amount_p1":            true, // float64
		"buy_amount_p2":            true, // float64
		"buy_amount_p3":            true, // float64
		"buy_amount_p4":            true, // float64
		"buy_amount_p5":            true, // float64
		"buy_amount_p6":            true, // float64
		"buy_slippage":             true, // float64
		"sell_amount_p1":           true, // float64
		"sell_amount_p2":           true, // float64
		"sell_amount_p3":           true, // float64
		"sell_slippage":            true, // float64
		"sell_protection":          true, // boolean
	}
	// Build the SET clause
	var setClauses []string
	var args []interface{}
	argIndex := 1
	for key, value := range updates {
		if !validColumns[key] {
			return nil, errors.New("invalid column name")
		}
		// Handle NOT logic for booleans
		if slices.Contains([]string{"confirm_trade_protection", "sell_protection"}, key) && value == "TOGGLE" {
			setClauses = append(setClauses, fmt.Sprintf("%s = NOT %s", key, key))
		} else {
			setClauses = append(setClauses, fmt.Sprintf("%s = $%d", key, argIndex))
			args = append(args, value)
			argIndex++
		}
	}
	// Add updated_at and WHERE clause
	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, time.Now().UnixMilli())
	argIndex++
	// Build the query
	query := fmt.Sprintf(`
		UPDATE users SET %s WHERE telegram_id = $%d
		RETURNING telegram_id, bot_language, accept_agreement, wallet_address,
		          trade_fees, confirm_trade_protection, buy_amount_p1, buy_amount_p2, buy_amount_p3,
		          buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage, sell_amount_p1, sell_amount_p2, 
		          sell_amount_p3, sell_slippage, sell_protection`, strings.Join(setClauses, ", "), argIndex)
	args = append(args, telegramID)
	row := repo.db.QueryRowContext(ctx, query, args...)
	return scan(row)
}

func scan(row *sql.Row) (*model.User, error) {
	var user model.User
	if err := row.Scan(
		&user.TelegramID, &user.BotLang, &user.AcceptAgreement,
		&user.WalletAddress, &user.TradeFees, &user.ConfirmTradeProtection,
		&user.BuyAmountP1, &user.BuyAmountP2, &user.BuyAmountP3,
		&user.BuyAmountP4, &user.BuyAmountP5, &user.BuyAmountP6, &user.BuySlippage,
		&user.SellAmountP1, &user.SellAmountP2, &user.SellAmountP3,
		&user.SellSlippage, &user.SellProtection,
	); err != nil {
		return nil, fmt.Errorf("failed to scan user data: %w", err)
	}
	return &user, nil
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db: db}
}
