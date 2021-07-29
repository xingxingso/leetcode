/*
Package kth_largest_element_in_an_array
https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

215. 数组中的第K个最大元素

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

提示：
	1 <= k <= nums.length <= 10^4
	-10^4 <= nums[i] <= 10^4

*/
package kth_largest_element_in_an_array

import "container/heap"

// --- 自己

// --- 他人
//https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247488820&idx=1&sn=e6a58b67b0050ae8144bb8ea579cf0d0&scene=21#wechat_redirect

/*
方法一: 二叉堆

时间复杂度：O(NlogK)
空间复杂度：O(K)
*/
func findKthLargestP1(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return h.Peek().(int)
}

//IntHeap 小顶堆类
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//Peek 获取堆顶元素 不弹出
func (h *IntHeap) Peek() interface{} { return (*h)[0] }

/*
方法二: 快速选择算法

时间复杂度：O(N)
空间复杂度：O(1)
*/
func findKthLargestP2(nums []int, k int) int {
	left, right := 0, len(nums)-1
	k = len(nums) - k
	for left <= right {
		p := partition(nums, left, right)
		if p < k {
			left = p + 1
		} else if p > k {
			right = p - 1
		} else {
			return nums[p]
		}
	}

	return -1
}

func partition(nums []int, left, right int) int {
	if left == right {
		return left
	}
	// 将 nums[left] 作为默认分界点 pivot
	pivot := nums[left]
	// j = right + 1 因为 for 中会先执行 --
	i, j := left, right+1
	for {
		// 保证 nums[left..i] 都小于 pivot
		for i++; nums[i] < pivot; i++ {
			if i == right {
				break
			}
		}
		// 保证 nums[j..right] 都大于 pivot
		for j--; nums[j] > pivot; j-- {
			if j == left {
				break
			}
		}
		if i >= j {
			break
		}
		// 如果走到这里，一定有：
		// nums[i] > pivot && nums[j] < pivot
		// 所以需要交换 nums[i] 和 nums[j]，
		// 保证 nums[left..i] < pivot < nums[j..right]
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 将 pivot 值交换到正确的位置
	nums[j], nums[left] = nums[left], nums[j]
	// 现在 nums[lo..j-1] < nums[j] < nums[j+1..right]
	return j
}
