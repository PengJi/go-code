package main

import (
	"fmt"
	"sync"
)

// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html#8-%E4%B8%8B%E9%9D%A2%E7%9A%84%E4%BB%A3%E7%A0%81%E6%9C%89%E4%BB%80%E4%B9%88%E9%97%AE%E9%A2%98

//type UserAges struct {
//	ages map[string]int
//	sync.Mutex
//}
//
//func (ua *UserAges) Add(name string, age int) {
//	ua.Lock()
//	defer ua.Unlock()
//	ua.ages[name] = age
//}
//
//func (ua *UserAges) Get(name string) int {
//	if age, ok := ua.ages[name]; ok {
//		return age
//	}
//
//	return -1
//}

//虽然有使用sync.Mutex做写锁，但是map是并发读写不安全的。
//map属于引用类型，并发读写时多个协程见是通过指针访问同一个地址，即访问共享变量，此时同时读写资源存在竞争关系。
//会报错误信息:“fatal error: concurrent map read and map write”。

type UserAges struct {
	ages map[string]int
	sync.RWMutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.RLock()
	defer ua.RUnlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}

	return -1
}

func main() {
	count := 10000
	gw := sync.WaitGroup{}
	gw.Add(count * 3)
	u := UserAges{ages: map[string]int{}}

	add := func(i int) {
		u.Add(fmt.Sprintf("user_%d", i), i)
		gw.Done()
	}

	for i := 0; i < count; i++ {
		go add(i)
		go add(i)
	}

	for i := 0; i < count; i++ {
		go func(i int) {
			defer gw.Done()
			u.Get(fmt.Sprintf("user_%d", i))
			fmt.Print(".")
		}(i)
	}

	gw.Wait()
	fmt.Println("Done")
}
