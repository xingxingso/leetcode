package public

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap(t *testing.T) {
	h := IntHeap{}
	//h.Push(2)
	//h.Push(3)
	//h.Push(1)
	//h.Push(4)
	//h.Push(6)
	//for h.Len() > 0 {
	//	fmt.Println(h.Pop())
	//}

	heap.Init(h)
}
