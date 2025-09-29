package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var id int16
	//var name string
	//var gpa float64

	//id = 999
	//name = "jallyyam"
	//gpa = 3.99

	//fmt.Println("학번은 ", id, ", 이름은", name)
	//fmt.Println("평점 : ", gpa)

	//var id int16 = 999
	//var name string = "jallyyam"
	//var gpa float32 = 3.99

	//fmt.Println("학번은 ", id, ", 이름은", name)
	//fmt.Println("평점 : ", gpa)

	//var id = 999
	//var name = "jallyyam"
	//var gpa = 3.99

	//id := 999
	//name := "jallyyam"
	//gpa := 3.99

	//fmt.Println("학번은 ", id, reflect.TypeOf(id), ", 이름은", name, reflect.TypeOf(name))
	//fmt.Println("평점 : ", gpa, reflect.TypeOf(gpa))

	// zero values
	//var f64 float64
	//var t bool
	///var s string
	//var i int
	//var i16 int16

	//fmt.Println(f64, reflect.TypeOf(f64))
	//fmt.Println(t, reflect.TypeOf(t))
	//fmt.Println(s, reflect.TypeOf(s))
	//fmt.Println(i, reflect.TypeOf(i))
	//fmt.Println(i16, reflect.TypeOf(i16))

	var f64 float64
	//total_price := 1000
	//totalPrice := 1000
	totalPrice := 1000

	fmt.Println(totalPrice)
	fmt.Println(totalPrice, reflect.TypeOf(f64))
}
