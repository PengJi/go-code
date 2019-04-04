package skiplist

import (
	"fmt"
	"math/rand"
)

const (
	//最高层数
	MAX_LEVEL = 16
)

//跳表节点结构体
type Node struct {
	//结点数值
	val int
	//层高
	level int
	//每层前进指针
	forwards []*Node
}

//跳表结构体
type SkipList struct {
	//跳表头结点
	head *Node
	//跳表当前层数
	level int
	//跳表长度
	length int
}

//新建跳表节点
func newNode(val, level int) *Node {
	return &Node{val: val, forwards: make([]*Node, level, level), level: level}
}

//实例化跳表对象
func NewSkipList() *SkipList {
	//头结点，便于操作
	head := newNode(0, MAX_LEVEL)
	return &SkipList{head, 1, 0}
}

//获取跳表长度
func (sl *SkipList) Length() int {
	return sl.length
}

//获取跳表层级
func (sl *SkipList) Level() int {
	return sl.level
}

//插入节点到跳表中
func (sl *SkipList) Insert(val int) int {
	//查找插入位置
	cur := sl.head
	//记录每层的路径
	update := [MAX_LEVEL]*Node{}
	i := MAX_LEVEL - 1

	for ; i >= 0; i-- {  //每次循环向下一层
		for cur.forwards[i] != nil {  //每次循环向右前进一步
			if cur.forwards[i].val == val {
				return 2
			}
			if cur.forwards[i].val > val {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	//通过随机算法获取该节点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	//创建一个新的跳表节点
	newNode := newNode(val, level)

	//原有节点连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	//如果当前节点的层数大于之前跳表的层数
	//更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	//更新跳表长度
	sl.length++

	return 0
}

//查找
func (sl *SkipList) Find(val int) *Node {
	if sl.length == 0 {
		return nil
	}

	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {  //每次循环向下一层
		for cur.forwards[i] != nil{  //每次循环向右移动一步
			if val == cur.forwards[i].val {
				return cur.forwards[i]
			} else if val < cur.forwards[i].val{
				fmt.Println(cur.forwards[i].val)
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

//查找
func (sl *SkipList) FindWithPath(val int) (*Node, []int) {
	var path []int

	if sl.length == 0 {
		return nil,nil
	}

	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {  //每次循环向下一层
		for cur.forwards[i] != nil {  //每次循环向右移动一步
			if val == cur.forwards[i].val {
				path = append(path, cur.forwards[i].val)
				return cur.forwards[i], path
			} else if val < cur.forwards[i].val{
				path = append(path, cur.forwards[i].val)
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil, nil
}

//删除节点
func (sl *SkipList) Delete(val int) int {
	//查找前驱节点
	cur := sl.head
	//记录前驱路径
	update := [MAX_LEVEL]*Node{}
	for i := sl.level - 1; i >= 0; i-- {  //向下移动一层
		update[i] = sl.head
		for nil != cur.forwards[i] {  //向右移动一步
			if cur.forwards[i].val == val {
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

		if nil == update[i].forwards[i] {
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

//func GetSkipList(sl *SkipList) []int {
//	//return fmt.Sprintf("level:%+v, length:%+v", sl.level, sl.length)
//	var path []int
//	cur := sl.head
//	fmt.Println(sl.length)
//	for i :=sl.length-1; i>=0; i-- {
//		fmt.Println(cur.forwards[0].val)
//		path = append(path, cur.forwards[0].val)
//		cur = cur.forwards[i]
//	}
//
//	return path
//}
//
//func main(){
//	sl := NewSkipList()
//
//	sl.Insert(95)
//	sl.Insert(88)
//	sl.Insert(100)
//	sl.Insert(101)
//
//	fmt.Println(GetSkipList(sl))
//}