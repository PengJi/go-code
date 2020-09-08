package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func test1() {
	var opts int64 = 0

	for i := 0; i < 50; i++ {
		// 第一个参数必须是地址
		atomic.AddInt64(&opts, 3)  // 加操作
		time.Sleep(time.Millisecond)
	}

	time.Sleep(time.Second)

	fmt.Println("opts: ", atomic.LoadInt64(&opts))
}

func test2() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops: ", ops)
}

func main() {
	test2()
}
