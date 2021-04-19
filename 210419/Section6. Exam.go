package main

import "fmt"

type Vertex struct {
	X, Y int
}

/* Q1
func (v Vertex) Plus() int {
	return v.X + v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Plus())
}
*/

// Q2
func (v Vertex) String() string {
	return fmt.Sprintf("X is %v! Y is %v!", v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
}
