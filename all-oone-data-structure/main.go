/*
Package all_oone_data_structure
https://leetcode-cn.com/problems/all-oone-data-structure/

432. 全 O(1) 的数据结构

请你设计一个用于存储字符串计数的数据结构，并能够返回计数最小和最大的字符串。

实现 AllOne 类：
	AllOne() 初始化数据结构的对象。
	inc(String key) 字符串 key 的计数增加 1 。如果数据结构中尚不存在 key ，那么插入计数为 1 的 key 。
	dec(String key) 字符串 key 的计数减少 1 。如果 key 的计数在减少后为 0 ，那么需要将这个 key 从数据结构中删除。测试用例保证：在减少计数前，key 存在于数据结构中。
	getMaxKey() 返回任意一个计数最大的字符串。如果没有元素存在，返回一个空字符串 "" 。
	getMinKey() 返回任意一个计数最小的字符串。如果没有元素存在，返回一个空字符串 "" 。

提示：
	1 <= key.length <= 10
	key 由小写英文字母组成
	测试用例保证：在每次调用 dec 时，数据结构中总存在 key
	最多调用 inc、dec、getMaxKey 和 getMinKey 方法 5 * 10^4 次

*/
package all_oone_data_structure

import "fmt"

//type AllOne struct {}
//func Constructor() AllOne {}
//func (this *AllOne) Inc(key string) {}
//func (this *AllOne) Dec(key string) {}
//func (this *AllOne) GetMaxKey() string {}
//func (this *AllOne) GetMinKey() string {}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

// --- 自己

/*
方法一: 哈希表+十字链表

时间复杂度：
空间复杂度：
*/

type AllOne struct {
	h          map[string]*keyList
	head, tail *valList
}

type valList struct {
	val       int
	pre, next *valList
	head      *keyList
	tail      *keyList
}

type keyList struct {
	val       string
	v         *valList // 全指向数量节点
	pre, next *keyList
}

func Constructor() AllOne {
	head, tail := &valList{val: 0}, &valList{val: -1}
	head.next = tail
	tail.pre = head
	return AllOne{
		h:    make(map[string]*keyList),
		head: head,
		tail: tail,
	}
}

func (s *AllOne) Inc(key string) {
	keyNode, ok := s.h[key]
	if !ok {
		keyNode = &keyList{val: key}
		s.h[key] = keyNode
		valNode := s.getNextValNode(s.head)
		s.moveKeyNode(keyNode, valNode)
		//printAllOne(s)
		return
	}
	valNode := s.getNextValNode(keyNode.v)
	s.removeKeyNode(keyNode)
	s.moveKeyNode(keyNode, valNode)

	//printAllOne(s)
}

func (s *AllOne) Dec(key string) {
	keyNode, ok := s.h[key]
	if !ok {
		return
	}

	valNode := s.getPreValNode(keyNode.v)
	s.removeKeyNode(keyNode)
	if valNode != s.head {
		s.moveKeyNode(keyNode, valNode)
	} else {
		delete(s.h, key)
	}
}

func (s *AllOne) GetMaxKey() string {
	if s.tail.pre == s.head {
		return ""
	}

	return s.tail.pre.head.next.val
}

func (s *AllOne) GetMinKey() string {
	if s.head.next == s.tail {
		return ""
	}

	return s.head.next.head.next.val
}

func (s *AllOne) getNextValNode(node *valList) *valList {
	if node == s.tail {
		return nil
	}

	if node.next.val != node.val+1 {
		head, tail := &keyList{}, &keyList{}
		head.next = tail
		tail.pre = head
		newNode := &valList{
			val:  node.val + 1,
			head: head,
			tail: tail,
		}
		newNode.pre = node
		newNode.next = node.next
		node.next = newNode
		newNode.next.pre = newNode
	}

	return node.next
}

func (s *AllOne) getPreValNode(node *valList) *valList {
	if node == s.head {
		return nil
	}

	// 只有前面没有数量大于0的节点 才新建
	if node.pre.val != node.val-1 {
		head, tail := &keyList{}, &keyList{}
		head.next = tail
		tail.pre = head
		newNode := &valList{
			val:  node.val - 1,
			head: head,
			tail: tail,
		}
		newNode.next = node
		newNode.pre = node.pre
		node.pre = newNode
		newNode.pre.next = newNode
	}

	return node.pre
}

func (s *AllOne) removeKeyNode(node *keyList) {
	node.next.pre = node.pre
	node.pre.next = node.next
	node.next = nil
	node.pre = nil
	valNode := node.v

	// 清理 valNode
	if valNode.head.next == valNode.tail {
		valNode.pre.next = valNode.next
		valNode.next.pre = valNode.pre
		valNode.pre = nil
		valNode.next = nil
	}
}

func (s *AllOne) moveKeyNode(keyNode *keyList, valNode *valList) {
	keyNode.pre = valNode.tail.pre
	keyNode.next = valNode.tail
	keyNode.pre.next = keyNode
	keyNode.next.pre = keyNode
	keyNode.v = valNode
}

func printAllOne(s *AllOne) {
	fmt.Println("-------------------s.h")
	for k, v := range s.h {
		fmt.Printf("k:%v, v:%v;", k, v.val)
	}
	fmt.Println()
	fmt.Println("s.list")
	head := s.head

	for head != nil {
		fmt.Printf("s.l: %v\n", head.val)
		node := head.head
		for node != nil {
			n := -1000
			if node.v != nil {
				n = node.v.val
			}
			fmt.Printf("\tkey:%v, n:%v;", node.val, n)
			node = node.next
		}
		fmt.Println()
		head = head.next
	}
	fmt.Println("------------------------")
}
