// package main

// import (
// 	"fmt"
// )

// func one(x *int) {
// 	*x = 1
// }

// func main() {
// 	var n int = 100
// 	fmt.Println(n)  // 100
// 	fmt.Println(&n) // 0xc000014088
// 	var p *int = &n
// 	fmt.Println(p)  // 0xc000014088
// 	fmt.Println(*p) // 100

// 	one(&n)
// 	fmt.Println(n)
// 	fmt.Println(&n)

// 	var p1 *int = new(int)
// 	fmt.Println(p1)
// 	fmt.Println(*p1)
// 	var p2 *int
// 	fmt.Println(p2)
// 	fmt.Printf("%T\n", p2)

// 	m := make(map[string]int)
// 	fmt.Printf("%T\n", m)
// }
