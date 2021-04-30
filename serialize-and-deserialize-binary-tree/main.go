/*
https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

297. 二叉树的序列化与反序列化

序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

提示:
	输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。
	你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

提示：
	树中结点数在范围 [0, 10^4] 内
	-1000 <= Node.val <= 1000

*/
package serialize_and_deserialize_binary_tree

import (
	"fmt"
	"strconv"
	"strings"
)

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485871&idx=1&sn=bcb24ea8927995b585629a8b9caeed01&chksm=9bd7f7a7aca07eb1b4c330382a4e0b916ef5a82ca48db28908ab16563e28a376b5ca6805bec2&scene=21#wechat_redirect

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

/*
方法一: 前序遍历

时间复杂度：
空间复杂度：
*/
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ans := ""
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			ans += "#,"
			return
		}
		ans += strconv.Itoa(root.Val) + ","
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	str := strings.Split(data[:len(data)-1], ",") // 多余的 ,
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if len(str) == 0 || str[0] == "#" {
			str = str[1:]
			return nil
		}
		val, _ := strconv.Atoi(str[0])
		str = str[1:]
		root := &TreeNode{Val: val}
		root.Left = dfs()
		root.Right = dfs()
		return root
	}

	return dfs()
}

/*
方法二: 后序遍历

时间复杂度：
空间复杂度：
*/
// Serializes a tree to a single string.
func (this *Codec) serializeP2(root *TreeNode) string {
	ans := ""
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			ans += "#,"
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		ans += strconv.Itoa(root.Val) + ","
	}
	dfs(root)
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializeP2(data string) *TreeNode {
	str := strings.Split(data[:len(data)-1], ",") // 多余的 ,
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if len(str) == 0 {
			return nil
		}
		last := str[len(str)-1]
		str = str[:len(str)-1]
		if last == "#" {
			return nil
		}
		val, _ := strconv.Atoi(last)
		root := &TreeNode{Val: val}
		root.Right = dfs()
		root.Left = dfs()
		return root
	}

	return dfs()
}

/*
方法三: 层级遍历
	中序遍历不行

时间复杂度：
空间复杂度：
*/
// Serializes a tree to a single string.
func (this *Codec) serializeP3(root *TreeNode) string {
	ans := ""
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			ans += "#,"
			continue
		}
		ans += strconv.Itoa(cur.Val) + ","
		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializeP3(data string) *TreeNode {
	str := strings.Split(data[:len(data)-1], ",") // 多余的 ,
	// fmt.Println(str)
	if len(str) == 0 {
		return nil
	}
	queue := make([]*TreeNode, 0)
	if str[0] == "#" {
		return nil
	}
	val, _ := strconv.Atoi(str[0])
	root := &TreeNode{Val: val}
	queue = append(queue, root)

	for i := 1; i < len(str); {
		parent := queue[0]
		queue = queue[1:]
		left := str[i]
		i++
		if left != "#" {
			val, _ := strconv.Atoi(left)
			parent.Left = &TreeNode{Val: val}
			queue = append(queue, parent.Left)
		}
		right := str[i]
		i++
		if right != "#" {
			val, _ := strconv.Atoi(right)
			parent.Right = &TreeNode{Val: val}
			queue = append(queue, parent.Right)
		}
	}
	return root
}

func inOrderPrint(root *TreeNode) {
	if root == nil {
		fmt.Printf("%s", "#,")
		return
	}
	inOrderPrint(root.Left)
	fmt.Printf("%d,", root.Val)
	inOrderPrint(root.Right)
	// fmt.Println()
}
