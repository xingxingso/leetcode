/*
Package queue_reconstruction_by_height
https://leetcode-cn.com/problems/queue-reconstruction-by-height/

406. 根据身高重建队列

假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。
每个 people[i] = [h_i, k_i] 表示第 i 个人的身高为 h_i ，前面 正好 有 k_i 个身高大于或等于 h_i 的人。

请你重新构造并返回输入数组 people 所表示的队列。
返回的队列应该格式化为数组 queue ，其中 queue[j] = [h_j, k_j] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。

提示：
	1 <= people.length <= 2000
	0 <= h_i <= 10^6
	0 <= k_i < people.length
	题目数据确保队列可以被重建

*/
package queue_reconstruction_by_height

import (
	"sort"
)

// --- 官方

/*
方法一: 从低到高考虑

时间复杂度：O(n^2)，其中 n 是数组 people 的长度。我们需要 O(nlogn) 的时间进行排序，
    随后需要 O(n^2) 的时间遍历每一个人并将他们放入队列中。
    由于前者在渐近意义下小于后者，因此总时间复杂度为 O(n^2)。
空间复杂度：O(logn)，即为排序需要使用的栈空间。
*/
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	ans := make([][]int, len(people))
	for _, person := range people {
		spaces := person[1] + 1
		for i := range ans {
			if ans[i] == nil {
				spaces--
				if spaces == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}
