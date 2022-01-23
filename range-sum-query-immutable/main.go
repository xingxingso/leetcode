/*
Package range_sum_query_immutable
https://leetcode-cn.com/problems/range-sum-query-immutable/

303. 区域和检索 - 数组不可变

给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。

实现 NumArray 类：

NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）

提示：
	0 <= nums.length <= 104
	-105 <= nums[i] <= 105
	0 <= i <= j < nums.length
	最多调用 104 次 sumRange 方法

*/
package range_sum_query_immutable

//type NumArray struct {}
//func Constructor(nums []int) NumArray {}
//func (this *NumArray) SumRange(left int, right int) int {}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	//fmt.Println(prefixSum)
	return NumArray{
		prefixSum: prefixSum,
	}
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.prefixSum[right+1] - n.prefixSum[left]
}
