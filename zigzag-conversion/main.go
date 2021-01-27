/*
https://leetcode-cn.com/problems/zigzag-conversion/

6. Z 字形变换

将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
请你实现这个将字符串进行指定行数变换的函数：
string convert(string s, int numRows);

提示：
1 <= s.length <= 1000
s 由英文字母（小写和大写）、',' 和 '.' 组成
1 <= numRows <= 1000
*/
package zigzag_conversion

import "strings"

// --- 自己

/*
方法一：二维数组
*/
func convert(s string, numRows int) string {
	res := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]byte, len(s))
	}

	zNum := 0
	if numRows >= 2 {
		zNum = numRows - 2
	} else {
		zNum = 0
	}

	zANum := numRows + zNum
	p := 0
	k := zNum
	for j := 0; j < len(s); j++ {
		quo := j / zANum
		rem := j % zANum
		if rem < numRows {
			res[rem][quo+p] = s[j]
		} else {
			p++
			res[k][quo+p] = s[j]
			k--
		}

		if k <= 0 {
			k = zNum
		}
	}

	var result []byte
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] != 0 {
				result = append(result, res[i][j])
			}
		}
	}

	return string(result)
}

// --- 官方

/*
方法一：按行排序

时间复杂度：O(n)，其中 n ==len(s)
空间复杂度：O(n)
*/
func convertO1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, min(numRows, len(s)))

	curRow := 0
	goingDown := false

	for _, c := range s {
		rows[curRow] += string(c)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow += 1
		} else {
			curRow += -1
		}
	}

	return strings.Join(rows, "")
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
方法二：按行访问

时间复杂度：O(n)，其中 n == len(s)。每个索引被访问一次。
空间复杂度：O(n)。对于 C++ 实现，如果返回字符串不被视为额外空间，则复杂度为 O(1)。
*/
func convertO2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var ret string
	n := len(s)
	cycleLen := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			ret += string(s[j+i])
			if i != 0 && (i != numRows-1) && (j+cycleLen-i < n) {
				ret += string(s[j+cycleLen-i])
			}
		}
	}
	return ret
}
