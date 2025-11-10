package main

import "fmt"

func main() {
	subject := []string{"Go", "", "Python"}

	for _, subject := range subject {
		fmt.Println(subject)
	}
}
