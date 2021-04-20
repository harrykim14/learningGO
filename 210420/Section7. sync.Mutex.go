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
