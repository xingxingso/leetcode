/*
Package number_of_valid_words_for_each_puzzle
https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle/

1178. 猜字谜

外国友人仿照中国字谜设计了一个英文版猜字谜小游戏，请你来猜猜看吧。
字谜的迷面 puzzle 按字符串形式给出，如果一个单词 word 符合下面两个条件，那么它就可以算作谜底：
	- 单词 word 中包含谜面 puzzle 的第一个字母。
	- 单词 word 中的每一个字母都可以在谜面 puzzle 中找到。
	例如，如果字谜的谜面是 "abcdefg"，那么可以作为谜底的单词有 "faced", "cabbage", 和 "baggage"；而 "beefed"（不含字母 "a"）
	以及 "based"（其中的 "s" 没有出现在谜面中）。
返回一个答案数组 answer，数组中的每个元素 answer[i] 是在给出的单词列表 words 中可以作为字谜迷面 puzzles[i] 所对应的谜底的单词数目。

提示：
	1 <= words.length <= 10^5
	4 <= words[i].length <= 50
	1 <= puzzles.length <= 10^4
	puzzles[i].length == 7
	words[i][j], puzzles[i][j] 都是小写英文字母。
	每个 puzzles[i] 所包含的字符都不重复。

*/
package number_of_valid_words_for_each_puzzle

import (
	"math/bits"
	"sort"
)

// --- 自己

/*
方法一: 暴力解法
	超出时限

时间复杂度：
空间复杂度：
*/
func findNumOfValidWords(words []string, puzzles []string) []int {
	ans := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		puzzleMap := map[byte]bool{}
		for j := 0; j < len(puzzle); j++ {
			puzzleMap[puzzle[j]] = true
		}
		for _, word := range words {
			if isValid(word, puzzleMap, puzzle[0]) {
				ans[i]++
			}
		}
	}
	return ans
}

func isValid(word string, puzzle map[byte]bool, first byte) bool {
	flag := false
	for i := 0; i < len(word); i++ {
		if !flag && word[i] == first {
			flag = true
		}

		if !puzzle[word[i]] {
			return false
		}
	}

	return flag
}

// --- 官方

/*
方法一: 二进制状态压缩

时间复杂度：O(m∣w∣+n2^∣p∣ )，其中 m 和 n 分别是数组 words 和 puzzles 的长度，∣w∣ 是 word 的最大长度 5050，
	∣p∣ 是 puzzle 的最大长度 7。时间复杂度分为三部分：
		- 计算所有 word 对应的二进制表示的时间复杂度为 O(m∣w∣)；
		- 计算所有 puzzle 对应的二进制表示的时间复杂度为 O(n∣p∣)；
		- 枚举 puzzle 的子集的时间复杂度为 O(n2^{|p|-1})，这里为使用第二种方法的时间复杂度，
		如果使用第一种方法，那么时间复杂度略高，为 O(n(|p|-1)2^{|p|-1})。
	由于 ∣p∣−1 与 ∣p∣ 同阶，因此写成 O(∣p∣) 更加简洁。并且由于第三部分的时间复杂度在渐进意义下严格大于第二部分，
	因此总时间复杂度即为第一部分与第三部分之和 O(m|w| + n2^{|p|})。
空间复杂度：O(m)，即为哈希映射需要使用的空间，其中最多只包含 m 个键值对。
*/
func findNumOfValidWordsO1(words []string, puzzles []string) []int {
	const puzzleLength = 7
	cnt := map[int]int{}
	for _, s := range words {
		mask := 0
		for _, ch := range s {
			mask |= 1 << (ch - 'a')
		}
		// fmt.Printf("w: %.27b\n", mask)
		if bits.OnesCount(uint(mask)) <= puzzleLength { // 大于 7种字符 不可能是谜底
			cnt[mask]++ // 不同排列的，键值相同，累加次数，行为相同(肯定同为谜底或者非谜底)
		}
	}

	ans := make([]int, len(puzzles))
	for i, s := range puzzles {
		first := 1 << (s[0] - 'a')

		// 枚举子集方法一
		//for choose := 0; choose < 1<<(puzzleLength-1); choose++ {
		//    mask := 0
		//    for j := 0; j < puzzleLength-1; j++ {
		//        if choose>>j&1 == 1 {
		//            mask |= 1 << (s[j+1] - 'a')
		//        }
		//    }
		//    ans[i] += cnt[mask|first]
		//}

		// 枚举子集方法二
		mask := 0
		for _, ch := range s[1:] {
			mask |= 1 << (ch - 'a')
		}
		subset := mask
		// fmt.Printf("p: %.27b\n", mask)
		for {
			ans[i] += cnt[subset|first]  // 必须包含首字符
			subset = (subset - 1) & mask // 枚举核心方法
			// fmt.Printf("subset: %.27b\n", subset)
			if subset == mask { // 当 subset 全是0后退出， 这时 (subset - 1) & mask = mask
				break
			}
		}
	}
	return ans
}

/*
方法二: 字典树

时间复杂度：
空间复杂度：
*/
func findNumOfValidWordsO2(words []string, puzzles []string) []int {
	root := &trieNode{}
	for _, word := range words {
		// 将 word 中的字母按照字典序排序并去重
		w := []byte(word)
		sort.Slice(w, func(i, j int) bool { return w[i] < w[j] })
		i := 0
		for _, ch := range w[1:] {
			if w[i] != ch {
				i++
				w[i] = ch
			}
		}
		w = w[:i+1]

		// 加入字典树中
		cur := root
		for _, ch := range w {
			ch -= 'a'
			if cur.son[ch] == nil {
				cur.son[ch] = &trieNode{}
			}
			cur = cur.son[ch]
		}
		cur.cnt++
	}

	ans := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		pz := []byte(puzzle)
		first := pz[0]
		sort.Slice(pz, func(i, j int) bool { return pz[i] < pz[j] })

		// 在回溯的过程中枚举 pz 的所有子集并统计答案
		var find func(int, *trieNode) int
		find = func(pos int, cur *trieNode) int {
			// 搜索到空节点，不合法，返回 0
			if cur == nil {
				return 0
			}

			// 整个 pz 搜索完毕，返回谜底的数量
			if pos == len(pz) {
				return cur.cnt
			}

			// 选择第 pos 个字母
			res := find(pos+1, cur.son[pz[pos]-'a'])

			// 当 pz[pos] 不为首字母时，可以不选择第 pos 个字母
			if pz[pos] != first {
				res += find(pos+1, cur)
			}

			return res
		}

		ans[i] = find(0, root)
	}

	return ans
}

type trieNode struct {
	son [26]*trieNode
	cnt int
}
