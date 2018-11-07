package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan []string)

	go worker(ch, context)

	sliStr := []string{"aaa", "bbb", "ccc", "ddd", "eee"}
	time.Sleep(1 * time.Second)
	ch <- sliStr
	time.Sleep(1 * time.Second)
	s := []string{"fff", "ggg"}
	ch <- s
	time.Sleep(1 * time.Second)
}

func worker(ch <-chan []string, context context.Context) {
	n := 5
	for i := 0; i < n; i++ {
		go func() {
			select {
			case sli := <-ch:
				for _, s := range sli {
					fmt.Println(s)
				}
			case <-context.Done():
				return
			}
		}()
	}
}
