package main

import (
	"errors"
	"fmt"
	"log"
)

var a, b, c int

func main() {
	fmt.Println(sayHello("Максим", 13))
	printMessage("вызов 1")
	printMessage("вызов 2")
	printMessage("вызов 3")

	message, err := enterTheClub(18)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return
	}
	fmt.Println(message)
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Привет, %s! Тебе %d лет.", name, age)
}

// func enterTheClub(age int) (string, bool) {
// 	if age >= 18 {
// 		return "Входи, только аккуратно.", true
// 	} else if age >= 45 {
// 		return "Вам точно тут понравится?", true
// 	} 
// 	return "Тебе еще рано.", false
// }

func enterTheClub(age int) (string, error) {
	if age >= 18 {
		return "Входи, только аккуратно.", nil
	} else if age >= 45 {
		return "Вам точно тут понравится?", errors.New("Староваты все-таки")
	} 
	return "Тебе еще рано.", errors.New("Подрасти, малой")
}