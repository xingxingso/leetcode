/*
Package corporate_flight_bookings
https://leetcode-cn.com/problems/corporate-flight-bookings/

1109. 航班预订统计

这里有 n 个航班，它们分别从 1 到 n 进行编号。
有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [first_i, last_i, seats_i]
意味着在从 first_i 到 last_i （包含 first_i 和 last_i ）的 每个航班 上预订了 seats_i 个座位。

请你返回一个长度为 n 的数组 answer，其中 answer[i] 是航班 i 上预订的座位总数。

提示：
	1 <= n <= 2 * 10^4
	1 <= bookings.length <= 2 * 10^4
	bookings[i].length == 3
	1 <= first_i <= last_i <= n
	1 <= seats_i <= 10^4

*/
package corporate_flight_bookings

// --- 自己

/*
方法一: 暴力解法

时间复杂度：
空间复杂度：
*/
func corpFlightBookingsS1(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, booking := range bookings {
		for i := booking[0]; i <= booking[1]; i++ {
			ans[i-1] += booking[2]
		}
	}
	return ans
}

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247487011&idx=1&sn=5e2b00c1c736fd7afbf3ed35edc4aeec&chksm=9bd7f02baca0793d569a9633cc14117e708ccc9eb41b7f0add430ea78f22e4f2443f421c6841&scene=21#wechat_redirect

/*
方法一: 差分数组

时间复杂度：
空间复杂度：
*/
func corpFlightBookingsP1(bookings [][]int, n int) []int {
	ans := make([]int, n, n)

	for _, booking := range bookings {
		ans[booking[0]-1] += booking[2]
		if booking[1] < n {
			ans[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}

	return ans
}
