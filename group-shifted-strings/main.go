/*
https://leetcode-cn.com/problems/group-shifted-strings/

249. 移位字符串分组

给定一个字符串，对该字符串可以进行 “移位” 的操作，也就是将字符串中每个字母都变为其在字母表中后续的字母，比如："abc" -> "bcd"。这样，
我们可以持续进行 “移位” 操作，从而生成如下移位序列：
	"abc" -> "bcd" -> ... -> "xyz"
给定一个包含仅小写字母字符串的列表，将该列表中所有满足 “移位” 操作规律的组合进行分组并返回。

解释：
	可以认为字母表首尾相接，所以 'z' 的后续为 'a'，所以 ["az","ba"] 也满足 “移位” 操作规律。

*/
package group_shifted_strings

// --- 自己

/*
方法一:
	自己想到了使用hash key 但是对于 az ba 没处理好， 参考了下面的方法
	https://leetcode-cn.com/problems/group-shifted-strings/solution/zhong-gui-zhong-ju-qiao-miao-ji-suan-has-ys3z/
	https://leetcode-cn.com/problems/group-shifted-strings/solution/golangha-xi-biao-yi-ci-bian-li-su-du-zui-ca39/

时间复杂度：
空间复杂度：
*/
func groupStrings(strings []string) [][]string {
	s := make(map[string][]string)
	for i := 0; i < len(strings); i++ {
		mark := []byte(strings[i])
		// fmt.Println("ok", mark)
		for j := 1; j < len(mark); j++ {
			mark[j-1] = (26 + mark[j] - mark[j-1]) % 26
		}
		// 最后一位没有处理 统一置为0
		mark[len(mark)-1] = '0'
		// fmt.Println(mark)
		s[string(mark)] = append(s[string(mark)], strings[i])
	}

	// fmt.Println(s)
	var res [][]string
	for _, v := range s {
		res = append(res, v)
	}

	return res
}
