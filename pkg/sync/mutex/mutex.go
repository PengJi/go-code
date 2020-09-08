package main

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	number int
	sync.Mutex
}

func (sc *safeCounter) Increment() {
	sc.Lock()
	defer sc.Unlock()
	sc.number++
}

func (sc *safeCounter) Decrement() {
	sc.Lock()
	defer sc.Unlock()
	sc.number--
}

func (sc *safeCounter) getNumber() int {
	sc.Lock()
	defer sc.Unlock()
	number := sc.number
	return number
}

func main() {
	sc := new(safeCounter)
	for i := 0; i < 100; i++ {
		go sc.Increment()
		go sc.Decrement()
	}
	fmt.Println(sc.getNumber())
}
