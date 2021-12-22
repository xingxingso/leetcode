/*
Package possible_bipartition
https://leetcode-cn.com/problems/possible-bipartition/

886. 可能的二分法

给定一组 N 人（编号为 1, 2, ..., N）， 我们想把每个人分进任意大小的两组。
每个人都可能不喜欢其他人，那么他们不应该属于同一组。
形式上，如果 dislikes[i] = [a, b]，表示不允许将编号为 a 和 b 的人归入同一组。
当可以用这种方法将所有人分进两组时，返回 true；否则返回 false。

提示：
	1 <= N <= 2000
	0 <= dislikes.length <= 10000
	dislikes[i].length == 2
	1 <= dislikes[i][j] <= N
	dislikes[i][0] < dislikes[i][1]
	对于 dislikes[i] == dislikes[j] 不存在 i != j

*/
package possible_bipartition

// --- 官方

/*
方法一: 深度优先搜索

时间复杂度：O(N+E)，其中 E 是 dislikes 的长度。
空间复杂度：O(N+E)。
*/
func possibleBipartition(N int, dislikes [][]int) bool {
	color := make(map[int]int)
	graph := make([][]int, N+1)

	var dfs func(node, c int) bool
	dfs = func(node, c int) bool {
		if s, ok := color[node]; ok {
			return s == c
		}
		color[node] = c
		for _, nei := range graph[node] {
			if !dfs(nei, c^1) {
				return false
			}
		}
		return true
	}
	for i := 1; i <= N; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range dislikes {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	// fmt.Println(graph)
	for node := 1; node <= N; node++ {
		if _, ok := color[node]; !ok && !dfs(node, 0) {
			// fmt.Println(color)
			return false
		}
	}
	return true
}
