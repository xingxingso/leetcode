/*
Package container_with_most_water
https://leetcode-cn.com/problems/container-with-most-water/

11. 盛最多水的容器

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。

提示：
	n == height.length
	2 <= n <= 10^5
	0 <= height[i] <= 10^4

*/
package container_with_most_water

// --- 官方

/*
方法一: 双指针

时间复杂度：
空间复杂度：
*/
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		ans = max(min(height[l], height[r])*(r-l), ans)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
