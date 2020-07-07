package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * 条件等待
 */

func main() {
	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			// Wait()等待通知: 阻塞当前线程，直到收到该条件变量发来的通知
			cond.Wait()
			wg.Done()
			// do other
			fmt.Println(i)
		}(i)
	}

	fmt.Println("正被阻塞。。。")

	time.Sleep(time.Second * 1)

	// Signal()单发通知: 让该条件变量向至少一个正在等待它的通知的线程发送通知，表示共享数据的状态已经改变。
	cond.Signal()

	fmt.Println("通知已被释放")

	time.Sleep(time.Second * 1)

	fmt.Println("广播")

	// Broadcast广播通知: 让条件变量给正在等待它的通知的所有线程都发送通知。
	cond.Broadcast()

	wg.Wait()

}
