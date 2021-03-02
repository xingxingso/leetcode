/*
https://leetcode-cn.com/problems/house-robber-iii/

337. 打家劫舍 III

在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。
除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

*/
package house_robber_iii

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 别人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484800&idx=1&sn=1016975b9e8df0b8f6df996a5fded0af&chksm=9bd7fb88aca0729eb2d450cca8111abd8f861236b04125ce556171cb520e298ddec4d90823b3&scene=21#wechat_redirect

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func robP1(root *TreeNode) int {
	memo := map[*TreeNode]int{}
	var dp func(root *TreeNode) int
	dp = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if n, ok := memo[root]; ok {
			return n
		}
		get, noGet := root.Val, 0
		if root.Left != nil {
			get += dp(root.Left.Left) + dp(root.Left.Right)
			noGet += dp(root.Left)
		}
		if root.Right != nil {
			get += dp(root.Right.Left) + dp(root.Right.Right)
			noGet += dp(root.Right)
		}
		profit := max(noGet, get)
		memo[root] = profit
		return profit
	}

	return dp(root)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func robP2(root *TreeNode) int {
	var dp func(root *TreeNode) (notGet, get int)
	dp = func(root *TreeNode) (notGet, get int) {
		if root == nil {
			return 0, 0
		}
		leftNotGet, leftGet := dp(root.Left)
		rightNotGet, rightGet := dp(root.Right)
		// 抢， 下家就不抢了
		get = root.Val + leftNotGet + rightNotGet
		// 不抢，下家可抢可不抢 取决于收益大小
		notGet = max(leftNotGet, leftGet) + max(rightNotGet, rightGet)
		return
	}
	return max(dp(root))
}
