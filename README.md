# Go를 배워보자

- **강의명** : 현역 실리콘밸리 엔지니어가 가르치는 Go입문 + 응용으로 비트코인 시스템 트레이딩 핀테크 어플리케이션의 개발 ([現役シリコンバレーエンジニアが教える Go 入門 + 応用でビットコインのシストレ Fintech アプリの開発](https://www.udemy.com/course/go-fintech/))

- **수강 기간** : 21. 04. 16 ~ 현재

- **수강 목적**

  1. JavaScript외에도 타 언어를 배우고자 함
  2. 왜 python도 아니고 Go를? → Go의 문법이 자바스크립트와 일부 유사하고 특히 함수가 일급객체라는 것이 매우 닮았다 ([참고자료](https://lannex.github.io/blog/2019/Golang-for-JavaScript-developers-1/))
  3. 동시성을 제어하는 고루틴과 채널 개념을 익히기 위해
  4. 구글이 만든 언어이니 앞으로 더 쓰임새가 많아질 언어임이 분명
  5. 프론트엔드 기술만 배워서는 좋은 웹 엔지니어가 될 수 없고 넓은 시야가 있어야 지금 사용하는 언어의 깊이 또한 깊어질 수 있을 것이라는 믿음 때문에

## 1. Go의 기본적인 형태

<details>
<summary> 열기 </summary>
<div markdonw="1">

```go
package main // package 형 언어

import "fmt" // 이와 같은 import 형태임

func main() { // C나 java같은 main 함수가 있어야 함
    // 1. 변수 선언
    var i int = 1

    var {
        j int = 2
        s string = "test"
        t, f bool = true, false
    }

    xi := 1 // 축약
    xt, xf : = true, false

    // 2. 배열과 슬라이스
    var a [2]int = [2]int{10, 20} // 배열 (불가변)
    var b []int = []int{100, 200} // 슬라이스 (가변)
    n := []int{1, 2, 3, 4, 5, 6}
    fmt.printLn(n[2:4])
    k := make([]int, 3, 5) // make(형, 길이, 메모리)

    // 3. map
    m := map[string]int{"apple": 100, "banana": 200}
    v, ok := m["apple"] // 두 번째 인자로 해당 값의 유무를 판별 가능, 이 경우 100 true가 출력될 것

    // 4. 함수
    r := add(10, 20)

    conter := incrementGenerator()
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2

}

func add(x int, y int) result int { // 매개변수들의 형과 반환형을 기입할 것
    result = x + y // 이렇게 써도 이미 result를 반환하기로 되어 있으므로 이 값을 알아서 반환
    return
}

func incrementGenerator() (func() int) {
    x := 0
    return func() int {
        x++
        return x
    }
}
// 자바스크립트에서 사용하는 spread operator로 여러개의 인자를 받을 수 있음 (아예 안 받을 수도 있음)
func sum() (params ...int) result int {
    for _, param := range params {
        result += param
    }
    return
}

```

</div>
</details>

## 2. Go의 문법과 포인터

<details>
<summary> (2-1) 문법 예제 코드 </summary>
<div markdonw="2-1">

```go
package main // package 형 언어


import "fmt"

func main() {
	num := 6
	if num%2 == 0 { // JS와는 달리 괄호 없이 조건문이 들어감
		fmt.Println("by 2")
	} else {
		fmt.Println("else")
	}

	for i := 0; i < 6; i++ { // 당연히 조건문 안에서도 형정의의 축약형을 사용할 수 있다
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	l := []string{"python", "java", "go"}

	for _, v := range l { // 이 부분은 파이썬과 비슷함
		fmt.Print(v + ` `)
	}

	os := "window"
    // 스위치도 괄호만 없고 똑같지만 대신 break가 없다
	switch os {
	case "mac":
		fmt.Println("Mac!")
	case "window":
		fmt.Println("Window!")
	default:
		fmt.Println("default...")
	}

}

```

</div>
</details>

<details>
<summary> (2-2) 포인터 예제 코드 </summary>
<div markdonw="2-2">

```go
package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	fmt.Println(n) // 100
	fmt.Println(&n) // 0xc000014088
	var p *int = &n
	fmt.Println(p) // 0xc000014088
	fmt.Println(*p) // 100

    one(&n)
	fmt.Println(n) // 1
	fmt.Println(&n) // 0xc000014088

    var p1 *int = new(int)
	fmt.Println(p1) // 0xc0000140a8
    fmt.Println(*p1) // 0
	var p2 *int
	fmt.Println(p2) // <nil>
    fmt.Printf("%T\n", p2) // *int

    m := make(map[string]int)
	fmt.Printf("%T\n", m)
    // fmt.Printf("%T\n", *m)
    // *m 을 보려고 하면 invalid operation: cannot indirect m (variable of type map[string]int)라고 표시된다
    // make로 생성된 자료구조는 포인터가 존재하지 않음
}
```

- struct 예제

```go
package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
	S string
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v) // {1 2 }

	v.X = 100
	fmt.Println(v.X, v.Y) // 100 2

	v2 := Vertex{X: 1}
	fmt.Println(v2) // {1 0 }

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3) // {1 2 test}

	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4) // main.Vertex {0 0 }

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5) // main.Vertex {0 0 }

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6) // *main.Vertex &{0 0 }

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7) //*main.Vertex &{0 0 }
}

```

</div>
</details>

<details>
<summary> (2-3) Go의 메소드와 포인터 리시버</summary>
<div markdonw="2-3">

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

// Vertex형의 v라는 변수에 Area()를 연결 (Go 메소드)
func (v Vertex) Area() int {
	return v.X * v.Y
}

// Vertex의 주소 내에서 직접 조작 (포인터 리시버)
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))  // 12
	fmt.Println(v.Area()) // 12

	v.Scale(10)
	fmt.Println(v.Area()) // 1200
}

```

</div>
</details>

<details>
<summary> (2-4) struct의 캡슐화</summary>
<div markdonw="2-4">

```go
package main

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

```

</div>
</details>

<details>
<summary> (2-5) Embeded </summary>
<div markdonw="2-5">

```go
package main

import "fmt"

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

type Vertex3D struct {
	Vertex // super() 같은 효과
	z      int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area3D()) // 30 * 40 * 50 = 60000
}
```

</div>
</details>

<details>
<summary> (2-6) non-struct 메소드 </summary>
<div markdonw="2-6">

```go
package main

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	fmt.Printf("%T %v\n", i, i) // main.MyInt 10
	fmt.Printf("%T %v\n", 1, 1) // int 1
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double()) // 20
}
```

</div>
</details>

<details>
<summary> (2-7) 인터페이스와 타입 단언(Type Assertion) </summary>
<div markdonw="2-7">

```go
// package main

import "fmt"

type Human interface {
	Say()
}

type Person struct {
	Name string
}

func (p *Person) Say() {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
}

func main() {
	var mike Human = &Person{"Mike"}
	mike.Say()
}

/* -------------------------------- */

package main

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
```

</div>
</details>

<details>
<summary> (2-8) String()과 Error() </summary>
<div markdonw="2-8">

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// fmt에 있는 String()을 오버로딩 한 것
func (p Person) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike.String())
}

/* -------------------------------- */
package main

import "fmt"

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}

```

</div>
</details>

## 3. Go의 goroutine과 channel

<details>
<summary> (3-1) goroutine과 sync.WaitGroup </summary>
<div markdown="3-1">

```go
package main

import (
	"fmt"
	// "time"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go goroutine("world")
	normal("hello")

	wg.Wait()
}

```

</div>
</details>

<details>
<summary> (3-2) channel과 buffered channel </summary>
<div markdown="3-2">

- sync.WaitGroup를 사용해 기다리지 않아도 channel을 통해 루틴간의 통신이 가능
- goroutine의 실행 순서가 일정하지 않아 출력값 15와 120이 실행할 때마다 달라질 수 있다

```go
package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	mul := 1
	for _, v := range s {
		mul *= v
	}
	c <- mul
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)

	go goroutine1(s, c)
	go goroutine2(s, c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}
```

- 채널을 만들 때 make 함수의 두번째 인자로 버퍼를 설정할 수 있는데 이 버퍼의 갯수에 따라 받을 수 있는 인자의 수가 결정된다

```go
package main

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
```

- goroutine과 channel을 써서 합을 구하는 과정을 출력하기

```go
package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))

	go goroutine1(s, c)
	for i := range c {
		fmt.Println(i)
	}
}
```

- goroutine과 channel을 사용한 Producer/Consumer 패턴

```go
package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process", i*1000)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
}
```

- goroutine과 channel을 사용한 fan-out/fan-in 패턴
  - 이 패턴을 사용한 예시 : 유저 정보를 받아 정보를 처리하고 이메일을 보내는 파이프라인 작업을 수행 등

```go
package main

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

```

- select를 사용한 병렬처리

```go
package main

import (
	"fmt"
	"time"
)

func goroutine1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(3 * time.Second)
	}
}

func goroutine2(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go goroutine1(c1)
	go goroutine2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
```

- default selection과 for break

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)

OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom!")
			break OuterLoop
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

</div>
</details>

<details>
<summary> (3-3) sync.Mutex</summary>
<div markdown="3-3">

- 여러 goroutine에서 같은 객체를 공유한다면 어떤 상황에서는 동시간에 해당 객체를 참조할 수 있는데 이럴 경우에 오류가 발생 할 수 있음
- 이를 안전하게 처리하기 위한 동기화 객체가 Mutex

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

// 같은 키값을 동시에 읽으면 오류남
func main() {
	c := Counter{v: make(map[string]int)}

	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c.Value("key"))
}
```

</div>
</details>

## 4. package, public과 private 그리고 테스트 모듈과 실제 테스트

<details>
<summary> (4-1) package, public과 private </summary>
<div markdown="4-1">

```go
/*
"GO111MODULE"
기존 GOPATH와 vendor/에 따라 동작하던 go 커맨드와의 공존을 위한 GO111MODULE이라는 임시 환경변수가 생김(~ 1.11 ver.)
`go env -w GO111MODULE=on`은 GOPATH에 전혀 관계없이 Go modules의 방식대로 동작하고
`go env -w GO111MODULE=off`는 Go modules는 전혀 사용되지 않고 기존에 사용되던 방식대로 GOPATH와 verdor/를 통해 go 커맨드가 동작함
`go env -w GO111MODULE=auto`의 경우 GOPATH/src 내부에서의 go 커맨드는 기존의 방식대로 외부에서의 go 커맨드는 Go modules의 방식대로 동작함
*/

package main

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
/*-----------------------------------------------------*/
// /PATH/mylib/math.go
package mylib

func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
/*-----------------------------------------------------*/
// /PATH/mylib/say.go
package mylib

import "fmt"

func Say() {
	fmt.Println("Hello, world!")
}
/*-----------------------------------------------------*/
// /PATH/mylib/under/hello.go
package under

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Hello() {
	fmt.Println("Hello!")
}
```

</div>
</details>

<details>
<summary> (4-2) 테스트 모듈과 실제 테스트 </summary>
<div markdown="4-2">

- 파일 내에 testing 모듈을 import하고 테스트할 함수의 변수로 `t *testing.T`를 넣는다

```go
// mylib/math_test.go
package mylib

import "testing"

func TestAverage(t *testing.T) {
	v := Average([]int{1, 2, 3, 4, 5})

	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}

```

- vscode에서 모듈을 테스트하는 방법에는 두가지 방법이 있는데 하나는 터미널에서 `go test ./...`와 같이 디렉토리 전체를 테스트하는 것이고 나머지는 launch.json에 Launch test function를 명시하고 기본 매개변수로 해당 파일을 넘겨주는 것

```json
// launch.json
"configurations": [
	{
        "name": "Launch test function",
        "type": "go",
        "request": "launch",
        "mode": "test",
        "program": "${workspaceFolder}/210421/mylib",
        "args": [
            "-test.v",
        ]
    },
	// 이 밑에는 launch 항목이 따로 있음
]
```

**실행 예제**

- 터미널에서 실행했을 때

```cli
PS %GOPATH%src> go test 210421/mylib
ok      210421/mylib    0.085s
```

- F5로 Launch test function을 실행했을 때 (디버그 콘솔)

```
API server listening at: 127.0.0.1:19984
=== RUN   TestAverage
--- PASS: TestAverage (0.00s)
PASS
```

- go의 테스팅은 매우 기본적인 것만 제공하므로 Ginkgo 혹은 Gomega의 테스트 라이브러리를 사용할 수도 있다

</div>
</details>

<details>
<summary> (4-3) 그 외 알아두면 좋은 것들 </summary>
<div markdown="4-3">

- gofmt : eslint나 prettier 같은 코드 스타일 정리 커맨드
- [서드 파티 패키지 검색 페이지](https://pkg.go.dev/?utm_source=godoc)
- 강의 내에서 talib 패키지를 다운로드

```
go get github.com/markcheno/go-talib
go get github.com/markcheno/go-quote
```

- 예제 코드를 실행해보기

```go
package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
```

- 실행 결과

```
datetime,open,high,low,close,volume
2016-01-04 00:00,200.49,201.03,198.59,181.92,222353500.00
2016-01-05 00:00,201.40,201.90,200.05,182.23,110845800.00
2016-01-06 00:00,198.34,200.06,197.60,179.93,152112600.00
2016-01-07 00:00,195.33,197.44,193.59,175.61,213436100.00
...
2016-03-31 00:00,205.91,206.41,205.33,186.95,94584100.00
```

- godoc (go의 공식문서 커맨드) : `go get golang.org/x/tools/cmd/godoc`로 다운로드하면 언제든 공식 문서를 참조 할 수 있음

</div>
</details>

## ext) 편리한 패키지와 네트워크에 관련된 라이브러리

<details>
<summary> (ex-1) 편리한 패키지 </summary>
<div markdown="ex-1">

1. time

```go
/*
const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)
*/
package sublib

import (
	"fmt"
	"time"
)

func TimeModuleExample() {
	t := time.Now()
	fmt.Println(t) // 2021-04-22 11:04:35.6950007 +0900 KST m=+0.002823501
	fmt.Println(t.Format(time.RFC3339)) // 2021-04-22T11:04:35+09:00
	fmt.Println(t.Year(), t.Month(), t.Day()) // 2021 April 22
}
```

2. regexp

```go
package sublib

import (
	"fmt"
	"regexp"
)

func RegexpExample() {
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match) // true

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms) // true

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs) // /view/test

	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2]) // [/view/test view test] /view/test view test
}
```

3. sort

```go
package sublib

import (
	"fmt"
	"sort"
)

func SortArr() {
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}

	fmt.Println(i, s, p) // [5 3 2 8 7] [d a f] [{Nancy 20} {Vera 40} {Mike 30} {Bob 50}]
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	fmt.Println(i, s, p) // [2 3 5 7 8] [a d f] [{Bob 50} {Mike 30} {Nancy 20} {Vera 40}]
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p) // [2 3 5 7 8] [a d f] [{Nancy 20} {Mike 30} {Vera 40} {Bob 50}]


}
```

4. iota

```go
package sublib

import "fmt"

const (
	c1 = iota
	c2
	c3
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
	TB
)

func IotaExample() {
	fmt.Println(c1, c2, c3) // 0 1 2
	fmt.Println(KB, MB, GB, TB) // 1024 1048576 1073741824 1099511627776
}
```

5. context

```go
package sublib

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(4 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func ContextExample() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	go longProcess(ctx, ch)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			// 함수의 실행 시간이 3초를 넘어가므로 이 케이스가 실행됨
			// context deadline exceeded
			break CTXLOOP
		case <-ch:
			fmt.Println("Success")
			break CTXLOOP
		}
	}
}

```

6. ioutil

```go
package sublib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func IoutilExample() {
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content)) // Example text

	r := bytes.NewBuffer([]byte("abc"))
	content2, _ := ioutil.ReadAll(r)
	fmt.Println(string(content2)) // abc
}
```

</div>
</details>

<details>
<summary> (ex-2) 네트워크 관련 패키지 </summary>
<div markdown="ex-2">

**1. http**

- http에서 리퀘스트를 보낼 때 NewRequest(method string, url string, body io.Reader)를 사용함

- \*http.Client의 주소값을 갖는 http.Client{} struct를 정의하고 이 객체를 통해 .Do, .Get, .Post 등을 실행할 수 있다

```go
package sub

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpExample() {
	// resp, _ := http.Get("https://nextjs-api-harry.herokuapp.com/api/detail-post/1/")
	// defer resp.Body.Close()

	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	fmt.Println(base) // http://example.com
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint) // http://example.com/test?a=1&b=2

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `W/"wyzzy`)
	q := req.URL.Query()
	fmt.Println(q) // map[a:[1] b:[2]]
	q.Add("c", "3")
	fmt.Println(q, q.Encode()) // map[a:[1] b:[2] c:[3]], a=1&b=2&c=3

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

**2. json.Unmarshal, json.Marshal**

```go
package sub

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"`      // 가리고 싶을 땐 "-"로 표기
	Age       int      `json:"age"`       // int를 string로 표기하고 싶을 땐 ,와 함께 string을 정의
	Nicknames []string `json:"nicknames"` // omitempty로 해당 값이 없다면 표기하지 않을 수도 있다
}

func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

func JsonExample() {
	b := []byte(`{"name":"Mike", "age":20, "nicknames":["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}

	fmt.Println(p.Name, p.Age, p.Nicknames) // Mike 20 [a b c]

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
	// before `json`setting : {"Name":"Mike","Age":20,"Nicknames":["a","b","c"]}
	// after : {"name":"Mike","age":20,"nicknames":["a","b","c"]}
	// even after custom Marshal func => {"Name":"Mr.Mike"}
}
```

**3. hmac**

- [go의 공식문서](https://golang.org/pkg/crypto/hmac/)에서 추천하는 hmac의 사용법

```go
// ValidMAC reports whether messageMAC is a valid HMAC tag for message.
func ValidMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
```

- 이를 토대로 간이 인증 서버를 구현

```go
package sub

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

func Server(apiKey, sign string, data []byte) bool {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	if sign == expectedHMAC {
		return true
	} else {
		return false
	}
}

func APIAuthexample() {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	data := []byte("data")
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	sign := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sign) // 076b55e7f7e126...

	fmt.Println(Server(apiKey, sign, data)) // true
}
```

</div>
</details>

<details>
<summary> (ex-3) 편리한 서드파티 패키지 </summary>
<div markdown="ex-3">

**1. Semaphore**

```go
package sub

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

/*
세마포어(Semaphore) : 공유된 자원의 데이터를 여러 프로세스가 접근하는 것을 막는 것
뮤텍스(Mutex) : 공유된 자원의 데이터를 여러 쓰레드가 접근하는 것을 막는 것
*/

var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}
	// if err := s.Acquire(ctx, 1); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	defer s.Release(1)

	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func SemaphoreExample() {
	ctx := context.TODO()

	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(2 * time.Second)

	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
```

**2. ini**

```go
// ini.go
package sub

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      int
	DBname    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DBname:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func IniExample() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DBname, Config.DBname)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
```

**3. talib**

```go
package sub

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func TalibExample() {
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2018-04-01", "2019-01-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	// Rsi란 종가의 변화로 추세 강도를 측정하는 선행지표
	// 헌재 추세강도가 어떠한지를 0~100퍼센트의 수치로 보여줌
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
	// Ema란 지수이동평균으로 과거 값(여기서는 14일)을 계산대상으로 단기변동성을 알 수 있는 값
	// Wma는 가중이동평균으로 현재의 추세를 알 수 있는 값
	mva := talib.Ema(spy.Close, 14)
	fmt.Println(mva)
}
```

**4. websocket으로 실시간 비트코인 가격 가져오기**

```go
package sub

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type JsonRPC2 struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Result  interface{} `json:"result,omitempty"`
	Id      *int        `json:"id,omitempty"`
}
type SubscribeParams struct {
	Channel string `json:"channel"`
}

func Websocket() {
	u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path: "/json-rpc"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	if err := c.WriteJSON(&JsonRPC2{Version: "2.0", Method: "subscribe", Params: &SubscribeParams{"lightning_ticker_BTC_JPY"}}); err != nil {
		log.Fatal("subscribe:", err)
		return
	}

	for {
		message := new(JsonRPC2)
		if err := c.ReadJSON(message); err != nil {
			log.Println("read:", err)
			return
		}

		if message.Method == "channelMessage" {
			log.Println(message.Params)
		}
	}
}
```

</div>
</details>
