package main

import (
	"sample-worker/ver5/worker"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/:id", worker.GetID)

	e.Logger.Fatal(e.Start(":8080"))
}
