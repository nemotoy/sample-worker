package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	m := newMessage()

	e.GET("/:id", m.getID)

	e.Logger.Fatal(e.Start(":8080"))
}

type messageService struct {
	ch chan *message
}

type message struct {
	data []string
	id   int
}

func newMessage() *messageService {
	ch := make(chan *message)
	m := &messageService{
		ch: ch,
	}

	go func() {
		for {
			message := <-ch
			go m.worker(message)
		}
	}()

	return m
}

func (m *messageService) getID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	m.ch <- &message{
		data: []string{"a", "b", "c"},
		id:   id,
	}

	return c.String(http.StatusOK, "ok")
}

func (m *messageService) worker(message *message) {
	fmt.Println(message.id)
}
