package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
)

type messageService struct {
	ch chan *message
	wg *sync.WaitGroup
}

type message struct {
	data []string
	id   int
}

const worker = 3

func main() {
	e := echo.New()
	m := newMessage()

	e.GET("/:id", m.getID)

	e.Logger.Fatal(e.Start(":8080"))
}

func newMessage() *messageService {

	wg := &sync.WaitGroup{}
	ch := make(chan *message)

	m := &messageService{
		ch: ch,
		wg: wg,
	}

	return m
}

func (m *messageService) getID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	message := &message{
		data: []string{"a", "b", "c"},
		id:   id,
	}

	go func() {
		for {
			message := <-m.ch
			time.Sleep(2 * time.Second)
			go m.worker(message)
		}
	}()

	m.ch <- message

	return c.String(http.StatusOK, "ok")
}

func (m *messageService) worker(message *message) {
	for _, m := range message.data {
		fmt.Println(m)
	}
	fmt.Println(message.id)
}
