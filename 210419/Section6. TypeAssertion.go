// package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) { // switch-type문
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't knwo %T\n", v)
	}
}

func main() {
	var i interface{} = 10 // i는 int형이 아닌 인터페이스임
	do(i)                  // 20
	do("Mike")             // Mike!
	do(true)               // I don't knwo bool
}
