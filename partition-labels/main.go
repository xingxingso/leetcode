/*
Package partition_labels
https://leetcode-cn.com/problems/partition-labels/

763. 划分字母区间

字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。
返回一个表示每个字符串片段的长度的列表。

提示：
	S的长度在[1, 500]之间。
	S只包含小写字母 'a' 到 'z' 。

*/
package partition_labels

// --- 自己

/*
方法一: 贪心

时间复杂度：
空间复杂度：
*/
func partitionLabels(s string) []int {
	h := make(map[byte]int) // the max index hash
	for i := 0; i < len(s); i++ {
		h[s[i]] = i
	}

	res := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if h[s[i]] == i {
			res = append(res, 1)
			continue
		}
		max := h[s[i]]
		for j := i + 1; j <= max; j++ {
			if h[s[j]] > max {
				max = h[s[j]]
			}
		}
		res = append(res, max-i+1)
		i = max
	}

	return res
}

// --- 官方

/*
方法一: 贪心

时间复杂度：
空间复杂度：
*/
func partitionLabelsO1(s string) []int {
	partition := make([]int, 0)
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}

	start, end := 0, 0
	for i, c := range s {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return partition
}
