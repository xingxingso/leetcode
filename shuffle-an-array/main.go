/*
Package shuffle_an_array
https://leetcode-cn.com/problems/shuffle-an-array/

384. 打乱数组

给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。

实现 Solution class:
	Solution(int[] nums) 使用整数数组 nums 初始化对象
	int[] reset() 重设数组到它的初始状态并返回
	int[] shuffle() 返回数组随机打乱后的结果

提示：
	1 <= nums.length <= 200
	-10^6 <= nums[i] <= 10^6
	nums 中的所有元素都是 唯一的
	最多可以调用 5 * 10^4 次 reset 和 shuffle

*/
package shuffle_an_array

import (
	"math/rand"
	"time"
)

//type Solution struct {}
//func Constructor(nums []int) Solution {}
//func (this *Solution) Reset() []int {}
//func (this *Solution) Shuffle() []int {}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

// --- 自己

/*
方法一: Fisher–Yates shuffle(洗牌算法)

时间复杂度：
空间复杂度：
*/

type Solution struct {
	origin []int
	upset  []int
}

func Constructor(nums []int) Solution {
	temp := make([]int, len(nums))
	copy(temp, nums)

	return Solution{
		origin: nums,
		upset:  temp,
	}
}
func (s *Solution) Reset() []int {
	return s.origin
}
func (s *Solution) Shuffle() []int {
	seed := time.Now().UnixNano() // 这个会影响性能
	r := rand.New(rand.NewSource(seed))
	//r := rand.New(rand.NewSource(96))
	//fmt.Println(seed)
	for i := len(s.upset) - 1; i >= 0; i-- { // 0 的时候可以交换了
		k := r.Intn(i + 1)
		//k := rand.Intn(i + 1) // 没有seed 的时候随机性其实是固定的
		//fmt.Println(k, i+1)
		s.upset[i], s.upset[k] = s.upset[k], s.upset[i]
	}
	return s.upset
}
