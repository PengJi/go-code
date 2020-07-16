package main

import (
	"fmt"
)

// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html#10-%E4%BB%A5%E4%B8%8B%E4%BB%A3%E7%A0%81%E8%83%BD%E7%BC%96%E8%AF%91%E8%BF%87%E5%8E%BB%E5%90%97-%E4%B8%BA%E4%BB%80%E4%B9%88

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	// var peo People = Stduent{}
	// 上述应修改为如下
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

// func (stu *Stduent) Speak(think string) (talk string) 是表示结构类型*Student的指针有提供该方法，但该方法并不属于结构类型Student的方法。因为struct是值类型。
