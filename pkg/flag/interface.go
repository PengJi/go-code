package main

/*
创建自定义 flag，只要实现 flag.Value 接口即可（要求 receiver 是指针）
*/

import (
	"flag"
	"fmt"
	"strings"
)

//定义一个类型，用于增加该类型方法
type sliceValue []string

// new 一个存放命令行参数值的slice
func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

/*
Value接口：
type Value interface {
    String() string
    Set(string) error
}
实现flag包中的Value接口，将命令行接收到的值用,分隔存到slice里
*/
func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}

// flag为slice的默认值default is me
func (s *sliceValue) String() string {
	*s = sliceValue(strings.Split("default is me", ","))
	return "It's none of my business"
}

/*
可执行文件名 -slice="java,go"  最后将输出[java,go]
可执行文件名 最后将输出[default is me]
*/
func main() {
	var languages []string
	flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`")
	flag.Parse()

	// 打印结果slice接收到的值
	fmt.Println(languages)
}
