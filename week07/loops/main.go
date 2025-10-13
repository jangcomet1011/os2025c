package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	i, _ := r.ReadString('\n') // 에러를 무시함 "_" 는
	fmt.Println(i)
}
