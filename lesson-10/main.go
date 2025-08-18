package main

import "fmt"

type Number interface {
	int64 | float64
}

type User struct {
	email string
	name string
}

func main() {

	a := []int64{1, 2, 3}
	b := []float64{1.1, 2.2, 3.3}
	c := []string{"1", "2", "3"}

	d := []User {
		{
			email: "asd@gmail.com",
			name: "Vasya",
		},
		{
			email: "wasd@ya.ru",
			name: "Petya",
		},
		{
			email: "dodo@mail.ru",
			name: "Sanya",
		},
	}

	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(searchElement(c, "2"))
	fmt.Println(searchElement(d, User{
		email: "wasd@ya.ru",
		name: "Petya",
	}))

	printAny(d)
}

func sum[V Number](input []V) V {
	var result V

	for _, number := range input {
		result += number
	}

	return result
}

// func sumOfFloat64(input []float64) float64 {
// 	var result float64

// 	for _, number := range input {
// 		result += number
// 	}

// 	return result
// }

func searchElement[C comparable](elements []C, searchEl C) bool {
	for _, el := range elements {
		if el == searchEl {
			return true
		}
	}

	return false
}

func printAny[A any](input A) {
	fmt.Println(input)
}
