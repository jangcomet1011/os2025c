package main

import (
	"week10/pkg/greeting" // 커스텀 패키지(standard path)
	"week10/pkg/greeting/deutsch"
	"week10/pkg/greeting/korean"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.GutenTag()
	korean.Anyung()
}
