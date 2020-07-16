package main

// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html#9-%E4%B8%8B%E9%9D%A2%E7%9A%84%E8%BF%AD%E4%BB%A3%E4%BC%9A%E6%9C%89%E4%BB%80%E4%B9%88%E9%97%AE%E9%A2%98

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()
		for elem := range set.s {
			ch <- elem
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

// 内部迭代出现阻塞。默认初始化时无缓冲区，需要等待接收者读取后才能继续写入。
