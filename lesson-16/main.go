package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan string
	fmt.Println(msg)

	msg = make(chan string)
	fmt.Println(msg)

	go func() {
		time.Sleep(2 * time.Second)
		msg <- "Канал ниндзя"
	}()

	// value := <-msg
	// fmt.Println(value)
	fmt.Println(<-msg)
}
