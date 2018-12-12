package main

/*
https://yar999.gitbooks.io/gopl-zh/content/ch4/ch4-03.html
通过map来标识所有的输入行对应的set集合，
以确保已经在集合存在的行不会被重复打印
*/

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    seen := make(map[string]bool) //a set of strings
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            seen[line] = true
            fmt.Println(line)
        }
    }

    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)
    }
}
