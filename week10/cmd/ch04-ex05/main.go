package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard" //go get github.com/headfirstgo/keyboard
)

func main() {
	fmt.Print("실수 입력 : ")
	score, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f\n", score)
}
