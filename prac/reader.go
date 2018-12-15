package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")

	//以 \n 为分隔符读取一段内容
	input, err := inputReader.ReadString('\n')
	if err != nil{
		fmt.Printf("Found an error: %s\n", err)
	}else{
		//对input进行切片操作，去掉内容中最后一个字节\n
		input = input[:len(input)-1]
		fmt.Printf("hello, %s!\n", input)
	}
}
