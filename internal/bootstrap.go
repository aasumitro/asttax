package internal

import (
	"context"
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
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run(ctx context.Context) {
	defer handlePanic()
	// load and init config
	cfg := config.LoadWith(ctx,
		config.SQLiteDBConnection(),
		config.InMemoryCache())
	log.Printf("Running %s v%s . . .",
		cfg.ServerName, cfg.ServerVersion)
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
	// register deps
	rpcClient := client.NewClient(rpc.DevnetRPCEndpoint)
	userRepo := sqlRepo.NewUserRepository(cfg.SQLPool)
	solanaRPCRepo := rpcRepo.NewSolanaRPCRepository(rpcClient, cfg.CachePool)
	coingeckoRESTRepo := restRepo.NewCoingeckoRepository(cfg.CoingeckoAPIURL, cfg.CachePool)
	userSrv := service.NewUserService(userRepo, coingeckoRESTRepo, solanaRPCRepo, cfg.SecretKey)
	commandHandler := handler.NewCommandHandler(bot, userSrv)
	callbackHandler := handler.NewCallbackHandler(bot, userSrv)
	settingHandler := handler.NewSettingHandler(bot, userSrv)
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
			}
			// Handling Callback Queries
			if update.CallbackQuery != nil {
				handleCallback(callbackHandler, settingHandler, update.CallbackQuery)
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

func handleCommand(
	h *handler.Command,
	msg *tgbotapi.Message,
) {
	switch msg.Command() {
	case common.Start:
		h.Start(msg)
	case common.Buy:
		h.Buy(msg)
	case common.Sell:
		h.Sell(msg)
	case common.Positions:
		h.Positions(msg)
	case common.Settings:
		h.Settings(msg)
	case common.Withdraw:
		h.Withdraw(msg)
	case common.Help:
		h.Help(msg)
	case common.Backup:
		h.Backup(msg)
	}
}

func handleCallback(
	h *handler.Callback,
	hs *handler.Setting,
	cq *tgbotapi.CallbackQuery,
) {
	switch cq.Data {
	case common.AcceptAgreement:
		h.AcceptAgreement(cq.Message)
	case common.Start:
		h.Start(cq.Message)
	case common.Buy:
		h.Buy(cq.Message)
	case common.Sell:
		h.Sell(cq.Message)
	case common.Positions:
		h.Positions(cq.Message)
	case common.NewPairs:
		h.NewPair(cq.Message)
	case common.Settings:
		h.Setting(cq.Message)
	case common.Help:
		h.Help(cq.Message)
	case common.LanguageSettings:
		h.LanguageSetting(cq.Message)
	case common.BackToStart:
		h.BackToStart(cq.Message)
	case common.BackToSetting:
		h.BackToSetting(cq.Message)
	case common.Refresh:
		h.Refresh(cq.Message)
	// settings handler
	case common.FastTradeFee:
		hs.EditTradeFee(cq.Message, "fast")
	case common.TurboTradeFee:
		hs.EditTradeFee(cq.Message, "turbo")
	case common.ConfirmTrade:
		hs.EditConfirmTrade(cq.Message)
	case common.BuyAmountP1:
		hs.EditBuyAmount(cq.Message, 1)
	case common.BuyAmountP2:
		hs.EditBuyAmount(cq.Message, 2)
	case common.BuyAmountP3:
		hs.EditBuyAmount(cq.Message, 3)
	case common.BuyAmountP4:
		hs.EditBuyAmount(cq.Message, 4)
	case common.BuyAmountP5:
		hs.EditBuyAmount(cq.Message, 5)
	case common.BuyAmountP6:
		hs.EditBuyAmount(cq.Message, 6)
	case common.BuySlippage:
		hs.EditBuySlippage(cq.Message)
	case common.SellAmountP1:
		hs.EditSellAmount(cq.Message, 1)
	case common.SellAmountP2:
		hs.EditSellAmount(cq.Message, 2)
	case common.SellAmountP3:
		hs.EditSellAmount(cq.Message, 3)
	case common.SellSlippage:
		hs.EditSellSlippage(cq.Message)
	case common.SellProtection:
		hs.EditSellProtection(cq.Message)
	}
}
