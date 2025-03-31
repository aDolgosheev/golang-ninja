package main

import (
	"errors"
	"fmt"
)

// var msg string

// func init() {
// 	msg = "from init()"
// }

func main() {
	// fmt.Println(msg)

	// var p *int
	// message := "Скоро я стану ниндзя!"
	// fmt.Println(message)

	// changeMessage(&message)

	// fmt.Println(message)

	// messages := [3]string{"1", "2", "3"}
	// fmt.Println(messages)
	// printMessages(messages)
	// fmt.Println(messages)

	// messages := []string{"1", "2", "3"}

	messages := make([]string, 5)

	messages = append(messages, "6")

	// printMessages(messages)
	fmt.Println(messages)
}

func printMessages(messages []string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}

	messages[1] = "5"

	fmt.Println(messages)

	return nil
}

// func changeMessage(message *string) {
// 	*message += " (из функции changeMessage())"
// }