/*
https://leetcode-cn.com/problems/meeting-rooms-ii/

253. 会议室 II

给你一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，为避免会议冲突，
同时要考虑充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。

提示：
	1 <= intervals.length <= 104
	0 <= starti < endi <= 106
*/
package meeting_rooms_ii

import (
	"container/heap"
	"sort"
)

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	count := 0
	mark := make(map[int]bool, len(intervals))
	for i := 0; i < len(intervals); i++ {
		if mark[i] {
			continue
		}
		mark[i] = true
		count++
		last := i
		for j := i + 1; j < len(intervals); j++ {
			if mark[j] {
				continue
			}
			if intervals[j][0] >= intervals[last][1] {
				mark[j] = true
				last = j
			}
		}
	}
	return count
}

// --- 官方

// https://leetcode-cn.com/problems/meeting-rooms-ii/solution/hui-yi-shi-ii-by-leetcode/
/*
方法一: 优先队列

时间复杂度：O(NlogN) 。
	- 时间开销主要有两部分。第一部分是数组的 排序 过程，消耗 O(NlogN) 的时间。数组中有 N 个元素。
	- 接下来是 最小堆 占用的时间。在最坏的情况下，全部 N 个会议都会互相冲突。在任何情况下，我们都要向堆执行 N 次插入操作。
      在最坏的情况下，我们要对堆进行 N 次查找并删除最小值操作。总的时间复杂度为 (NlogN) ，
      因为查找并删除最小值操作只消耗 O(logN) 的时间。
空间复杂度：O(N) 。额外空间用于建立 最小堆 。在最坏的情况下，堆需要容纳全部 N 个元素。因此空间复杂度为 O(N) 。
*/
func minMeetingRoomsO1(intervals [][]int) int {
	h := &IntHeap{}
	heap.Init(h)
	res := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); i++ {
		for len(*h) != 0 && intervals[i][0] >= (*h)[0] {
			heap.Pop(h)
		}
		heap.Push(h, intervals[i][1])
		res = max(res, len(*h))
	}
	return res
}

// 小顶堆类
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法二：有序化

时间复杂度：O(NlogN) 。我们所做的只是将 开始时间 和 结束时间 两个数组分别进行排序。每个数组有 N 个元素，因为有 N 个时间间隔。
空间复杂度：O(N) 。我们建立了两个 N 大小的数组。分别用于记录会议的开始时间和结束时间。
*/
func minMeetingRoomsO2(intervals [][]int) int {
	// Check for the base case. If there are no intervals, return 0
	if len(intervals) == 0 {
		return 0
	}

	start := make([]int, len(intervals))
	end := make([]int, len(intervals))

	for i := 0; i < len(intervals); i++ {
		start[i] = intervals[i][0]
		end[i] = intervals[i][1]
	}
	// Sort the intervals by start time
	sort.Slice(start, func(i, j int) bool {
		return start[i] < start[j]
	})

	// Sort the intervals by end time
	sort.Slice(end, func(i, j int) bool {
		return end[i] < end[j]
	})

	// The two pointers in the algorithm: e_ptr and s_ptr.
	startPointer, endPointer := 0, 0

	// Variables to keep track of maximum number of rooms used.
	usedRooms := 0

	// Iterate over intervals.
	for startPointer < len(intervals) {
		// If there is a meeting that has ended by the time the meeting at `start_pointer` starts
		if start[startPointer] >= end[endPointer] {
			usedRooms--
			endPointer++
		}

		// We do this irrespective of whether a room frees up or not.
		// If a room got free, then this used_rooms += 1 wouldn't have any effect. used_rooms would
		// remain the same in that case. If no room was free, then this would increase used_rooms
		usedRooms++
		startPointer++
	}
	return usedRooms
}

// --- 他人

// https://leetcode-cn.com/problems/meeting-rooms-ii/solution/golang-shang-xia-che-fa-by-bloodborne/
/*
方法一: 上下车法

时间复杂度：
空间复杂度：
*/
func minMeetingRoomsP1(intervals [][]int) int {
	var nums []int
	for _, v := range intervals {
		nums = append(nums, v[0]*10+2)
		nums = append(nums, v[1]*10+1)
	}
	sort.Ints(nums)
	maxRoom := 0
	curNeedRoom := 0
	for _, v := range nums {
		if v%10 == 1 {
			curNeedRoom--
		} else {
			curNeedRoom++
		}
		if curNeedRoom > maxRoom {
			maxRoom = curNeedRoom
		}
	}
	return maxRoom
}
