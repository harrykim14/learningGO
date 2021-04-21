/*
"GO111MODULE"
기존 GOPATH와 vendor/에 따라 동작하던 go 커맨드와의 공존을 위한 GO111MODULE이라는 임시 환경변수가 생김(~ 1.11 ver.)
`go env -w GO111MODULE=on`은 GOPATH에 전혀 관계없이 Go modules의 방식대로 동작하고
`go env -w GO111MODULE=off`는 Go modules는 전혀 사용되지 않고 기존에 사용되던 방식대로 GOPATH와 verdor/를 통해 go 커맨드가 동작함
`go env -w GO111MODULE=auto`의 경우 GOPATH/src 내부에서의 go 커맨드는 기존의 방식대로 외부에서의 go 커맨드는 Go modules의 방식대로 동작함
*/

// package main

import (
	"210421/mylib"
	"210421/mylib/under"
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))
	mylib.Say()

	person := under.Person{Name: "Mike", Age: 20}
	fmt.Print(person.Name + ` `)
	under.Hello() // 함수의 첫글자를 소문자로 작성하면 private임을 잊지 말자
}
