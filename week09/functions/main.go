package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func swap(first *int, second *int) {
	var temp int = 0
	temp = *first
	*first = *second
	*second = temp
	fmt.Println(*first, *second)
}
func main() {
	//fmt.Printf("%0.3f\n", math.Sqrt(-25.0))
	//a, b := 10, 20
	//fmt.Println(a, b)
	//swap(&a, &b)
	//fmt.Println(a, b)
	fmt.Println("점수입력")
	score, err := GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	passFail := ""
	if score >= 60 {
		passFail = "pass!"
	} else {
		passFail = "Fail!"
	}
	fmt.Printf("%0.2f점은 %v\n", score, passFail)
}
