package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	total := 0

	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of gorountines: ", runtime.NumGoroutine())
	}()

	var mutex sync.Mutex
	for i:=0; i < 2; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()  // 如果不加，则会出现 goroutines 泄露
			total += 1
		}()
	}
}
