package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func handle() {
	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		fmt.Println("1")
		wg.Done()
	}()

	go func() {
		fmt.Println(2)
		wg.Done()
	}()

	go func() {
		fmt.Println(3)
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of gorountines: ", runtime.NumGoroutine())
	}()

	go handle()

	time.Sleep(time.Second)
}