package main

import "fmt"

// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html#4-%E4%B8%8B%E9%9D%A2%E4%BB%A3%E7%A0%81%E4%BC%9A%E8%BE%93%E5%87%BA%E4%BB%80%E4%B9%88

type People struct{}

func (p *People) ShowA() {
	fmt.Println("show a")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("show b")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher show b")
}

func main() {
	t := Teacher{}
	t.ShowA()
	t.ShowB()
}
