package main

import (
	"fmt"
)

func swap(first int, second int) {
	var temp int = 0
	temp = first
	first = second
	second = temp
	fmt.Println(first, second)
}
func main() {
	//fmt.Printf("%0.3f\n", math.Sqrt(-25.0))
	a, b := 10, 20
	fmt.Println(a, b)
	swap(a, b)
	fmt.Println(a, b)
}
