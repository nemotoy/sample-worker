package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []string{"aaa", "bbb", "ccc"}
	job := make(chan []string)
	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			data := <-job
			for _, d := range data {
				fmt.Println(d)
			}
			wg.Done()
		}()
	}

	job <- data
	close(job)

	wg.Wait()
}
