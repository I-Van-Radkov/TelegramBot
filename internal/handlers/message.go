package handlers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/I-Van-Radkov/TelegramBot/clients/telegram"
	"github.com/I-Van-Radkov/TelegramBot/internal/constants"
	"github.com/I-Van-Radkov/TelegramBot/pkg/graph"
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
		log.Printf("[NORMAL] got new event: %s", update.Message.Text)

		if update.Message.Text == "/start" {
			u.tgClient.SendMessage(update.Message.Chat.ID, constants.MsgStart)
			return
		}
		if update.Message.Text == "/help" {
			u.tgClient.SendMessage(update.Message.Chat.ID, constants.MsgHelp)
			return
		}

		u.tgClient.SendMessage(update.Message.Chat.ID, constants.MsgDataAccepted) // Сообщение, что запрос принят

		matrix, start, end, err := parseMessage(update.Message.Text)

		if err != nil {
			switch err.Error() {
			case "insufficient data":
				u.tgClient.SendMessage(update.Message.Chat.ID, constants.MsgInsufficientData)
			case "error when parsing the matrix":
				u.tgClient.SendMessage(update.Message.Chat.ID, constants.MsgErrMatrix)
			case "error when parsing the start vertex":
				u.tgClient.SendMessage(update.Message.Chat.ID, constants.MsgErrStartVertex)
			case "error when parsing the end vertex":
				u.tgClient.SendMessage(update.Message.Chat.ID, constants.MsgErrEndVertex)
			default:
				u.tgClient.SendMessage(update.Message.Chat.ID, constants.MsgUnknown)
			}
			return
		}

		myGraph := graph.NewGraph(matrix, start, end)
		channel := make(chan string)

		go myGraph.Dijkstra(channel)

		// Ожидание результата вычисления: 30 секунд
		timeout := time.AfterFunc(30*time.Second, func() {
			channel <- constants.MsgTimeIsUp
		})

		result := <-channel
		timeout.Stop()

		u.tgClient.SendMessage(update.Message.Chat.ID, result)
	}
}

func parseMessage(message string) ([][]int, int, int, error) {
	var numCols = -1 //для проверки на квадратную матрицу

	lines := strings.Split(message, "\n")

	if len(lines) < 3 {
		return nil, 0, 0, fmt.Errorf("insufficient data")
	}

	var matrix [][]int

	for _, line := range lines[:len(lines)-2] {
		row := strings.Fields(line)

		if numCols == -1 {
			numCols = len(row)
		} else if len(row) != numCols { //если матрица не квадратная
			return nil, 0, 0, fmt.Errorf("error when parsing the matrix")
		}

		intRow := make([]int, len(row))

		for i, val := range row {
			/*if val == "беск" {
				intRow[i] = math.MaxInt64
			} else {*/
			num, err := strconv.Atoi(val)
			if err != nil {
				return nil, 0, 0, fmt.Errorf("error when parsing the matrix")
			}
			intRow[i] = num
		}
		matrix = append(matrix, intRow)
	}

	if numCols != len(matrix) { //если матрица не квадратная
		return nil, 0, 0, fmt.Errorf("error when parsing the matrix")
	}

	start, err := strconv.Atoi(lines[len(lines)-2]) //начальная вершина
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error when parsing the start vertex")
	}

	end, err := strconv.Atoi(lines[len(lines)-1]) //конечная вершина
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error when parsing the end vertex")
	}

	return matrix, start, end, nil
}
