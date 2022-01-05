/*
Package random_pick_with_blacklist
https://leetcode-cn.com/problems/random-pick-with-blacklist/

710. 黑名单中的随机数

给定一个包含 [0，n) 中不重复整数的黑名单 blacklist ，写一个函数从 [0, n) 中返回一个不在 blacklist 中的随机整数。
对它进行优化使其尽量少调用系统方法 Math.random() 。

提示:
	1 <= n <= 1000000000
	0 <= blacklist.length < min(100000, N)
	[0, n) 不包含 n ，详细参见 interval notation 。

输入语法说明：
	输入是两个列表：调用成员函数名和调用的参数。Solution的构造函数有两个参数，n 和黑名单 blacklist。
	pick 没有参数，输入参数是一个列表，即使参数为空，也会输入一个 [] 空列表。

*/
package random_pick_with_blacklist

import "math/rand"

// type Solution struct{}
// func Constructor(n int, blacklist []int) Solution {}
// func (this *Solution) Pick() int                  {}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */

// --- 他人
// labuladong
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247487414&idx=1&sn=2be87c0c9279da447f8ac8b8406230fe&chksm=9bd7f1beaca078a865357f58ba2ff12b46490b0a773c0221e0a846c67950fa9c661664ad500e&scene=21#wechat_redirect

/*
方法一:

时间复杂度：
空间复杂度：
*/

type Solution struct {
	mapping map[int]int
	size    int
}

func Constructor(n int, blacklist []int) Solution {
	// 实际可用数量
	size := n - len(blacklist)
	mapping := make(map[int]int, len(blacklist))
	for _, v := range blacklist {
		mapping[v] = 0
	}
	last := n - 1
	for _, v := range blacklist {
		if v >= size {
			continue
		}
		for {
			if _, ok := mapping[last]; ok {
				last--
			} else {
				break
			}
		}
		mapping[v] = last
		last--
	}
	return Solution{
		mapping: mapping,
		size:    size,
	}
}

func (this *Solution) Pick() int {
	// r := rand.New(rand.NewSource(time.Now().UnixNano())) // 这种方式会超时
	// index := r.Intn(this.size)
	index := rand.Intn(this.size)
	if n, ok := this.mapping[index]; ok {
		return n
	}
	return index
}
