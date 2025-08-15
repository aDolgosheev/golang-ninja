package main

import "fmt"

func main() {

	defer handlerPanic()

	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}

	messages[4] = "message 5"

	fmt.Println(messages)

	//panic("aaaaaaaa help!")
}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
