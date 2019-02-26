package tree

import "fmt"

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{NewNode(rootV)}
}

func (this *BinaryTree) InOrderTraverse() {
	p := this.root
	s := NewArrayStack()

	for !s.IsEmpty() || p != nil {
		if p != nil {
			s.Push(p)
			p = p.left
		} else {
			tmp := s.Pop().(*Node)
			fmt.Printf("%+v", tmp.data)
			p = tmp.right
		}
	}
	fmt.Println()
}

func (this *BinaryTree) PreOrderTraverse() {
	p := this.root
	s := NewArrayStack()

	for !s.IsEmpty() || p != nil {
		if p != nil {
			fmt.Printf("%+v", p.data)
			s.Push(p)
			p = p.left
		} else {
			p = s.Pop().(*Node).right
		}
	}
	fmt.Println()
}

func (this *BinaryTree) PostOrderTraverse() {
	s1 := NewArrayStack()
	s2 := NewArrayStack()
	s1.Push(this.root)
	for !s1.IsEmpty() {
		p := s1.Pop().(*Node)
		s2.Push(p)
		if p.left != nil {
			s1.Push(p.left)
		}
		if p.right != nil {
			s1.Push(p.right)
		}
	}

	for !s2.IsEmpty() {
		fmt.Printf("%+v", s2.Pop().(*Node).data)
	}
}

func (this *BinaryTree) PostOrderTraverse2() {
	r := this.root
	s := NewArrayStack()

	//point to last visit node
	var pre *Node

	for !s.IsEmpty() {
		r = s.Top().(*Node)
		if (r.left == nil && r.right == nil) || (pre != nil && (pre == r.left || pre == r.right)){
			fmt.Printf("%+v", r.data)
			s.Pop()
			pre = r
		} else {
			if r.right != nil {
				s.Push(r.right)
			}

			if r.left != nil {
				s.Push(r.left)
			}
		}
	}
}