package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type message struct {
	data string
}

func main() {
	e := echo.New()
	m := &message{}
	e.GET("/:id", m.getID)

	e.Logger.Fatal(e.Start(":8080"))
}

func (m *message) getID(c echo.Context) error {
	id := c.Param("id")
	ch := make(chan *message)

	m.data = id

	go func() {
		fmt.Println(m)
		time.Sleep(2 * time.Second)
		ch <- m
	}()

	fmt.Println("---start---")
	<-ch
	fmt.Println("---end---")

	return c.String(http.StatusOK, fmt.Sprintf("return id: ", m.data))
}
