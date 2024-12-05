package sql_test

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/asttax/internal/model"
	sqlRepo "github.com/aasumitro/asttax/internal/repository/sql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type userRepositoryTestSuite struct {
	suite.Suite
	mock     sqlmock.Sqlmock
	userRepo sqlRepo.IUserRepository
}

func (suite *userRepositoryTestSuite) SetupSuite() {
	var err error
	var db *sql.DB

	db, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.userRepo = sqlRepo.NewUserRepository(db)
}

func (suite *userRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *userRepositoryTestSuite) Test_Find_ExpectedReturnDataRows() {
	user := suite.mock.
		NewRows([]string{"telegram_id", "bot_language", "accept_agreement", "wallet_address", "private_key",
			"trade_fees", "confirm_trade_protection",
			"buy_amount_p1", "buy_amount_p2", "buy_amount_p3", "buy_amount_p4", "buy_amount_p5", "buy_amount_p6", "buy_slippage",
			"sell_amount_p1", "sell_amount_p2", "sell_amount_p3", "sell_slippage", "sell_protection"}).
		AddRow(12345, "en", 1, "1w23asd222222233test", "this-mock-private-key", "fast",
			0, 0.25, 0.5, 1, 2.5, 5, 10, 15, 25, 50, 100, 15, 0)
	query := `
    SELECT telegram_id, bot_language, accept_agreement, wallet_address, private_key,
           trade_fees, confirm_trade_protection,
           buy_amount_p1, buy_amount_p2, buy_amount_p3, buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
           sell_amount_p1, sell_amount_p2, sell_amount_p3, sell_slippage, sell_protection
    FROM users
    WHERE telegram_id = ?`
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(user)
	res, err := suite.userRepo.Find(context.TODO(), 12345)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *userRepositoryTestSuite) Test_Find_ExpectedReturnError() {
	user := suite.mock.
		NewRows([]string{"telegram_id", "bot_language", "accept_agreement", "wallet_address", "private_key",
			"trade_fees", "confirm_trade_protection",
			"buy_amount_p1", "buy_amount_p2", "buy_amount_p3", "buy_amount_p4", "buy_amount_p5", "buy_amount_p6", "buy_slippage",
			"sell_amount_p1", "sell_amount_p2", "sell_amount_p3", "sell_slippage", "sell_protection"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil,
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	query := `
    SELECT telegram_id, bot_language, accept_agreement, wallet_address, private_key,
           trade_fees, confirm_trade_protection,
           buy_amount_p1, buy_amount_p2, buy_amount_p3, buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
           sell_amount_p1, sell_amount_p2, sell_amount_p3, sell_slippage, sell_protection
    FROM users
    WHERE telegram_id = ?`
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(user)
	res, err := suite.userRepo.Find(context.TODO(), 12345)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *userRepositoryTestSuite) Test_Insert_ExpectedReturnDataRows() {
	user := &model.User{AcceptAgreement: true}
	user.TelegramID = 12345
	user.WalletAddress = "1w23asd222222233test"
	user.PrivateKey = "this-mock-private-key"
	rows := suite.mock.
		NewRows([]string{"telegram_id", "bot_language", "accept_agreement", "wallet_address", "private_key",
			"trade_fees", "confirm_trade_protection",
			"buy_amount_p1", "buy_amount_p2", "buy_amount_p3", "buy_amount_p4", "buy_amount_p5", "buy_amount_p6", "buy_slippage",
			"sell_amount_p1", "sell_amount_p2", "sell_amount_p3", "sell_slippage", "sell_protection"}).
		AddRow(12345, "en", 1, "1w23asd222222233test", "this-mock-private-key", "fast",
			0, 0.25, 0.5, 1, 2.5, 5, 10, 15, 25, 50, 100, 15, 0)
	query := `
	INSERT INTO users (telegram_id, accept_agreement, wallet_address, private_key, created_at) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING telegram_id, bot_language, accept_agreement, wallet_address, private_key,
           trade_fees, confirm_trade_protection,
           buy_amount_p1, buy_amount_p2, buy_amount_p3, buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
           sell_amount_p1, sell_amount_p2, sell_amount_p3, sell_slippage, sell_protection
	`
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(user.TelegramID, user.AcceptAgreement,
			user.WalletAddress, user.PrivateKey,
			time.Now().UnixMilli()).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.userRepo.Insert(context.TODO(), user)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *userRepositoryTestSuite) Test_Insert_ExpectedReturnError() {
	user := &model.User{AcceptAgreement: true}
	user.TelegramID = 12345
	user.WalletAddress = "1w23asd222222233test"
	user.PrivateKey = "this-mock-private-key"
	rows := suite.mock.
		NewRows([]string{"telegram_id", "bot_language", "accept_agreement", "wallet_address", "private_key",
			"trade_fees", "confirm_trade_protection",
			"buy_amount_p1", "buy_amount_p2", "buy_amount_p3", "buy_amount_p4", "buy_amount_p5", "buy_amount_p6", "buy_slippage",
			"sell_amount_p1", "sell_amount_p2", "sell_amount_p3", "sell_slippage", "sell_protection"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil,
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	query := `
	INSERT INTO users (telegram_id, accept_agreement, wallet_address, private_key, created_at) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING telegram_id, bot_language, accept_agreement, wallet_address, private_key,
           trade_fees, confirm_trade_protection,
           buy_amount_p1, buy_amount_p2, buy_amount_p3, buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
           sell_amount_p1, sell_amount_p2, sell_amount_p3, sell_slippage, sell_protection
	`
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(user.TelegramID, user.AcceptAgreement,
			user.WalletAddress, user.PrivateKey,
			time.Now().UnixMilli()).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.userRepo.Insert(context.TODO(), user)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *userRepositoryTestSuite) Test_Update_ExpectedReturnDataRows() {
	rows := suite.mock.
		NewRows([]string{"telegram_id", "bot_language", "accept_agreement", "wallet_address", "private_key",
			"trade_fees", "confirm_trade_protection",
			"buy_amount_p1", "buy_amount_p2", "buy_amount_p3", "buy_amount_p4", "buy_amount_p5", "buy_amount_p6", "buy_slippage",
			"sell_amount_p1", "sell_amount_p2", "sell_amount_p3", "sell_slippage", "sell_protection"}).
		AddRow(12345, "en", 1, "1w23asd222222233test", "this-mock-private-key", "fast",
			0, 0.25, 0.5, 1, 2.5, 5, 10, 15, 25, 50, 100, 15, 0)
	query := fmt.Sprintf(`UPDATE users SET %s = '%s', updated_at = $1 WHERE telegram_id = $2 
		RETURNING telegram_id, bot_language, accept_agreement, wallet_address, private_key,
		trade_fees, confirm_trade_protection,
		buy_amount_p1, buy_amount_p2, buy_amount_p3, buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
		sell_amount_p1, sell_amount_p2, sell_amount_p3, sell_slippage, sell_protection`, "trade_fees", "fast")
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(time.Now().UnixMilli(), 12345).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.userRepo.Update(context.TODO(), "trade_fees", "fast", 12345)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *userRepositoryTestSuite) Test_Update_ExpectedReturnError() {
	rows := suite.mock.
		NewRows([]string{"telegram_id", "bot_language", "accept_agreement", "wallet_address", "private_key",
			"trade_fees", "confirm_trade_protection",
			"buy_amount_p1", "buy_amount_p2", "buy_amount_p3", "buy_amount_p4", "buy_amount_p5", "buy_amount_p6", "buy_slippage",
			"sell_amount_p1", "sell_amount_p2", "sell_amount_p3", "sell_slippage", "sell_protection"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil,
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	query := fmt.Sprintf(`UPDATE users SET %s = '%s', updated_at = $1 WHERE telegram_id = $2 
		RETURNING telegram_id, bot_language, accept_agreement, wallet_address, private_key,
		trade_fees, confirm_trade_protection,
		buy_amount_p1, buy_amount_p2, buy_amount_p3, buy_amount_p4, buy_amount_p5, buy_amount_p6, buy_slippage,
		sell_amount_p1, sell_amount_p2, sell_amount_p3, sell_slippage, sell_protection`, "trade_fees", "fast")
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(time.Now().UnixMilli(), 12345).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.userRepo.Update(context.TODO(), "trade_fees", "fast", 12345)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(userRepositoryTestSuite))
}
