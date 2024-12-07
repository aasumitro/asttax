package handler

import (
	"log"
	"time"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/service"
	"github.com/aasumitro/asttax/internal/util/cache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	ContextDuration time.Duration = 5
	CacheDuration   time.Duration = 1
)

type Handler struct {
	bot     *tgbotapi.BotAPI
	userSrv service.IUserService
	cache   *cache.Cache
}

func (h *Handler) reply(r interface{}) {
	switch msg := r.(type) {
	case *tgbotapi.MessageConfig:
		if msg.Text == "" && msg.ChatID == 0 {
			log.Println(common.ErrNoMsg)
			return
		}
		if _, err := h.bot.Send(msg); err != nil {
			log.Printf("error sending reply: %v\n", err)
			return
		}
	case *tgbotapi.EditMessageTextConfig:
		if msg.Text == "" && msg.ChatID == 0 {
			log.Println(common.ErrNoMsg)
			return
		}
		if _, err := h.bot.Send(msg); err != nil {
			log.Printf("error sending reply: %v\n", err)
			return
		}
	default:
		log.Printf("unexpected reply type: %T", r)
	}
}

func (h *Handler) remove(chatID int64, msgID int) {
	deleteMsg := tgbotapi.NewDeleteMessage(chatID, msgID)
	if _, err := h.bot.Request(deleteMsg); err != nil {
		log.Printf("failed delete message: %v", err)
		return
	}
}

func (h *Handler) promptForUpdate(
	stateKey, stateValue, msgKey string,
	msg *tgbotapi.Message,
) {
	h.cache.Set(stateKey, stateValue, CacheDuration*time.Minute)
	h.cache.Set(msgKey, msg.MessageID, CacheDuration*time.Minute)
	txt := `
*Note:*

Please provide a valid decimal value. Examples include:
- Whole numbers: 1, 2, 3 . . . 
- Decimals: 0.0, 1.5, 2.75 . . .

**Important:**  
- *Sell Slippage*, *Sell Amount*, and *Buy Slippage* must be between 1 and 100.
- *Buy Amount* must start from 0.1 to ♾️

Kindly send us your update.
`
	reply := tgbotapi.NewEditMessageText(msg.Chat.ID, msg.MessageID, txt)
	reply.ParseMode = common.MessageParseMarkdown
	h.reply(&reply)
}

func (h *Handler) notifyProcessingUpdate(msg *tgbotapi.Message) {
	txt := "Your update is being processed. Please wait a moment while we apply the changes . . ."
	reply := tgbotapi.NewEditMessageText(msg.Chat.ID, msg.MessageID, txt)
	reply.ParseMode = common.MessageParseMarkdown
	h.reply(&reply)
}

func NewCommandHandler(
	bot *tgbotapi.BotAPI,
	userSrv service.IUserService,
	cachePool *cache.Cache,
) *Handler {
	return &Handler{
		bot:     bot,
		userSrv: userSrv,
		cache:   cachePool,
	}
}
