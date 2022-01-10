/*
Package implement_rand10_using_rand7
https://leetcode-cn.com/problems/implement-rand10-using-rand7/

470. 用 Rand7() 实现 Rand10()

已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。

不要使用系统的 Math.random() 方法。

提示:
	rand7 已定义。
	传入参数: n 表示 rand10 的调用次数。

进阶:
	rand7()调用次数的 期望值 是多少 ?
	你能否尽量少调用 rand7() ?

*/
package implement_rand10_using_rand7

import (
	"math/rand"
	"time"
)

func rand7() int {
	//return rand.Intn(7) + 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(7) + 1
}

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func rand10() int {
	a := rand7()
	for a == 6 || a == 7 {
		a = rand7()
	}
	a *= 2

	b := rand7()
	for b == 7 {
		b = rand7()
	}

	if b > 3 {
		a--
	}

	return a
}
