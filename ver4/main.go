package main

import (
	"sample-worker/ver4/worker"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	m := worker.NewMessage()

	e.GET("/:id", m.GetID)

	e.Logger.Fatal(e.Start(":8080"))
}
