/*
Package top_k_frequent_elements
https://leetcode-cn.com/problems/top-k-frequent-elements/

347. 前 K 个高频元素

给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

提示：
	1 <= nums.length <= 10^5
	k 的取值范围是 [1, 数组中不相同的元素的个数]
	题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

进阶：
	你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。

*/
package top_k_frequent_elements

import (
	"container/heap"
	"sort"
)

// --- 自己

/*
方法一: 桶排序

时间复杂度：
空间复杂度：
*/
func topKFrequent(nums []int, k int) []int {
	h := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		h[nums[i]]++
	}
	s := make([]int, 0, k)
	for _, v := range h {
		s = append(s, v)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	m := make(map[int]bool)
	for _, v := range s[:k] {
		m[v] = true
	}
	ans := make([]int, 0, k)
	for k, v := range h {
		if m[v] {
			ans = append(ans, k)
		}
	}
	return ans
}

/*
方法一: 桶排序

时间复杂度：
空间复杂度：
*/
func topKFrequentS2(nums []int, k int) []int {
	h := make(map[int]int)
	max := 0
	for _, v := range nums {
		h[v]++
		if max < h[v] {
			max++
		}
	}
	s := make([][]int, max+1)
	for k, v := range h {
		s[v] = append(s[v], k)
	}
	ans := make([]int, 0, k)
	for i := max; i > 0 && k > 0; i-- {
		ans = append(ans, s[i]...) // 题目数据保证答案唯一
		k = k - len(s[i])
	}
	return ans
}

// --- 官方

/*
方法一:堆

时间复杂度：O(N log k)，其中 N 为数组的长度。我们首先遍历原数组，并使用哈希表记录出现次数，每个元素需要 O(1) 的时间，共需 O(N) 的时间。
	随后，我们遍历「出现次数数组」，由于堆的大小至多为 k，因此每次堆操作需要 O(log k) 的时间，共需 O(N log k) 的时间。
	二者之和为 O(N log k)
空间复杂度：O(N)。哈希表的大小为 O(N)，而堆的大小为 O(k)，共计为 O(N)。
*/
func topKFrequentO1(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
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
