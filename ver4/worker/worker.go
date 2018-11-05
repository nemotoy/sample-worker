package worker

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type messageService struct {
	ch chan *message
}

type message struct {
	data []string
	id   int
}

func NewMessage() *messageService {

	return &messageService{
		ch: make(chan *message),
	}
}

func (m *messageService) GetID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	go func() {
		for {
			message := <-m.ch
			go m.worker(message)
		}
	}()

	m.ch <- &message{
		data: []string{"a", "b", "c"},
		id:   id,
	}

	return c.String(http.StatusOK, "ok")
}

func (m *messageService) worker(message *message) {
	// for _, m := range message.data {
	// 	fmt.Println(m)
	// }
	fmt.Println(message.id)
}
