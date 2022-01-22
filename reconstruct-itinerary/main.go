/*
Package reconstruct_itinerary
https://leetcode-cn.com/problems/reconstruct-itinerary/

332. 重新安排行程

给你一份航线列表 tickets ，其中 tickets[i] = [from_i, to_i] 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。

所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。

例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。

提示：
	1 <= tickets.length <= 300
	tickets[i].length == 2
	from_i.length == 3
	to_i.length == 3
	from_i 和 to_i 由大写英文字母组成
	from_i != to_i

*/
package reconstruct_itinerary

import (
	"fmt"
	"sort"
)

// --- 自己

/*
方法一: 回溯

时间复杂度：
空间复杂度：
*/
func findItinerary(tickets [][]string) []string {
	h := make(map[string][]string)
	memo := make(map[string]int) // 记录机票
	maxLen := len(tickets) + 1
	for _, ticket := range tickets {
		//if _, ok := h[ticket[0]]; !ok {
		//	h[ticket[0]] = []string{ticket[1]}
		//} else {
		h[ticket[0]] = append(h[ticket[0]], ticket[1])
		//}
		memo[ticket[0]+ticket[1]]++
	}

	for _, v := range h {
		sort.Slice(v, func(i, j int) bool {
			return v[i][0] < v[j][0] ||
				(v[i][0] == v[j][0] && v[i][1] < v[j][1]) ||
				(v[i][0] == v[j][0] && v[i][1] == v[j][1] && v[i][2] < v[j][2])
		})
	}

	temp := []string{"JFK"}
	ans := make([]string, 0, maxLen)
	var dfs func(from string)
	dfs = func(from string) {
		//fmt.Println(temp)
		if len(ans) > 0 {
			return
		}
		if len(temp) == maxLen {
			ans = append(ans, temp...)
			return
		}
		toArr, ok := h[from]
		if !ok {
			return
		}
		var to string
		for i := 0; i < len(toArr); i++ {
			if memo[from+toArr[i]] <= 0 {
				continue
			}
			to = toArr[i]
			memo[from+to]--
			temp = append(temp, to)
			dfs(to)
			temp = temp[:len(temp)-1]
			memo[from+to]++
		}
	}

	from := "JFK"
	//fmt.Println(h)
	dfs(from)
	return ans
}

// --- 他人

/*
方法一:
	leetcode101 11.8

时间复杂度：
空间复杂度：
*/
func findItineraryP1(tickets [][]string) []string {
	hash := make(map[string][]string)
	for _, ticket := range tickets {
		hash[ticket[0]] = append(hash[ticket[0]], ticket[1])
	}

	for _, v := range hash {
		sort.Slice(v, func(i, j int) bool {
			return v[i][0] < v[j][0] ||
				(v[i][0] == v[j][0] && v[i][1] < v[j][1]) ||
				(v[i][0] == v[j][0] && v[i][1] == v[j][1] && v[i][2] < v[j][2])
		})
	}
	//fmt.Println(hash)
	stack := make([]string, 0)
	stack = append(stack, "JFK")
	ans := make([]string, 0)
	for len(stack) > 0 {
		next := stack[len(stack)-1]
		if len(hash[next]) == 0 { // 先找到终点
			ans = append(ans, next)
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, hash[next][0])
			hash[next] = hash[next][1:]
		}
	}

	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	fmt.Println(ans)
	return ans
}
