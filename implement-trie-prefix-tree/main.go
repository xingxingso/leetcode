/*
Package implement_trie_prefix_tree
https://leetcode-cn.com/problems/implement-trie-prefix-tree/

208. 实现 Trie (前缀树)

Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：
	Trie() 初始化前缀树对象。
	void insert(String word) 向前缀树中插入字符串 word 。
	boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
	boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

提示：
	1 <= word.length, prefix.length <= 2000
	word 和 prefix 仅由小写英文字母组成
	insert、search 和 startsWith 调用次数 总计 不超过 3 * 10^4 次

*/
package implement_trie_prefix_tree

//type Trie struct {}
//func Constructor() Trie {}
//func (this *Trie) Insert(word string)  {}
//func (this *Trie) Search(word string) bool {}
//func (this *Trie) StartsWith(prefix string) bool {}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/

type Trie struct {
	root *node
}

type node struct {
	val      byte
	children []*node
	tail     bool
}

func Constructor() Trie {
	return Trie{
		root: &node{},
	}
}

func (t *Trie) Insert(word string) {
	for i, n := 0, t.root; i < len(word); i++ {
		flag := false
		for _, child := range n.children {
			if child.val == word[i] {
				if i+1 == len(word) {
					child.tail = true
				}
				flag = true
				n = child
				break
			}
		}
		if flag {
			continue
		}
		l := &node{val: word[i]}
		if i+1 == len(word) {
			l.tail = true
		}

		n.children = append(n.children, l)
		n = l
	}
}

func (t *Trie) Search(word string) bool {
	for i, n := 0, t.root; i < len(word); i++ {
		flag := false
		for _, child := range n.children {
			if child.val == word[i] {
				if i+1 < len(word) {
					flag = true
				} else {
					flag = child.tail
				}
				n = child
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

func (t *Trie) StartsWith(prefix string) bool {
	for i, n := 0, t.root; i < len(prefix); i++ {
		flag := false
		for _, child := range n.children {
			if child.val == prefix[i] {
				n = child
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
