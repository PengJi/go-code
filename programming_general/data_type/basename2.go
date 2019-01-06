package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//basename2 reads file names from stdin and prints the base name of each one

func main(){
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		fmt.Println(basename(input.Text()))
	}
}

// basename removes directory components and a trailing .suffix
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string{
	slash := strings.LastIndex(s, "/")  //-1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >=0{
		s = s[:dot]
	}

	return s
}
