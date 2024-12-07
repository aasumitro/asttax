package internal

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/config"
	"github.com/aasumitro/asttax/internal/handler"
	restRepo "github.com/aasumitro/asttax/internal/repository/rest"
	rpcRepo "github.com/aasumitro/asttax/internal/repository/rpc"
	sqlRepo "github.com/aasumitro/asttax/internal/repository/sql"
	"github.com/aasumitro/asttax/internal/service"
	"github.com/aasumitro/asttax/internal/util/cache"
	"github.com/blocto/solana-go-sdk/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run(ctx context.Context) {
	defer handlePanic()
	// load and init config
	cfg := config.LoadWith(ctx,
		config.SQLiteDBConnection(),
		config.InMemoryCache())
	log.Printf("Running %s . . .",
		cfg.GetServerIdentity())
	// make context notify
	ctxNC, stop := signal.NotifyContext(ctx,
		syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// make telegram bot instance
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}
	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)
	// register deps build handler
	commandHandler := registerHandler(bot, cfg)
	// stream update request
	for {
		select {
		case <-ctxNC.Done():
			updates.Clear()
			return
		case update, ok := <-updates:
			// validate updates
			if !ok {
				continue
			}
			time.Sleep(1 * time.Millisecond)
			// handle empty request
			if update.Message == nil && update.CallbackQuery == nil {
				continue
			}
			// Handling main commands
			if update.Message != nil && update.Message.IsCommand() {
				handleCommand(commandHandler, update.Message)
				continue
			}
			// Handling Callback Queries
			if update.CallbackQuery != nil {
				handleCallback(commandHandler, update.CallbackQuery)
				continue
			}
			// Handle user responses based on state
			if update.Message != nil {
				handleState(commandHandler, cfg.CachePool, update.Message)
				continue
			}
		}
	}
}

func handlePanic() {
	if r := recover(); r != nil {
		log.Printf("Recovered from panic: %v\nStack trace:\n%s",
			r, debug.Stack())
	}
}

func registerHandler(
	bot *tgbotapi.BotAPI,
	cfg *config.Config,
) *handler.Handler {
	rpcClient := client.NewClient(cfg.GetRPCEndpoint())
	userRepo := sqlRepo.NewUserRepository(cfg.SQLPool)
	solanaRPCRepo := rpcRepo.NewSolanaRPCRepository(rpcClient, cfg.CachePool)
	coingeckoRESTRepo := restRepo.NewCoingeckoRepository(cfg.CoingeckoAPIURL, cfg.CachePool)
	userSrv := service.NewUserService(userRepo, coingeckoRESTRepo, solanaRPCRepo, cfg.SecretKey)
	return handler.NewCommandHandler(bot, userSrv, cfg.CachePool)
}

func handleCommand(
	h *handler.Handler,
	msg *tgbotapi.Message,
) {
	switch msg.Command() {
	case common.Start:
		h.StartCommand(msg)
	case common.Buy:
		h.BuyCommand(msg)
	case common.Sell:
		h.SellCommand(msg)
	case common.Positions:
		h.PositionsCommand(msg)
	case common.Settings:
		h.SettingsCommand(msg)
	case common.Withdraw:
		h.WithdrawCommand(msg)
	case common.Help:
		h.HelpCommand(msg)
	case common.Backup:
		h.BackupCommand(msg)
	}
}

func handleCallback(
	h *handler.Handler,
	cq *tgbotapi.CallbackQuery,
) {
	switch cq.Data {
	case common.AcceptAgreement:
		h.AcceptAgreementCallback(cq.Message)
	case common.Start:
		h.StartCallback(cq.Message)
	case common.Buy:
		h.BuyCallback(cq.Message)
	case common.Sell:
		h.SellCallback(cq.Message)
	case common.Positions:
		h.PositionsCallback(cq.Message)
	case common.NewPairs:
		h.NewPairCallback(cq.Message)
	case common.Settings:
		h.SettingCallback(cq.Message)
	case common.Help:
		h.HelpCallback(cq.Message)
	case common.LanguageSettings:
		h.LanguageSettingCallback(cq.Message)
	case common.BackToStart:
		h.BackToStartCallback(cq.Message)
	case common.BackToSetting:
		h.BackToSettingCallback(cq.Message)
	case common.Refresh:
		h.RefreshCallback(cq.Message)
	// settings handler
	case common.FastTradeFee:
		h.EditTradeFeeState(cq.Message, "fast")
	case common.TurboTradeFee:
		h.EditTradeFeeState(cq.Message, "turbo")
	case common.ConfirmTrade:
		h.EditConfirmTradeState(cq.Message)
	case common.BuyAmountP1:
		h.EditBuyAmountState(cq.Data, cq.Message, 1)
	case common.BuyAmountP2:
		h.EditBuyAmountState(cq.Data, cq.Message, 2)
	case common.BuyAmountP3:
		h.EditBuyAmountState(cq.Data, cq.Message, 3)
	case common.BuyAmountP4:
		h.EditBuyAmountState(cq.Data, cq.Message, 4)
	case common.BuyAmountP5:
		h.EditBuyAmountState(cq.Data, cq.Message, 5)
	case common.BuyAmountP6:
		h.EditBuyAmountState(cq.Data, cq.Message, 6)
	case common.BuySlippage:
		h.EditBuySlippageState(cq.Data, cq.Message)
	case common.SellAmountP1:
		h.EditSellAmountState(cq.Data, cq.Message, 1)
	case common.SellAmountP2:
		h.EditSellAmountState(cq.Data, cq.Message, 2)
	case common.SellAmountP3:
		h.EditSellAmountState(cq.Data, cq.Message, 3)
	case common.SellSlippage:
		h.EditSellSlippageState(cq.Data, cq.Message)
	case common.SellProtection:
		h.EditSellProtectionState(cq.Message)
	}
}

func handleState(
	h *handler.Handler,
	cache *cache.Cache,
	msg *tgbotapi.Message,
) {
	chatID := msg.Chat.ID
	key := fmt.Sprintf("%d_state", chatID)
	var command string
	if cacheData, ok := cache.Get(key); ok {
		if cmd, ok := cacheData.(string); ok {
			command = cmd
		}
	}
	switch command {
	case common.AwaitingBuySlippage:
		h.EditBuySlippageState(command, msg)
	case common.AwaitingBuyAmount:
		h.EditBuyAmountState(command, msg, 0)
	case common.AwaitingSellSlippage:
		h.EditSellSlippageState(command, msg)
	case common.AwaitingSellAmount:
		h.EditSellAmountState(command, msg, 0)
	}
}
