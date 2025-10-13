package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	//i, _ := r.ReadString('\n') // 에러를 무시함 "_" 는
	i, err := r.ReadString('\n')
	//fmt.Println(err)
	if err != nil {
		log.Fatal(err) //report 에러 그리고 프로그램을 끝냄
	}

	fmt.Println(i)
}
