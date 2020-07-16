package main

import "fmt"

// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html#1-%E5%86%99%E5%87%BA%E4%B8%8B%E9%9D%A2%E4%BB%A3%E7%A0%81%E8%BE%93%E5%87%BA%E5%86%85%E5%AE%B9

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

// 函数的return value 不是原子操作，而是在编译器中分解为两部分：返回值赋值 和 return 。
// 而defer刚好被插入到末尾的return前执行。故可以在derfer函数中修改返回值
func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			//将影响返回值
			score = source
		}
	}()
	return source * 2
	//上面等价于：
	//score = source * 2
	//return
}

func main() {
	//defer_call()
	fmt.Println(doubleScore(0))    //0
	fmt.Println(doubleScore(20.0)) //40
	fmt.Println(doubleScore(50.0)) //50
}