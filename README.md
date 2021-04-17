# GO를 배워보자

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
