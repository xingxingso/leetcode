/*
Package linked_list_random_node
https://leetcode-cn.com/problems/linked-list-random-node/

382. 链表随机节点

给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。保证每个节点被选的概率一样。

进阶:
	如果链表十分大且长度未知，如何解决这个问题？你能否使用常数级空间复杂度实现？

*/
package linked_list_random_node

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//type Solution struct{}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
//func Constructor(head *ListNode) Solution {}

/** Returns a random node's value. */
//func (this *Solution) GetRandom() int {}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484974&idx=1&sn=795a33c338d4a5bd8d265bc7f9f63c03&scene=21#wechat_redirect

/*
方法一: 水塘抽样算法

时间复杂度：
空间复杂度：
*/

type SolutionS1 struct {
	head *ListNode
	r    *rand.Rand
}

func ConstructorS1(head *ListNode) SolutionS1 {
	return SolutionS1{
		head: head,
		r:    rand.New(rand.NewSource(time.Now().UnixNano())), // 保证随机性
	}
}

func (this *SolutionS1) GetRandom() int {
	ans, count := 0, 1
	for cur := this.head; cur != nil; cur = cur.Next {
		if this.r.Intn(count)+1 == count {
			//if rand.Intn(count)+1 == count { // 这样每次都是固定的随机
			ans = cur.Val
		}
		count++
	}
	return ans
}
