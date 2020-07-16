package main

import "fmt"

// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html#2-%E4%BB%A5%E4%B8%8B%E4%BB%A3%E7%A0%81%E6%9C%89%E4%BB%80%E4%B9%88%E9%97%AE%E9%A2%98-%E8%AF%B4%E6%98%8E%E5%8E%9F%E5%9B%A0

type student struct {
	Name string
	Age int
}

func parse_stu() map[string]*student {
	m := make(map[string]*student)

	stus := []student{
		{Name: "zhou", Age: 20},
		{Name: "li", Age: 21},
		{Name: "wang", Age: 22},
	}

	// for遍历时，变量stu指针不变，每次遍历仅进行struct值拷贝，故m[stu.Name]=&stu实际上一致指向同一个指针，最终该指针的值为遍历的最后一个struct的值拷贝。
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	// 类似
	var stu student
	for _, stu = range stus {
		m[stu.Name] = &stu
	}

	// 修正方案
	for i, _ := range stus {
		tmp := stus[i]
		m[tmp.Name] = &tmp
	}

	return m
}

func main() {
	students := parse_stu()

	for k, v := range students {
		fmt.Println(k, v)
	}
}
