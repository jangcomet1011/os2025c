package korean

import "fmt"

func Anyunghaseyo() {
	fmt.Println("안녕하세요!")
	//anyung()
}

func Anyung() { //함수명이 소문자로 시작하는 경우 외부 파일에서 접근불가
	fmt.Println("안녕!")
}
