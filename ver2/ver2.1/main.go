package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

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
	ch := make(chan *message, worker)

	m := &messageService{
		ch: ch,
		wg: wg,
	}
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go m.worker()

	}

	return m
}

func (m *messageService) getID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	message := &message{
		data: []string{"a", "b", "c"},
		id:   id,
	}
	m.ch <- message
	close(m.ch)
	m.wg.Wait()
	// http://cuto.unirita.co.jp/gostudy/post/go_chanel/

	return c.String(http.StatusOK, "ok")
}

func (m *messageService) worker() {
	defer m.wg.Done()
	for {
		message := <-m.ch
		for _, m := range message.data {
			fmt.Println(m)
		}
		fmt.Println(message.id)
	}
}
