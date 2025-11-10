package main

import "fmt"

func main() {
	subject := make([]string, 3)
	subject[0] = "Go"
	subject[2] = "Python"

	for _, subject := range subject {
		fmt.Println(subject)
	}
}
