package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "JavaScript", "Python", "Linux"}
	subjectslice := subjects[:3]
	subjectslice[0] = "Database"
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("======================")
	for i := 0; i < len(subjectslice); i++ {
		fmt.Println(subjectslice[i])
	}

}
