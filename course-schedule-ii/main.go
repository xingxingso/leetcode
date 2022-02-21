/*
Package course_schedule_ii
https://leetcode-c.com/problems/course-schedule-ii/

210. 课程表 II

现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。
给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
	例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。

返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。
如果不可能完成所有课程，返回 一个空数组 。

提示：
	1 <= numCourses <= 2000
	0 <= prerequisites.length <= numCourses * (numCourses - 1)
	prerequisites[i].length == 2
	0 <= ai, bi < numCourses
	ai != bi
	所有[ai, bi] 互不相同

*/
package course_schedule_ii

// --- 他人

/*
方法一:拓扑排序
	leetcode101 15.3
	拓扑排序也可以被看成是广度优先搜索的一种情况:
	这里注意我们将所有的边反向，使得如果课程 i 指向课程 j，那么课程 i 需要在课程 j 前面先修完。
	我们先遍历一遍所有节点，把入度为 0 的节点(即没有前置课程要求)放在队列中。
	在每次从队列中获得节点时，我们将该节点放在目 前排序的末尾，并且把它指向的课程的入度各减 1;
	如果在这个过程中有课程的所有前置必修课 都已修完(即入度为 0)，我们把这个节点加入队列中。
	当队列的节点都被处理完时，说明所有的节点都已排好序，或因图中存在循环而无法上完所有课程。

时间复杂度：
空间复杂度：
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, 0)
	}
	res, indegree := make([]int, 0), make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		// prerequisite[1] 课程是 prerequisite[0] 的前置课程
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		indegree[prerequisite[0]]++ // prerequisite[0] 的前置课程数目，0表示可以学习了
	}
	queue := make([]int, 0)
	for i := 0; i < len(indegree); i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		u := queue[0]
		res = append(res, u)
		queue = queue[1:]
		for _, v := range graph[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	for _, v := range indegree {
		if v > 0 {
			return nil
		}
	}
	return res
}
