/*
Package gas_station
https://leetcode-cn.com/problems/gas-station/

134. 加油站

在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 gas 和 cost ，如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。

提示:
	gas.length == n
	cost.length == n
	1 <= n <= 10^5
	0 <= gas[i], cost[i] <= 10^4

*/
package gas_station

// --- 自己

/*
方法一: 暴力模拟
	// 超出时限

时间复杂度：
空间复杂度：
*/
func canCompleteCircuitS1(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		flag := true
		sum := gas[i]
		j := i
		for { // 这里应该可以优化i下一步的位置
			if cost[j] > sum {
				flag = false
				break
			}
			sum -= cost[j]
			j = (j + 1) % len(gas)
			sum += gas[j]
			if j == i {
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

/*
方法一: 一次遍历

时间复杂度：
空间复杂度：
*/
func canCompleteCircuitS2(gas []int, cost []int) int {
	dep, sum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		dif := gas[i] - cost[i]
		sum += dif
		dep += dif
		if sum < 0 { // 说明不能从 i 出发
			start = i + 1
			sum = 0
		}
	}

	if dep >= 0 { // 最终剩余的油大于等于0必有出发点可以完成
		return start
	}

	return -1
}
