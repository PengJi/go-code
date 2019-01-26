package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pengji/go-code/programming_general/func/links"
)

// crawl2 crawls web links starting with the command-line arguments.
//
// The version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}  //acquire a token
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	worklist := make(chan []string)
	var n int  //number of pending sends to worklist

	// start with the command-line arguments.
	n++
	go func(){worklist <- os.Args[1:]}()

	// crawl the web concurently
	seen := make(map[string] bool)
	for ; n>0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}