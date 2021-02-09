/*
https://leetcode-cn.com/problems/construct-binary-tree-from-string/

536. 从字符串生成二叉树

你需要从一个包括括号和整数的字符串构建一棵二叉树。
输入的字符串代表一棵二叉树。它包括整数和随后的 0 ，1 或 2 对括号。整数代表根的值，一对括号内表示同样结构的子树。
若存在左子结点，则从左子结点开始构建。

提示：
	输入字符串中只包含 '(', ')', '-' 和 '0' ~ '9'
	空树由 "" 而非"()"表示。

*/
package construct_binary_tree_from_string

import (
	"strconv"
)

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 栈
	() 是一个节点

时间复杂度：
空间复杂度：
*/
func str2tree(s string) *TreeNode {
	type cTreeNode struct {
		t     *TreeNode
		count int
	}

	stackNodes := make([]*cTreeNode, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}

		if isNumber(s[i]) {
			num, idx := readNumber(s, i)
			i = idx

			node := &TreeNode{
				Val: num,
			}
			if len(stackNodes) <= 0 {
				stackNodes = append(stackNodes, &cTreeNode{
					t: node,
				})
				continue
			}

			topNode := stackNodes[len(stackNodes)-1]
			if topNode.count == 0 {
				topNode.t.Left = node
				topNode.count++
			} else {
				topNode.t.Right = node
			}
			stackNodes = append(stackNodes, &cTreeNode{
				t: node,
			})
		} else if s[i] == ')' {
			stackNodes = stackNodes[:len(stackNodes)-1]
		}
	}

	if len(stackNodes) > 0 {
		return stackNodes[0].t
	}

	return nil
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9' || b == '-'
}

/**
i < len(s)
*/
func readNumber(s string, i int) (int, int) {
	start := i

	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' || s[i] == '-' {
			i++
		} else {
			break
		}
	}

	num, _ := strconv.Atoi(s[start:i])

	return num, i - 1
}
