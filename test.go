package main

import (
	"fmt"
	"time"
)

func putter(ch chan int) {
	val := 5

	ch <- val
}

func test() {
	ch := make(chan int)

	go putter(ch)

	time.Sleep(5 * 1e9)

	fmt.Println(<-ch)

}
