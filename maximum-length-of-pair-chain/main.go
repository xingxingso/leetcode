/*
Package maximum_length_of_pair_chain
https://leetcode-cn.com/problems/maximum-length-of-pair-chain/

646. 最长数对链

给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。

现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。

给定一个数对集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。

提示：
	给出数对的个数在 [1, 1000] 范围内。

*/
package maximum_length_of_pair_chain

import (
	"sort"
)

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1] || (pairs[i][1] == pairs[j][1] && pairs[i][0] < pairs[j][0])
	})
	//fmt.Println(pairs)
	dp := make([]int, len(pairs))
	dp[0] = 1
	ans := 1
	for i := 1; i < len(pairs); i++ {
		dp[i] = 1
		// 二分查找, 找到最靠右的小于 pairs[i][0] 的 pairs[j][1], 其实不用也行，只是加快了一点速度
		for j := binarySearchRight(pairs[:i], pairs[i][0]); j >= 0; j-- {
			if s := dp[j] + 1; s > dp[i] { // 所有前面的都要考虑
				dp[i] = s
			}
			dp[i] = max(dp[i], dp[j]+1)
		}
		ans = max(dp[i], ans)
	}

	return ans
}

func binarySearchRight(pairs [][]int, target int) int {
	left, right := 0, len(pairs)-1
	for left <= right {
		mid := left + (right-left)/2
		if pairs[mid][1] == target {
			right = mid - 1
		} else if pairs[mid][1] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
