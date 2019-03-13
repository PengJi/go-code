package linkedlist

import "fmt"

// 单链表节点
type ListNode struct {
	next *ListNode
	value interface{}
}

// 单链表
type LinkedList struct {
	head *ListNode
	length uint
}

// 新建单链表节点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// 在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}

	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

// 在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil || p == this.head {
		return false
	}

	cur := this.head.next
	pre := this.head

	for nil != cur {
		if cur == p {
			break
		}

		pre = cur
		cur = cur.next
	}

	if nil == cur {
		return false
	}

	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

// 在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i<index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}

	cur := this.head.next
	pre := this.head

	for cur != nil {
		if p == cur {
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	pre.next = p.next
	p = nil
	this.length--
	return true
}

// 打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

// 单链表反转
func (this *LinkedList) Reverse(){
	if this.head == nil || this.head.next == nil || this.head.next.next == nil {
		return
	}

	var pre *ListNode = nil
	cur := this.head.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}

	this.head.next = pre
}