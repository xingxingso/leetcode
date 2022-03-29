/*
Package course_schedule
https://leetcode-cn.com/problems/course-schedule/

207. 课程表

你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。
先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

提示：
	1 <= numCourses <= 10^5
	0 <= prerequisites.length <= 5000
	prerequisites[i].length == 2
	0 <= ai, bi < numCourses
	prerequisites[i] 中的所有课程对 互不相同

*/
package course_schedule

// --- 自己

/*
方法一:拓扑排序

时间复杂度：
空间复杂度：
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	need := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		need[i] = make([]int, 0)
	}

	countMap := make(map[int]int)
	revers := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		need[prerequisites[i][0]] = append(need[i], prerequisites[i][1])
		revers[prerequisites[i][1]] = append(revers[prerequisites[i][1]], prerequisites[i][0])
		countMap[prerequisites[i][0]]++
	}

	queue := make([]int, 0)
	count := numCourses
	for i := 0; i < numCourses; i++ {
		if len(need[i]) == 0 {
			queue = append(queue, i)
		}
		for len(queue) > 0 {
			size := len(queue)
			count -= size
			for i := 0; i < size; i++ {
				node := queue[i]
				if nodes, ok := revers[node]; ok {
					for _, v := range nodes {
						countMap[v]--
						if countMap[v] == 0 {
							queue = append(queue, v)
						}
					}
				}
			}

			queue = queue[size:]
		}
	}

	return count == 0
}
