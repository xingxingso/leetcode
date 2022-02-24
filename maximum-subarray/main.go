/*
Package maximum_subarray
https://leetcode-cn.com/problems/maximum-subarray/

53. 最大子序和

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

提示：
	1 <= nums.length <= 3 * 10^4
	-10^5 <= nums[i] <= 10^5

进阶：
	如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

*/
package maximum_subarray

// --- 自己

/*
方法一: 动态规划

时间复杂度：O(n)，其中 n 为nums 数组的长度。我们只需要遍历一遍数组即可求得答案。
空间复杂度：O(1)。我们只需要常数空间存放若干变量。
*/
func maxSubArray(nums []int) int {
	sum := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(nums[i], nums[i]+sum) // 包含当前数的最大和
		ans = max(ans, sum)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法一: 前缀和

时间复杂度：O(n)。
空间复杂度：O(n)。
*/
func maxSubArrayS2(nums []int) int {
	prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	//fmt.Println(prefixSum)

	// 从后向前 找到最大差值
	biggest := prefixSum[len(prefixSum)-1]
	ans := prefixSum[len(prefixSum)-1]
	for i := len(prefixSum) - 2; i >= 0; i-- {
		sum := biggest - prefixSum[i]
		if sum > ans {
			ans = sum
		}
		if prefixSum[i] > biggest {
			biggest = prefixSum[i]
		}
	}

	return ans
}

// --- 官方

/*
方法二: 分治

时间复杂度：假设我们把递归的过程看作是一颗二叉树的先序遍历，那么这颗二叉树的深度的渐进上界为 O(logn)，
	这里的总时间相当于遍历这颗二叉树的所有节点，故总时间的渐进上界是 O(\sum_{i=1}^{\log n} 2^{i-1})=O(n)，故渐进时间复杂度为 O(n)。
空间复杂度：递归会使用 O(logn) 的栈空间，故渐进空间复杂度为 O(logn)。
*/
func maxSubArrayO2(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

type Status struct {
	lSum, rSum, mSum, iSum int
}
