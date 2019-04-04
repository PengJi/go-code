package skiplist

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	sl := NewSkipList()

	sl.Insert(95)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")

	sl.Insert(88)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")

	sl.Insert(100)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")

	sl.Insert(99)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0].forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")

	sl.Insert(105)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0].forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0].forwards[0].forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")

	sl.Insert(70)
	t.Log(sl)
	//t.Log(GetSkipList(sl))
	//fmt.Println(GetSkipList(sl))
	sl.Insert(115)
	t.Log(sl)
	sl.Insert(130)
	t.Log(sl)

	t.Log(sl.Find(88))
	t.Log("-----------------------------")

	t.Log(sl.FindWithPath(115))
	t.Log("-----------------------------")

	t.Log(sl.FindWithPath(130))
	t.Log("-----------------------------")

	sl.Delete(95)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")
}