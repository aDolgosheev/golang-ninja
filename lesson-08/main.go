package main

import (
	"fmt"
	"math"
)

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	//square := Square{5}
	//circle := Circle{8}

	//printShapeArea(square)
	//printShapeArea(circle)

	//printInterface(square)
	//printInterface(circle)
	printInterface("12341234124")
	printInterface(22)
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

func printInterface(i interface{}) {
	//switch value := i.(type) {
	//case int:
	//	fmt.Println("int", value)
	//case bool:
	//	fmt.Println("bool", value)
	//default:
	//	fmt.Println("unknown type", value)
	//}
	//fmt.Println(i)

	str, ok := i.(string)
	if !ok {
		fmt.Println("interface is not string")
	}
	fmt.Println(len(str))
}
