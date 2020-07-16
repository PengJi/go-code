package main

import (
	"fmt"
	"runtime"
)

// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html#5-%E4%B8%8B%E9%9D%A2%E4%BB%A3%E7%A0%81%E4%BC%9A%E8%A7%A6%E5%8F%91%E5%BC%82%E5%B8%B8%E5%90%97-%E8%AF%B7%E8%AF%A6%E7%BB%86%E8%AF%B4%E6%98%8E

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	int_chan <- 1
	string_chan <- "hello"

	// 单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：
	//
	// select 中只要有一个case能return，则立刻执行。
	// 当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
	// 如果没有一个case能return则可以执行”default”块。
	select {
	case value := <- int_chan:
		fmt.Println(value)
	case value := <- string_chan:
		panic(value)
	}
}
