package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	MAX_LEVEL = 16  //最高层数
)

// 跳表节点结构体
type skipListNode struct {
	v interface{}  //跳表保存的值
	score int  //用于排序的分值
	level int  //层高
	forwards []*skipListNode  //每层前进指针
}

// 跳表结构体
type SkipList struct{
	head *skipListNode  //跳表头结点
	level int  //跳表当前层数
	length int  //跳表长度
}

// 新建跳表节点
func newSkipListNode(v interface{}, score, level int ) *skipListNode {
	return &skipListNode{v: v, score: score, forwards: make([]*skipListNode, level, level), level: level}
}

// 新建跳表对象
func NewSkipList() *SkipList{
	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL)  //头结点，便于操作
	return &SkipList{head, 1, 0}
}

// 获取跳表长度
func (sl *SkipList) Length() int{
	return sl.length
}

func (sl *SkipList) Level() int{
	return sl.level
}

// 插入节点到跳表中
func (sl *SkipList) Insert(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	cur := sl.head  //查找插入位置
	update := [MAX_LEVEL]*skipListNode{}  //记录每层的路径
	i := MAX_LEVEL -1

	for ; i>=0; i-- {
		for cur.forwards[i] != nil{
			if cur.forwards[i].v == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	// 通过随机算法获取节点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1{
			level++
		}
	}

	// 创建一个新的跳表节点
	newNode := newSkipListNode(v, score, level)

	// 连接
	for i:=0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	// 如果当前节点的层数大于之前跳表的层数
	// 1. 更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}
	// 更新跳表长度
	sl.length++

	return 0
}

// 查找
func (sl *SkipList) Find(v interface{}, score int) *skipListNode {
	if v == nil || sl.length == 0 {
		return nil
	}

	cur := sl.head
	for i := sl.level-1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

// 删除节点
func (sl *SkipList) Delete(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	cur := sl.head  //查找前驱节点
	update := [MAX_LEVEL] * skipListNode{}  //记录前驱路径
	for i := sl.level-1; i >= 0; i-- {
		update[i] = sl.head
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}

		if update[i].forwards[i] == nil {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	sl.length--

	return 0
}

func (sl *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", sl.level, sl.length)
}