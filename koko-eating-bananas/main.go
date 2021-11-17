/*
Package koko_eating_bananas
https://leetcode-cn.com/problems/koko-eating-bananas/

875. 爱吃香蕉的珂珂

珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。
珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，
她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。

提示：
	1 <= piles.length <= 10^4
	piles.length <= H <= 10^9
	1 <= piles[i] <= 10^9

*/
package koko_eating_bananas

import (
	"math"
	"sort"
)

// --- 自己

/*
方法一: 二分查找
	使用 golang 的 Search

时间复杂度：O(NlogW)，其中 N 是香蕉堆的数量，W 最大的香蕉堆的大小。
空间复杂度：O(1)。
*/
func minEatingSpeed(piles []int, h int) int {
	// 这里是为了获取 max, 也可以不排序直接 用  10^9 或者一遍循环找到最大的
	sort.Slice(piles, func(i, j int) bool {
		return piles[i] > piles[j]
	})

	return sort.Search(piles[0], func(i int) bool {
		times := 0
		if i <= 0 {
			return false
		}
		for j := 0; j < len(piles); j++ {
			// n := piles[j] / i
			// if n*i < piles[j] {
			// 	n++
			// }
			// times += n
			times += int(math.Ceil(float64(piles[j]) / float64(i)))
		}

		return times <= h
	})
}

/*
方法一: 二分查找
	自己实现 Search

时间复杂度：O(NlogW)，其中 N 是香蕉堆的数量，W 最大的香蕉堆的大小。
空间复杂度：O(1)。
*/
func minEatingSpeedS1(piles []int, h int) int {
	var possible func(k int) bool
	possible = func(k int) bool {
		times := 0
		for i := 0; i < len(piles); i++ {
			times += int(math.Ceil(float64(piles[i]) / float64(k)))
		}
		return times <= h
	}

	lo, hi := 1, 1000000000
	for lo < hi {
		mi := int(uint(lo+hi) >> 1) // 这里参考了 golang 的 sort.Search
		if possible(mi) {
			hi = mi
		} else {
			lo = mi + 1
		}
		// fmt.Println(hi)
	}
	// return lo
	return hi
}
