/*
Package open_the_lock
https://leetcode-cn.com/problems/open-the-lock/

752. 打开转盘锁

你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。

锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。

列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。

字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。

提示：
	1 <= deadends.length <= 500
	deadends[i].length == 4
	target.length == 4
	target 不在 deadends 之中
	target 和 deadends[i] 仅由若干位数字组成

*/
package open_the_lock

// --- 自己

/*
方法一: 广度优先搜索 BFS

时间复杂度：
空间复杂度：
*/
func openLock(deadends []string, target string) int {
	step := 0
	mark := make(map[string]bool, len(deadends))
	for i := 0; i < len(deadends); i++ {
		mark[deadends[i]] = true
	}
	if mark["0000"] {
		return -1
	}
	mark["0000"] = true
	queue := []string{"0000"}
	for len(queue) > 0 {
		l := len(queue)
		for _, str := range queue[:l] {
			if str == target {
				return step
			}
			for _, v := range rotate(str) {
				if mark[v] {
					continue
				}
				queue = append(queue, v)
				mark[v] = true
			}
		}
		step++
		queue = queue[l:]
	}

	return -1
}

func rotate(cur string) []string {
	res := make([]string, 0)
	for i := 0; i < len(cur); i++ {
		temp := []byte(cur)
		if cur[i] == '9' {
			temp[i] = '0'
		} else {
			temp[i] = cur[i] + 1
		}

		temp2 := []byte(cur)
		if cur[i] == '0' {
			temp2[i] = '9'
		} else {
			temp2[i] = cur[i] - 1
		}

		res = append(res, string(temp), string(temp2))
	}
	return res
}
