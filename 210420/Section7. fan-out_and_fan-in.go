// fan-out & fan-in Pattern
// goroutine1(producer)
// -(channel)-> goroutine2(2nd Stage)
// -(channel)-> goroutine3(3rd Stage)
// -(channel)-> func main()

// package main

import "fmt"

func producer(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func mulit4(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	go multi2(first, second)
	go multi2(second, third)

	for result := range third {
		fmt.Println(result)
	}
}
