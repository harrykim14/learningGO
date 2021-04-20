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
