package main

import "fmt"

//单向通道

//只能发送
func counter(out chan<-int){
	for x:=0; x<100; x++{
		out <- x
	}
	close(out)
}

//接收，发送
func squarer(out chan<-int, in <-chan int){
	for v:= range in{
		out <- v*v
	}

	close(out)
}

//只能接收
func printer(in <-chan int){
	for v := range in{
		fmt.Println(v)
	}
}

func main(){
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}