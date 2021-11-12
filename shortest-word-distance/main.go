/*
Package shortest_word_distance
https://leetcode-cn.com/problems/shortest-word-distance/

243. 最短单词距离

给定一个单词列表和两个单词 word1 和 word2，返回列表中这两个单词之间的最短距离。

注意:
	你可以假设 word1 不等于 word2, 并且 word1 和 word2 都在列表里。

*/
package shortest_word_distance

// --- 自己

/*
方法一:
	找到word1 word2 下标，分别构成集合 m,n, 问题转变为求两个有序集合中绝对值最小值

时间复杂度：O(N^2) ??
空间复杂度：O(N)
*/

/*
方法二:
	遍历单词列表，取首先找到的 word1 或 word2 记为 wordF, 继续查找，如果未匹配到 另一个单词,记为wordP，
	但找到 wordF, 则 指向新 wordF 位置；如果匹配到 则计算距离一次, 最后取距离最小值

时间复杂度：O(n)
空间复杂度：O(1)
*/
func shortestDistance(words []string, word1 string, word2 string) int {
	var wordF, wordP string
	var minDis, curDis int
	var fp int
	for i := 0; i < len(words); i++ {
		if wordF == "" {
			if words[i] == word1 {
				wordF = word1
				wordP = word2
				fp = i
			} else if words[i] == word2 {
				wordF = word2
				wordP = word1
				fp = i
			}
			continue
		}

		if words[i] == wordF {
			fp = i
			continue
		}

		if words[i] == wordP {
			curDis = i - fp
			if minDis == 0 {
				minDis = curDis
			} else {
				minDis = min(minDis, curDis)
			}

			tmp := wordF
			wordF = wordP
			wordP = tmp
			fp = i
		}
	}

	return minDis
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
*/
