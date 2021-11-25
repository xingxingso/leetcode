/*
Package word_ladder_ii
https://leetcode-cn.com/problems/word-ladder-ii/

126. 单词接龙 II

按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，
一个表示此过程的 转换序列 是形式上像 beginWord -> s1 -> s2 -> ... -> sk 这样的单词序列，并满足：
	- 每对相邻的单词之间仅有单个字母不同。
	- 转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList 中的单词。
	- sk == endWord

给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。
请你找出并返回所有从 beginWord 到 endWord 的 最短转换序列 ，如果不存在这样的转换序列，返回一个空列表。
每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk] 的形式返回。

提示：
	1 <= beginWord.length <= 7
	endWord.length == beginWord.length
	1 <= wordList.length <= 5000
	wordList[i].length == beginWord.length
	beginWord、endWord 和 wordList[i] 由小写英文字母组成
	beginWord != endWord
	wordList 中的所有单词 互不相同

*/
package word_ladder_ii

// ---官方

/*
方法一: 广度优先搜索

时间复杂度：
空间复杂度：
*/
func findLaddersO1(beginWord string, endWord string, wordList []string) [][]string {
	res := make([][]string, 0)
	// 因为需要快速判断扩展出的单词是否在 wordList 里，因此需要将 wordList 存入哈希表，这里命名为「字典」
	dict := make(map[string]bool)
	for _, v := range wordList {
		dict[v] = true
	}
	// 特殊用例判断
	if !dict[endWord] {
		return res
	}
	dict[beginWord] = false

	// 第 1 步：广度优先遍历建图
	// 记录扩展出的单词是在第几次扩展的时候得到的，key：单词，value：在广度优先遍历的第几层
	steps := make(map[string]int)
	steps[beginWord] = 0
	// 记录了单词是从哪些单词扩展而来，key：单词，value：单词列表，这些单词可以变换到 key ，它们是一对多关系
	from := make(map[string][]string)
	step := 1
	found := false
	wordLen := len(beginWord)
	queue := []string{beginWord}
	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			currWord := queue[0]
			currByte := []byte(currWord)
			queue = queue[1:]
			// 将每一位替换成 26 个小写英文字母
			for i := 0; i < wordLen; i++ {
				ch := currByte[i]
				for c := byte('a'); c <= 'z'; c++ {
					currByte[i] = c
					nextWord := string(currByte)
					if nextStep, ok := steps[nextWord]; ok && step == nextStep {
						from[nextWord] = append(from[nextWord], currWord)
					}
					if !dict[nextWord] {
						continue
					}
					//fmt.Printf("%s\n", nextWord)
					// 如果从一个单词扩展出来的单词以前遍历过，距离一定更远，为了避免搜索到已经遍历到，且距离更远的单词，需要将它从 dict 中删除
					dict[nextWord] = false
					// 这一层扩展出的单词进入队列
					queue = append(queue, nextWord)
					//fmt.Printf("%v\n", queue)

					// 记录 nextWord 从 currWord 而来
					from[nextWord] = append(from[nextWord], currWord)
					// 记录 nextWord 的 step
					steps[nextWord] = step
					//fmt.Printf("%v, %v\n", from, steps)
					if nextWord == endWord {
						found = true
					}
				}
				currByte[i] = ch
			}
		}
		step++
		if found {
			break
		}
	}

	//fmt.Println(from)
	// 第 2 步：深度优先遍历找到所有解，从 endWord 恢复到 beginWord ，所以每次尝试操作 path 列表的头部
	if found {
		path := []string{endWord}
		dfs(from, path, beginWord, endWord, &res)
	}

	return res
}

func dfs(from map[string][]string, path []string, beginWord, cur string, res *[][]string) {
	if cur == beginWord {
		*res = append(*res, path)
		return
	}
	for _, precursor := range from[cur] {
		path = append([]string{precursor}, path...)
		dfs(from, path, beginWord, precursor, res)
		path = path[1:] // 回溯
	}
}
