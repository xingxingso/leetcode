/*
Package letter_combinations_of_a_phone_number
https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

17. 电话号码的字母组合

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

提示：
	0 <= digits.length <= 4
	digits[i] 是范围 ['2', '9'] 的一个数字。

*/
package letter_combinations_of_a_phone_number

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func letterCombinations(digits string) []string {
	h := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	ans := make([]string, 0)
	temp := ""
	var backTrack func(digits string)
	backTrack = func(digits string) {
		if len(digits) == 0 {
			ans = append(ans, temp)
			return
		}
		for _, v := range h[digits[0]] {
			temp += string(v)
			backTrack(digits[1:])
			temp = temp[:len(temp)-1]
		}
	}
	if len(digits) > 0 {
		backTrack(digits)
	}
	//fmt.Printf("%T, %[1]v, %d", ans, len(ans))
	return ans
}
