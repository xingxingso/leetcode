/*
https://leetcode-cn.com/problems/longest-substring-with-at-most-two-distinct-characters/

159. 至多包含两个不同字符的最长子串

给定一个字符串 s ，找出 至多 包含两个不同字符的最长子串 t ，并返回该子串的长度。
*/
package longest_substring_with_at_most_two_distinct_characters

// --- 自己

/*
方法一: 滑动窗口
*/
func lengthOfLongestSubstringTwoDistinct(s string) int {
	// 维护一个hash 存储了两个字符的最后出现位置
	// 指针遍历字符串，当出现新的字符的时候， 指针指向两个字符中的较小值的下一个
	var maxLen, currentLen int
	var mark = make(map[byte]int)
	var twoChar []byte

	for j := 0; j < len(s); j++ {
		if _, ok := mark[s[j]]; !ok {
			mark[s[j]] = j
			twoChar = append(twoChar, s[j])
		} else {
			mark[s[j]] = j
		}
		// fmt.Println(mark)
		if len(twoChar) > 2 {
			if mark[twoChar[0]] < mark[twoChar[1]] {
				currentLen = j - mark[twoChar[0]]
				delete(mark, twoChar[0])
				twoChar = twoChar[1:]
			} else {
				currentLen = j - mark[twoChar[1]]
				delete(mark, twoChar[1])
				twoChar = []byte{twoChar[0], twoChar[2]}
			}
		} else {
			currentLen++
		}
		// fmt.Println("nwe", mark, twoChar)
		maxLen = max(maxLen, currentLen)
	}

	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// --- 官方

/*
方法一：滑动窗口

时间复杂度：O(N) 其中 N 是输入串的字符数目。
空间复杂度：O(1)，这是因为额外的空间只有 hashmap ，且它包含不超过 3 个元素。
*/
func lengthOfLongestSubstringTwoDistinctO1(s string) int {
	n := len(s)
	if n < 3 {
		return n
	}

	// sliding window left and right pointers
	left, right := 0, 0
	// hashmap character -> its rightmost position
	// in the sliding window
	hashmap := make(map[byte]int)
	maxLen := 2
	for right < n {
		// slide_window contains less than 3 characters
		if len(hashmap) < 3 {
			hashmap[s[right]] = right
			right += 1
		}

		// slide_window contains 3 characters
		if len(hashmap) == 3 {
			// delete the leftmost character
			delIdx := min(hashmap)
			delete(hashmap, s[delIdx])
			// move left pointer of the slide_window
			left = delIdx + 1
		}
		maxLen = max(maxLen, right-left)
	}
	return maxLen
}

func min(hashmap map[byte]int) int {
	var minValue int
	for _, v := range hashmap {
		if minValue == 0 {
			minValue = v
		} else {
			if minValue > v {
				minValue = v
			}
		}
	}
	return minValue
}
