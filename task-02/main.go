package main

import (
	"errors"
	"fmt"
)

func main() {
	message, err := enterTheClub(12)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(message)
}

func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Входите, только аккуратно!", nil
	} else if age >= 45 && age < 65 {
		return "Вам точно понравится эта музыка?", nil
	} else if age >= 65 {
		return "Это для вас уже слишком", errors.New("you are too old")
	}
	return "Вход запрещен", errors.New("you are too young")
}
