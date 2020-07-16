package main

import (
	"fmt"
	"runtime"
	"sync"
)

// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html#3-%E4%B8%8B%E9%9D%A2%E7%9A%84%E4%BB%A3%E7%A0%81%E4%BC%9A%E8%BE%93%E5%87%BA%E4%BB%80%E4%B9%88-%E5%B9%B6%E8%AF%B4%E6%98%8E%E5%8E%9F%E5%9B%A0

func main() {
	// 将GOMAXPROCS设置为1，将影响goroutine的并发，后续代码中的go func()相当于串行执行。
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)

	// go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。
	// 故go func执行时，i的值始终是10（10次遍历很快完成）。
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}

	// go func中i是函数参数，与外部for中的i完全是两个变量。
	// 尾部(i)将发生值拷贝，go func内部指向值拷贝地址。
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("j: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}