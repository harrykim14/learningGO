// package main

import "fmt"

// 소문자로 작성하면 private의 효과를 갖는다
type Vertex struct {
	x, y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func Area(v Vertex) int {
	return v.x * v.y
}

// x, y를 받는 New함수(New는 디자인패턴)를 만들어 Vertex의 포인터를 리턴하도록 함
// 이 때, 리턴하는 것은 Vertex{x, y}로 만들어진 struct의 주소
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area()) // 1200
}
