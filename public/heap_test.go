package public

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap(t *testing.T) {
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, 2)
	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 4)
	heap.Push(h, 6)
	for h.Len() > 0 {
		fmt.Println(h.Peek())
		fmt.Println(heap.Pop(h))
	}
}
