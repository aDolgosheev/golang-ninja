package main

import "fmt"

func main() {
	users := map[string]int{
		"Vasya":  15,
		"Petya":  23,
		"Kostya": 48,
	}

	// users["Serega"] = 21

	// age, exists := users["Kostya"]
	// if exists {
	// 	fmt.Println("Kostya", age)
	// } else {
	// 	fmt.Println("нет в списке")
	// }

	// delete(users, "Vasya")

	// for key, value := range users {
	// 	fmt.Println(key, value)
	// }

	fmt.Println(len(users))
}
