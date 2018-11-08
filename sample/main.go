package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan []string)

	go worker(ctx, ch)

	sliStr := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg"}
	ch <- sliStr

	time.Sleep(1 * time.Second)
	fmt.Println("---end---: ", time.Now())
}

func worker(ctx context.Context, ch <-chan []string) {
	n := 5
	fmt.Println("---start---: ", time.Now())
	for i := 0; i < n; i++ {
		go func() {
			select {
			case sli := <-ch:
				for _, s := range sli {
					fmt.Println(s)
				}
			case <-ctx.Done():
				fmt.Println("context is close: ", ctx.Err())
				return
			}
		}()
	}
}
