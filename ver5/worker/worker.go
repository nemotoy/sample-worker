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

func GetID(c echo.Context) error {

	ch := make(chan *message, 1)

	id, _ := strconv.Atoi(c.Param("id"))

	go func() {
		defer close(ch)
		message := <-ch
		go worker(message)
	}()

	ch <- &message{
		data: []string{"a", "b", "c"},
		id:   id,
	}

	return c.String(http.StatusOK, "ok")
}

func worker(message *message) {
	fmt.Println(message.id)
}
