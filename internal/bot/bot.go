package bot

import (
	"log"
	"time"

	telegram "github.com/I-Van-Radkov/TelegramBot/client/telegram"
)

type Worker struct {
	tgClient *telegram.Client
	offset   int
	limit    int
}

func NewWorker(client *telegram.Client) *Worker {
	return &Worker{
		tgClient: client,
		limit:    100,
	}
}

func (w *Worker) Start() error {
	for {
		updates, err := w.tgClient.Updates(w.offset, w.limit)
		if err != nil {
			log.Printf("[ERR] getUpdates %s", err.Error())
			continue
		}

		if len(updates) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		w.offset = updates[len(updates)-1].ID + 1
		w.handleUpdates(updates)
	}
}

func (w *Worker) handleUpdates(updates []telegram.Update) {
	for _, update := range updates {
		w.tgClient.SendMessage(update.Message.Chat.ID, "Пошел нахуй, уеба, я еще в разработке, хули ты тут делаешь")
	}
}
