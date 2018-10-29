package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		for {
			select {
			case <-ch:
				fmt.Println(ch)
			}
		}

	}()
	ch <- "hello"

	time.Sleep(1 * time.Second)
}
