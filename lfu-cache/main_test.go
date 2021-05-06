package lfu_cache

import (
	"testing"
)

func TestLFUCache(t *testing.T) {
	lfu := Constructor(2)
	testEx1(&lfu, t)
	//lfu2 := Constructor(0)
	//testEx2(&lfu2, t)
}

//func TestLFUCacheO1(t *testing.T) {
//	lfuO1 := ConstructorO1(2)
//	testEx1(&lfuO1, t)
//}

type LFU interface {
	Get(key int) int
	Put(key int, value int)
}

func testEx1(lfu LFU, t *testing.T) {
	lfu.Put(1, 1)
	lfu.Put(2, 2)

	if got := lfu.Get(1); got != 1 {
		t.Errorf("Get() = %v, want %v", got, 1)
	}
	// cache=[1,2], cnt(2)=1, cnt(1)=2

	lfu.Put(3, 3) // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
	// cache=[3,1], cnt(3)=1, cnt(1)=2

	if got := lfu.Get(2); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}

	if got := lfu.Get(3); got != 3 {
		t.Errorf("Get() = %v, want %v", got, 3)
	}
	// cache=[3,1], cnt(3)=2, cnt(1)=2

	//printLFUCache(lfu.(*LFUCache))
	lfu.Put(4, 4) // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
	// cache=[4,3], cnt(4)=1, cnt(3)=2

	if got := lfu.Get(1); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}

	if got := lfu.Get(3); got != 3 {
		t.Errorf("Get() = %v, want %v", got, 3)
	}
	// cache=[3,4], cnt(4)=1, cnt(3)=3

	if got := lfu.Get(4); got != 4 {
		t.Errorf("Get() = %v, want %v", got, 4)
	}
	// cache=[3,4], cnt(4)=2, cnt(3)=3
}

func testEx2(lfu LFU, t *testing.T) {
	lfu.Put(0, 0)

	if got := lfu.Get(0); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}
	// cache=[]
}
