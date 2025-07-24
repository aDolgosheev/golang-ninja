package main

import "fmt"

// var msg string

// func init() {
// 	msg = "from init()"
// }

func main() {
	// message := "Скоро я стану ниндзя!"

	// changeMessage(&message)
	// fmt.Println(message)

	messages := [3]string {"1", "2", "3"}
	fmt.Println(messages)
	fmt.Println(messages[1])
}

// func changeMessage(message *string) {
// 	*message += " (из функции changeMessage())"
// }


