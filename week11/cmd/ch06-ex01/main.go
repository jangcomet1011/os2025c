package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "JavaScript", "Python", "Linux"}
	subjectslice := subjects[:3]
	subjects[0] = "Java"
	subjectslice = append(subjectslice, "Go")
	//subjectslice = append(subjectslice, "Go", "DB")
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("======================")
	for i := 0; i < len(subjectslice); i++ {
		fmt.Println(subjectslice[i])
	}

}
