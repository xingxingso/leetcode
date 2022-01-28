/*
Package advantage_shuffle
https://leetcode-cn.com/problems/advantage-shuffle/

870. 优势洗牌

给定两个大小相等的数组 A 和 B，A 相对于 B 的优势可以用满足 A[i] > B[i] 的索引 i 的数目来描述。

返回 A 的任意排列，使其相对于 B 的优势最大化。

提示：
	1 <= A.length = B.length <= 10000
	0 <= A[i] <= 10^9
	0 <= B[i] <= 10^9

*/
package advantage_shuffle

import (
	"sort"
)

// --- 自己

/*
方法一: 暴力贪心
	// 超出时间限制

时间复杂度：
空间复杂度：
*/
func advantageCount(nums1 []int, nums2 []int) []int {
	for i := 0; i < len(nums2); i++ {
		idx := -1
		minIdx := i
		for j := i; j < len(nums1); j++ {
			if nums1[j] < nums1[minIdx] {
				minIdx = j
			}
			if nums1[j] > nums2[i] {
				if idx == -1 {
					idx = j
				} else {
					if nums1[idx] > nums1[j] {
						idx = j
					}
				}
			}
		}
		if idx == -1 {
			idx = minIdx
		}

		nums1[i], nums1[idx] = nums1[idx], nums1[i]
	}

	return nums1
}

/*
方法一: 暴力

时间复杂度：
空间复杂度：
*/
func advantageCountS2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	clone2 := make([]int, len(nums2))
	copy(clone2, nums2)
	sort.Ints(clone2)
	memo := make(map[int][]int)
	use := make(map[int]bool)
	for i, j := 0, 0; i < len(clone2); i++ {
		if j >= len(nums1) {
			continue
		}
		for j < len(nums1) {
			if nums1[j] > clone2[i] {
				use[j] = true
				memo[clone2[i]] = append(memo[clone2[i]], j)
				j++
				break
			}
			j++
		}
	}
	unuse := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		if !use[i] {
			unuse = append(unuse, i)
		}
	}

	//fmt.Println(memo)

	ans := make([]int, len(nums1))
	for i := 0; i < len(nums2); i++ {
		if len(memo[nums2[i]]) == 0 { // 最小的未使用的数
			ans[i] = nums1[unuse[len(unuse)-1]]
			unuse = unuse[:len(unuse)-1]
			continue
		}

		ans[i] = nums1[memo[nums2[i]][len(memo[nums2[i]])-1]]
		memo[nums2[i]] = memo[nums2[i]][:len(memo[nums2[i]])-1]
	}

	return ans
}
