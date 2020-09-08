package main
import (
	"fmt"
	"math/rand"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	curNode := root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}

		top := stack[len(stack) - 1]
		res = append(res, top.Val)
		stack = stack[:len(stack) - 1]

		curNode = top.Right
	}

	return res
}

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			select {
			case ch <- rand.Intn(5):
			case <- done:
				return
			default:
			}
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
		}
		done <- true
		return
	}()

	<-done
}