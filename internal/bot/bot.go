package bot

import (
	"log"
	"time"

	telegram "github.com/I-Van-Radkov/TelegramBot/clients/telegram"
	handlers "github.com/I-Van-Radkov/TelegramBot/internal/handlers"
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
		updates := handlers.NewUpdates(w.tgClient) //структура для получения обновлений

		dataUpdate, err := w.tgClient.Updates(w.offset, w.limit) //получаем слайс с обновлениями
		if err != nil {
			log.Printf("[ERR] getUpdates %s", err.Error())
			continue
		}
		updates.Updates = dataUpdate

		if len(updates.Updates) == 0 { //если не надено обновлений
			time.Sleep(1 * time.Second)
			continue
		}

		w.offset = updates.Updates[len(updates.Updates)-1].ID + 1

		updates.HandleUpdates() //обработка обновлений
	}
}
