/*
Package find_smallest_letter_greater_than_target
https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/

744. 寻找比目标字母大的最小字母

给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母 target，请你寻找在这一有序列表里比目标字母大的最小字母。

在比较时，字母是依序循环出现的。举个例子：
	如果目标字母 target = 'z' 并且字符列表为 letters = ['a', 'b']，则答案返回 'a'

提示：
	2 <= letters.length <= 10^4
	letters[i] 是一个小写字母
	letters 按非递减顺序排序
	letters 最少包含两个不同的字母
	target 是一个小写字母

*/
package find_smallest_letter_greater_than_target

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	for l <= r {
		mid := l + (r-l)/2
		if letters[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l >= len(letters) {
		l = 0
	}
	return letters[l]
}
