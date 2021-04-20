// package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	// 버퍼를 2개로 설정했으므로 ch는 2개의 값만을 받을 수 있다

	x := <-ch      // 여기서 ch 내 버퍼에서 값을 하나 끄집어 냄
	fmt.Println(x) // 100

	fmt.Println(len(ch)) // ch의 길이는 1이 된다

	ch <- 300            // 값을 하나 집어넣으면
	fmt.Println(len(ch)) // 2가 된다

	close(ch)
	//len 이 있다는 뜻은 채널은 순환 가능한 iterator이다
	for c := range ch {
		fmt.Println(c)
	}
	// 하지만 이렇게 하면 오류가 남 -> 순환하기 전에 채널을 닫아주어야 함

}
