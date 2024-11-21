package internal

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run(ctx context.Context) {
	cfg := config.LoadWith(ctx, config.SQLiteDBConnection())
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
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for {
		select {
		case <-ctxNC.Done():
			updates.Clear()
			return
		case update, ok := <-updates:
			if !ok {
				continue
			}
			if update.Message == nil && update.CallbackQuery == nil {
				continue
			}
			time.Sleep(1 * time.Millisecond)

			// Handling commands (e.g., /start, /help)
			if update.Message != nil && update.Message.IsCommand() {
				var msg tgbotapi.MessageConfig
				switch update.Message.Command() {
				case "start":
					// Create the message object for the /start command
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, mainText())
					msg.ParseMode = "Markdown"
					msg.ReplyMarkup = common.StartReplyMarkup
					msg.ReplyToMessageID = update.Message.MessageID
				case "help":
					// Edit the message to show help content when "help" is selected
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, helpText())
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = common.BackReplyMarkup
				}
				if msg.Text == "" && msg.ChatID == 0 {
					log.Println("skip no message")
					continue
				}
				if _, err := bot.Send(msg); err != nil {
					log.Println(err)
				}
			}

			// Handling Callback Queries for help and back
			if update.CallbackQuery != nil {
				chatID := update.CallbackQuery.Message.Chat.ID
				data := update.CallbackQuery.Data
				var editMsg tgbotapi.EditMessageTextConfig
				switch data {
				case "help":
					editMsg = tgbotapi.NewEditMessageTextAndMarkup(chatID,
						update.CallbackQuery.Message.MessageID, helpText(), common.BackReplyMarkup)
					editMsg.ParseMode = "HTML"
				case "refresh":
					// handle if a new message same as before
					editMsg = tgbotapi.NewEditMessageTextAndMarkup(chatID,
						update.CallbackQuery.Message.MessageID, mainText(), common.StartReplyMarkup)
					editMsg.ParseMode = "Markdown"
				case "back_to_start":
					editMsg = tgbotapi.NewEditMessageTextAndMarkup(chatID,
						update.CallbackQuery.Message.MessageID, mainText(), common.StartReplyMarkup)
					editMsg.ParseMode = "Markdown"
				}
				if editMsg.Text == "" && editMsg.ChatID == 0 {
					log.Println("skip no message")
					continue
				}
				if _, err := bot.Send(editMsg); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func mainText() string {
	solanaAddress := "3x6QDiKyZR4vDjtJGXyXcEsQyh4CX2QoxyLjjhVFkqCG"
	solanaURL := "https://solscan.io/account/" + solanaAddress
	solanaValue := "0 SOL ($0.00)"
	return "Welcome to AsttaX on Solana! \n\n" +
		" *Solana* • [🅴](" + solanaURL + ")\n" +
		"💰`" + solanaAddress + "` _(Tap to copy)_\n" +
		"🪙Balance: `" + solanaValue + "`\n\n\n" +
		"Click on the buttons below to navigate:\n" +
		"• Use the Buy/Sell buttons to trade.\n" +
		"• Use the Refresh button to update your current balance.\n" +
		"• Check your positions, set limit orders, and more!\n\n"
}

func helpText() string {
	youtubePlaylistURL := "https://www.youtube.com/watch?v=HavGDGUTmgs&list=PLmAMfj0qP2wwfnuRJQge2ss4sJxnhIqyt&ab_channel=Solandy%5Bsolandy.sol%5D"
	return fmt.Sprintf(`
📖 <b>Help</b>

<b><u>How do I use AsttaX?</u></b>
Check out our <a href="%s">Youtube playlist</a> where we explain it all.

<b><u>Which tokens can I trade?</u></b>
Any SPL token that is tradeable via Jupiter, including SOL and USDC pairs. We also support directly trading through Raydium if Jupiter fails to find a route. You can trade newly created SOL pairs (not USDC) directly through Raydium.

<b><u>My transaction timed out. What happened?</u></b>
Transaction timeouts can occur when there is heavy network load or instability. This is simply the nature of the current Solana network.

<b><u>What are the fees for using AsttaX?</u></b>
Transactions through Trojan incur a fee of 1%%, or 0.9%% if you were referred by another user. We don't charge a subscription fee or pay-wall any features.

<b><u>My net profit seems wrong, why is that?</u></b>
The net profit of a trade takes into consideration the trade's transaction fees. Confirm the details of your trade on Solscan.io to verify the net profit.

<b><u>Additional questions or need support?</u></b>
send us mail hello@astta.xyz and one of our admins can assist you.
`, youtubePlaylistURL)
}
