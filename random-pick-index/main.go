/*
Package random_pick_index
https://leetcode-cn.com/problems/random-pick-index/

398. 随机数索引

给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。

注意：
	数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。

*/
package random_pick_index

import (
	"math/rand"
	"time"
)

//type Solution struct {}
//func Constructor(nums []int) Solution {}
//func (this *Solution) Pick(target int) int {}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

// --- 他人
//https://leetcode-cn.com/problems/random-pick-index/solution/sui-ji-shu-suo-yin-ji-yu-yu-dao-de-targe-41xg/

/*
方法一:

时间复杂度：
空间复杂度：
*/

type SolutionS1 struct {
	data []int
	r    *rand.Rand
}

func ConstructorS1(nums []int) SolutionS1 {
	return SolutionS1{
		data: nums,
		r:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *SolutionS1) Pick(target int) int {
	count := 0
	i := 0
	index := -1
	for i < len(this.data) {
		if this.data[i] == target {
			count++
			if this.r.Intn(count)+1 == count {
				index = i
			}
		}
		i++
	}
	return index
}
