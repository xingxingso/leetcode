/*
Package pancake_sorting
https://leetcode-cn.com/problems/pancake-sorting/

969. 煎饼排序

给你一个整数数组 arr ，请使用 煎饼翻转 完成对数组的排序。
一次煎饼翻转的执行过程如下：
	- 选择一个整数 k ，1 <= k <= arr.length
	- 反转子数组 arr[0...k-1]（下标从 0 开始）
例如，arr = [3,2,1,4] ，选择 k = 3 进行一次煎饼翻转，反转子数组 [3,2,1] ，得到 arr = [1,2,3,4] 。

以数组形式返回能使 arr 有序的煎饼翻转操作所对应的 k 值序列。任何将数组排序且翻转次数在 10 * arr.length 范围内的有效答案都将被判断为正确。

提示：
	1 <= arr.length <= 100
	1 <= arr[i] <= arr.length
	arr 中的所有整数互不相同（即，arr 是从 1 到 arr.length 整数的一个排列）

*/
package pancake_sorting

// --- 自己

/*
方法一: 深度优先遍历

时间复杂度：
空间复杂度：
*/
func pancakeSort(arr []int) []int {
	ans := make([]int, 0)
	var dfs func(arr []int)
	dfs = func(arr []int) {
		if len(arr) == 1 {
			return
		}
		maxIndex := getMaxIndex(arr)
		if maxIndex+1 != len(arr) { // 已处于最下层 没有必要
			if maxIndex != 0 { // reverse(arr, 0, 1) 没有必要
				reverse(arr, 0, maxIndex)
				ans = append(ans, maxIndex+1)
			}
			reverse(arr, 0, len(arr)-1)
			ans = append(ans, len(arr))
		}
		dfs(arr[:len(arr)-1])
	}

	dfs(arr)
	return ans
}

func reverse(arr []int, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func getMaxIndex(arr []int) int {
	maxIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}
