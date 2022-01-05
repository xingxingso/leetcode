/*
Package random_pick_with_weight
https://leetcode-cn.com/problems/random-pick-with-weight/

528. 按权重随机选择

给你一个 下标从 0 开始 的正整数数组 w ，其中 w[i] 代表第 i 个下标的权重。

请你实现一个函数 pickIndex ，它可以 随机地 从范围 [0, w.length - 1] 内（含 0 和 w.length - 1）选出并返回一个下标。选取下标 i 的 概率 为 w[i] / sum(w) 。

例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3) = 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3) = 0.75（即，75%）。

提示：
	1 <= w.length <= 10^4
	1 <= w[i] <= 10^5
	pickIndex 将被调用不超过 10^4 次

*/
package random_pick_with_weight

import (
	"math/rand"
	"sort"
)

//type Solution struct {}
//func Constructor(w []int) Solution {}
//func (this *Solution) PickIndex() int {}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

// --- 自己

/*
方法一:
	转化为前缀和，之后可以使用二分查找

时间复杂度：
空间复杂度：
*/

type Solution struct {
	preSum []int
}

func Constructor(w []int) Solution {
	prefixSum := make([]int, len(w))
	for sum, i := 0, 0; i < len(w); i++ {
		sum += w[i]
		prefixSum[i] = sum
	}

	return Solution{
		preSum: prefixSum,
	}
}

// PickIndex 问题转化为求 preSum 中 最左侧大于等于 rand 的下标
func (s *Solution) PickIndex() int {

	// 影响性能
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//pick := r.Intn(s.preSum[len(s.preSum)-1]) + 1 // [0,n) + 1 => [1, n]

	pick := rand.Intn(s.preSum[len(s.preSum)-1]) + 1 // [0,n)

	//index := s.searchBiggerMinIndex(s.preSum, pick)

	index := sort.SearchInts(s.preSum, pick)
	return index
}

// searchBiggerMinIndex 寻找升序数组中第一个大于等于目标值的数字的下标
func (s *Solution) searchBiggerMinIndex(nums []int, target int) int {
	var searchIndex func(left, right int) int
	searchIndex = func(left, right int) int {
		for left <= right {
			mid := left + (right-left)/2
			if target > nums[mid] {
				left = mid + 1
			} else if target < nums[mid] {
				right = mid - 1
			} else {
				right = mid - 1
			}
		}
		if left >= len(nums) {
			return -1
		}
		return left
	}

	return searchIndex(0, len(nums)-1)
}
