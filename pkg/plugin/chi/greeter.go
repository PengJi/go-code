package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("你好")
}

// exported
var Greeter greeting
