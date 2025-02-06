package handlers

import (
	"github.com/I-Van-Radkov/TelegramBot/clients/telegram"
	"github.com/I-Van-Radkov/TelegramBot/internal/constants"
)

type Updates struct {
	tgClient *telegram.Client
	Updates  []telegram.Update
}

func NewUpdates(client *telegram.Client) *Updates {
	return &Updates{
		tgClient: client,
		Updates:  []telegram.Update{},
	}
}

func (u *Updates) HandleUpdates() {
	for _, update := range u.Updates {
		if update.Message.Text != "" {
			u.tgClient.SendMessage(update.Message.Chat.ID, update.Message.Text)
		} else {
			u.tgClient.SendMessage(update.Message.Chat.ID, constants.MsgUnknown)
		}
	}
}
