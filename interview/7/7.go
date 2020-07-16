package main

import "fmt"

// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html#7-%E8%AF%B7%E5%86%99%E5%87%BA%E4%BB%A5%E4%B8%8B%E8%BE%93%E5%85%A5%E5%86%85%E5%AE%B9

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

//make可用于初始化数组，第二个可选参数表示数组的长度。数组是不可变的。
//
//当执行make([]int,5)时返回的是一个含义默认值(int的默认值为0)的数组:[0,0,0,0,0]。
//而append函数是便是在一个数组或slice后面追加新的元素，并返回一个新的数组或slice。
//
//这里append(s,1,2,3)是在数组s的继承上追加三个新元素:1、2、3，故返回的新数组为[0 0 0 0 0 1 2 3]
