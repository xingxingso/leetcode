/*
Package range_sum_query_mutable
https://leetcode-cn.com/problems/range-sum-query-mutable/

307. 区域和检索 - 数组可修改

给你一个数组 nums ，请你完成两类查询，其中一类查询要求更新数组下标对应的值，另一类查询要求返回数组中某个范围内元素的总和。

实现 NumArray 类：
	- NumArray(int[] nums) 用整数数组 nums 初始化对象
	- void update(int index, int val) 将 nums[index] 的值更新为 val
	- int sumRange(int left, int right) 返回子数组 nums[left, right] 的总和（即，nums[left] + nums[left + 1], ..., nums[right]）

提示：
	1 <= nums.length <= 3 * 10^4
	-100 <= nums[i] <= 100
	0 <= index < nums.length
	-100 <= val <= 100
	0 <= left <= right < nums.length
	最多调用 3 * 10^4 次 update 和 sumRange 方法

*/
package range_sum_query_mutable

//type NumArray struct {}
//func Constructor(nums []int) NumArray {}
//func (this *NumArray) Update(index int, val int)  {}
//func (this *NumArray) SumRange(left int, right int) int {}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

// --- 自己

/*
方法一: 暴力
	时间复杂度较高

时间复杂度：
空间复杂度：
*/

type NumArrayS1 struct {
	origin    []int
	prefixSum []int
}

func ConstructorS1(nums []int) NumArrayS1 {
	prefixSum := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	return NumArrayS1{
		origin:    nums,
		prefixSum: prefixSum,
	}
}

func (n *NumArrayS1) Update(index int, val int) {
	//fmt.Println(n.prefixSum)
	diff := n.origin[index] - val
	for i := index + 1; i < len(n.prefixSum); i++ {
		n.prefixSum[i] -= diff
	}

	n.origin[index] = val
	//fmt.Println(n.prefixSum)
}

func (n *NumArrayS1) SumRange(left int, right int) int {
	return n.prefixSum[right+1] - n.prefixSum[left]
}

/*
方法一: 线段树

时间复杂度：
空间复杂度：
*/

type NumArray struct {
	data []int
	tree []int
}

func Constructor(nums []int) NumArray {
	tree := make([]int, 4*len(nums))
	var dfs func(index, left, right int)
	dfs = func(index, left, right int) {
		if left == right {
			tree[index] = nums[left]
			return
		}
		mid, leftIndex, rightIndex := left+(right-left)>>1, 2*index+1, 2*index+2
		dfs(leftIndex, left, mid)
		dfs(rightIndex, mid+1, right)
		tree[index] = tree[leftIndex] + tree[rightIndex]
		return
	}
	dfs(0, 0, len(nums)-1)
	//fmt.Println(tree)
	return NumArray{
		data: nums,
		tree: tree,
	}
}

func (n *NumArray) Update(index int, val int) {
	var dfs func(treeIndex, left, right, index, val int)
	dfs = func(treeIndex, left, right, index, val int) {
		if left == right {
			n.tree[treeIndex] = val
			return
		}
		mid, leftIndex, rightIndex := left+(right-left)>>1, 2*treeIndex+1, 2*treeIndex+2
		if index > mid {
			dfs(rightIndex, mid+1, right, index, val)
		} else {
			dfs(leftIndex, left, mid, index, val)
		}
		n.tree[treeIndex] = n.tree[leftIndex] + n.tree[rightIndex]
	}

	dfs(0, 0, len(n.data)-1, index, val)
	//fmt.Println(n.tree)
}

func (n *NumArray) SumRange(left int, right int) int {
	var dfs func(index, l, r, left, right int) int
	dfs = func(index, l, r, left, right int) int {
		if l == left && r == right {
			return n.tree[index]
		}

		mid, leftIndex, rightIndex := l+(r-l)>>1, 2*index+1, 2*index+2
		if right <= mid {
			return dfs(leftIndex, l, mid, left, right)
		}
		if left > mid {
			return dfs(rightIndex, mid+1, r, left, right)
		}
		return dfs(leftIndex, l, mid, left, mid) + dfs(rightIndex, mid+1, r, mid+1, right)
	}

	return dfs(0, 0, len(n.data)-1, left, right)
}
