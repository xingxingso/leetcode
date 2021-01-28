/*
https://leetcode-cn.com/problems/meeting-rooms/

252. 会议室

给定一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [start_i, end_i] ，
请你判断一个人是否能够参加这里面的全部会议。

提示：
	0 <= intervals.length <= 104
	intervals[i].length == 2
	0 <= start_i < end_i <= 106

*/
package meeting_rooms

// --- 自己

/*
方法一: 暴力法
	各个集合不能存在交集，依次与后面的比较

时间复杂度：O(n^2)
空间复杂度：O(1)
*/
func canAttendMeetings(intervals [][]int) bool {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][1] <= intervals[i][0] ||
				intervals[j][0] >= intervals[i][1] {
				continue
			}

			return false
		}
	}

	return true
}

/*
方法二: 排序

时间复杂度 : O(n log n) 。时间复杂度由排序决定。一旦排序完成，只需要 O(n) 的时间来判断交叠。
空间复杂度 : O(1)。没有使用额外空间。
*/
