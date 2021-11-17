/*
Package sort_characters_by_frequency
https://leetcode-cn.com/problems/sort-characters-by-frequency/

451. 根据字符出现频率排序

给定一个字符串，请将字符串里的字符按照出现的频率降序排列。

*/
package sort_characters_by_frequency

import "container/heap"

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func frequencySort(s string) string {
	occurrences := map[byte]int{}
	for i := 0; i < len(s); i++ {
		occurrences[s[i]]++
	}

	h := &IHeap{}
	heap.Init(h)
	for k, v := range occurrences {
		heap.Push(h, [2]int{int(k), v})
	}
	ret := make([]byte, 0, len(s))
	for h.Len() > 0 {
		x := heap.Pop(h).([2]int)
		for i := 0; i < x[1]; i++ {
			ret = append(ret, byte(x[0]))
		}
	}

	return string(ret)
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
