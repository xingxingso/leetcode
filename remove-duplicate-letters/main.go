/*
Package remove_duplicate_letters
https://leetcode-cn.com/problems/remove-duplicate-letters/

316. 去除重复字母

给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
注意：
	该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters 相同

提示：
	1 <= s.length <= 10^4
	s 由小写英文字母组成

*/
package remove_duplicate_letters

type StackSlice []byte

func NewStackSlice() *StackSlice {
	return &StackSlice{}
}

func (s *StackSlice) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StackSlice) Pop() byte {
	if s.IsEmpty() {
		return 0
	}

	val := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]

	return val
}

func (s *StackSlice) Push(val byte) {
	*s = append(*s, val)
}

func (s *StackSlice) Peek() byte {
	if s.IsEmpty() {
		return 0
	}

	return (*s)[len(*s)-1]
}

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247486946&idx=1&sn=94804eb15be33428582544a1cd90da4d&chksm=9bd7f3eaaca07afc6fdfa94d05fa3007d9ecc54914a238e6deabeafd5032a299155505b40f2d&scene=21#wechat_redirect

/*
方法一:

时间复杂度：
空间复杂度：
*/
func removeDuplicateLetters(s string) string {
	count := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	memo := make(map[byte]bool, 0)
	var buf []byte
	stack := NewStackSlice()

	for i := 0; i < len(s); i++ {
		count[s[i]]--
		if memo[s[i]] {
			continue
		}

		for !stack.IsEmpty() && stack.Peek() > s[i] {
			if count[stack.Peek()] == 0 {
				break
			}
			memo[stack.Pop()] = false
		}

		stack.Push(s[i])
		memo[s[i]] = true
	}

	for !stack.IsEmpty() {
		buf = append(buf, stack.Pop())
	}
	// 翻转
	for i := 0; i < len(buf)/2; i++ {
		buf[i], buf[len(buf)-1-i] = buf[len(buf)-1-i], buf[i]
	}
	return string(buf)
}

// --- 官方

/*
方法一: 贪心+单调栈

时间复杂度：O(N)，其中 N 为字符串长度。代码中虽然有双重循环，但是每个字符至多只会入栈、出栈各一次。
空间复杂度：O(∣Σ∣)，其中 Σ 为字符集合，本题中字符均为小写字母，所以 ∣Σ∣=26。由于栈中的字符不能重复，
	因此栈中最多只能有 ∣Σ∣ 个字符，另外需要维护两个数组，分别记录每个字符是否出现在栈中以及每个字符的剩余数量。
*/
func removeDuplicateLettersO1(s string) string {
	left := [26]int{}
	for i := 0; i < len(s); i++ {
		left[s[i]-'a']++
	}
	stack := []byte{}
	instack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !instack[ch-'a'] {
			for len(stack) > 0 && stack[len(stack)-1] > ch {
				last := stack[len(stack)-1] - 'a'
				if left[last] <= 0 {
					break
				}
				stack = stack[:len(stack)-1]
				instack[last] = false
			}

			instack[ch-'a'] = true
			stack = append(stack, ch)
		}
		left[ch-'a']--
	}
	return string(stack)
}
