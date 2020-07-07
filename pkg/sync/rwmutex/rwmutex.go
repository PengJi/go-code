package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MapCounter struct {
	m map[int]int
	sync.RWMutex
}

func (mapCounter *MapCounter) Reader(n int) {
	for {
		mapCounter.RLock()
		v := mapCounter.m[rand.Intn(n)]
		mapCounter.RUnlock()
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
}

func (mapCounter *MapCounter) Writer(n int) {
	for i := 0; i < n; i++ {
		mapCounter.Lock()
		mapCounter.m[i] = i * 10
		mapCounter.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	mc := MapCounter{m: make(map[int]int)}
	go mc.Writer(10)
	go mc.Reader(10)
	go mc.Reader(10)
	time.Sleep(15 * time.Second)
}
