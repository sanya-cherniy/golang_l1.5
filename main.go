package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var w sync.WaitGroup
	w.Add(1)

	go func() {
		defer w.Done()
		time.Sleep(5 * time.Second)
	}()
	go func(ch chan int) {
		x := 1
		for {
			ch <- x
			time.Sleep(1 * time.Second)
			x += x
		}

	}(ch)
	go func(ch chan int) {
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
			}
		}

	}(ch)
	w.Wait()
}
