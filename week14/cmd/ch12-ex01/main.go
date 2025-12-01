package main

import "fmt"

func main() {
	//Last in First out
	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	defer fmt.Println("3rd defer")
	fmt.Println("main logic")
}
