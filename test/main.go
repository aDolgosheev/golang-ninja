package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...

    bytes := []byte(text)
	var res string


	if len(bytes) <= width {
		res = text
	} else {
		bytes = bytes[0:width]
		res = string(bytes)
		res = res + "..."
	}
    
	fmt.Println(res)
}