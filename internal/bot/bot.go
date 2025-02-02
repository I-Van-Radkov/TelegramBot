package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/I-Van-Radkov/TelegramBot/client/telegram"
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
		fmt.Println(updates)
		if err != nil {
			log.Printf("[ERR] getUpdates %s", err.Error())
			continue
		}

		if len(updates) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

	}
}

func (w *Worker) handleUpdates(updates []telegram.Update) {
	log.Print("хуй")
	for _, update := range updates {
		w.tgClient.SendMessage(update.Message.Chat.ID, "Тестовый запуск, проверка работы сообщений")
	}
}
