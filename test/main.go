package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!!")
	ch := make(chan int, 10)

	//	product
	go func() {
		val := 1
		ticker := time.NewTicker(5 * time.Second)
		for _ = range ticker.C {
			ch <- val
			fmt.Println("write done: ", val)
			val++
		}
	}()

	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			val, ok := <-ch
			if ok {
				fmt.Println("consumer done: ", val)
			} else {
				fmt.Println("consumer nil")
			}
		}
	}()
	time.Sleep(time.Second * 100)

}
