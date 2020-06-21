package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello")
}

// exported
var Greeter greeting
