package main

import (
	"fmt"
	"runtime"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n  // 可能造成泄露
		}
		close(out)
	}()
	return out
}

func gen_no_leak(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	// Set up the pipeline.
	done := make(chan struct{})
	defer close(done)

	out := gen_no_leak(done, 2, 3)

	for n := range out {
		fmt.Println(n) // 2
		time.Sleep(5 * time.Second) // done thing, 可能异常中断接收
		if true { // if err != nil
			break
		}
	}
}
