package main

import "fmt"

func main() {
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}

	for i := range messages {
		fmt.Println(i)
		fmt.Println(messages[i])
	}
}

// import (
// 	"errors"
// 	"fmt"
// )

// // var msg string

// // func init() {
// // 	msg = "from init()"
// // }

// func main() {
// 	// message := "Скоро я стану ниндзя!"

// 	// changeMessage(&message)
// 	// fmt.Println(message)

// 	// messages := [3]string {"1", "2", "3"}

// 	// messages := []string {"1", "2", "3"}

// 	// var messages []string
// 	// messages[0] = "1"
// 	// fmt.Println(messages)
// 	// fmt.Println(messages[1])
// 	// printMessage(messages)
// 	// fmt.Println(messages)

// 	// matrix := make([][]int, 10)

// 	// for x := 0; x < 10; x++ {
// 	// 	for y := 0; y < 10; y++ {
// 	// 		matrix[y] = make([]int, 10)
// 	// 		matrix[y][x] = x
// 	// 	}
// 	// 	fmt.Println(matrix[x])
// 	// }

// 	matrix := make([][]int, 10)

// 	counter := 0
// 	for x := 0; x < 10; x++ {
// 		matrix[x] = make([]int, 10)
// 		for y := 0; y < 10; y++ {
// 			counter++
// 			matrix[x][y] = counter
// 		}
// 		fmt.Println(matrix[x])
// 	}

// 	// for x := 0; true; x++ {
// 	// 	fmt.Println(x)
// 	// }
// }

// func printMessage(messages []string) error {
// 	if len(messages) == 0 {
// 		return errors.New("empty array")
// 	}

// 	messages[1] = "5"

// 	fmt.Println(messages)

// 	return nil
// }

// // func changeMessage(message *string) {
// // 	*message += " (из функции changeMessage())"
// // }
