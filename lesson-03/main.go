package main

import "fmt"

func main() {

	fmt.Println(findMin(1, 2, 3, 4, 5, -1, -12, 1341, -1241, -12, 45))
	fmt.Println(findMax(1, 2, 3, 4, 5, -1, -12, 1341, -1241, -12, 45))

	//Анонимная функция
	func() {
		fmt.Println("анонимная функция")
	}()

	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}

	return min
}

func findMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]

	for _, i := range numbers {
		if i > max {
			max = i
		}
	}

	return max
}

// func calendar(dayOfWeek string) string {
// 	if dayOfWeek == "пн" {
// 		return "Желаю тебе хорошего начала недели"
// 	} else if dayOfWeek == "вт" {
// 		return "Хорошего вторника"
// 	} else if dayOfWeek == "ср" {
// 		return "Хорошей среды"
// 	} else if dayOfWeek == "чт" {
// 		return "Хорошего четверга"
// 	}
// 	return ""
// }

// func prediction(dayOfWeek string) (string, error) {
// 	switch dayOfWeek {
// 	case "пн":
// 		return "Хорошего начала недели", nil
// 	case "вт":
// 		return "Хорошего вторника", nil
// 	case "ср":
// 		return "Хорошей среды", nil
// 	case "чт":
// 		return "Хорошего четверга", nil
// 	case "пт":
// 		return "Хорошей пятницы", nil
// 	case "сб":
// 		return "Хорошей субботы", nil
// 	case "вс":
// 		return "Хорошего конца недели", nil
// 	default:
// 		return "невалидный день недели", errors.New("invalid day of the week")
// 	}
// }
