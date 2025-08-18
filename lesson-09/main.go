package main

import (
	"fmt"
	"golang-ninja/basic/shape"
)

func main() {
	square := shape.NewSquare(5)
	//circle := Circle{8}

	printShapeArea(square)
	//printShapeArea(circle)

	//printInterface(square)
	//printInterface(circle)
	// printInterface("12341234124")
	// printInterface(22)

	// scheduler := scheduler.NewScheduler()
}

func printShapeArea(s shape.Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}

// func printInterface(i interface{}) {
// 	//switch value := i.(type) {
// 	//case int:
// 	//	fmt.Println("int", value)
// 	//case bool:
// 	//	fmt.Println("bool", value)
// 	//default:
// 	//	fmt.Println("unknown type", value)
// 	//}
// 	//fmt.Println(i)

// 	str, ok := i.(string)
// 	if !ok {
// 		fmt.Println("interface is not string")
// 	}
// 	fmt.Println(len(str))
// }
