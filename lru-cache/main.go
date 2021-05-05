/*
https://leetcode-cn.com/problems/lru-cache/

146. LRU 缓存机制

运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：
	- LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
	- int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
	- void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
      当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

进阶：
	你是否可以在 O(1) 时间复杂度内完成这两种操作？

提示：

	- 1 <= capacity <= 3000
	- 0 <= key <= 3000
	- 0 <= value <= 10^4
	- 最多调用 3 * 10^4 次 get 和 put

*/
package lru_cache

import "fmt"

//type LRUCache struct{}
//func Constructor(capacity int) LRUCache       {}
//func (this *LRUCache) Get(key int) int        {}
//func (this *LRUCache) Put(key int, value int) {}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/

type LRUCacheS1 struct {
	cap, len   int
	cache      map[int]*DoubleList
	head, tail *DoubleList
}

type DoubleList struct {
	key, val  int
	pre, next *DoubleList
}

func ConstructorS1(capacity int) LRUCacheS1 {
	head := &DoubleList{}
	tail := &DoubleList{}
	head.next = tail
	tail.pre = head

	return LRUCacheS1{
		cap:   capacity,
		len:   0,
		cache: make(map[int]*DoubleList, 0),
		head:  head,
		tail:  tail,
	}
}

func (this *LRUCacheS1) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCacheS1) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.moveToHead(node)
		return
	}
	node := &DoubleList{
		key: key,
		val: value,
	}
	this.cache[key] = node
	this.len++
	this.addToHead(node)
	this.removeTail()
}

func (this *LRUCacheS1) moveToHead(node *DoubleList) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCacheS1) removeNode(node *DoubleList) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCacheS1) addToHead(node *DoubleList) {
	node.pre = this.head
	node.next = this.head.next
	//this.head.next.pre = node
	node.next.pre = node
	this.head.next = node
}

func (this *LRUCacheS1) removeTail() {
	if this.len <= this.cap {
		return
	}
	node := this.tail.pre
	// remove from list
	//this.tail.pre = node.pre
	//node.pre.next = this.tail
	this.removeNode(node)
	// remove from cache
	delete(this.cache, node.key)
	this.len--
}

func printList(node *DoubleList) {
	for node != nil {
		fmt.Printf("%d,", node.val)
		node = node.next
	}
	fmt.Println()
}

// --- 官方

/*
方法一: 哈希表 + 双向链表

时间复杂度：对于 put 和 get 都是 O(1)。
空间复杂度：O(capacity)，因为哈希表和双向链表最多存储 capacity+1 个元素。
*/
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.value = value
		this.moveToHead(node)
		return
	}

	node = initDLinkedNode(key, value)
	this.cache[key] = node
	this.addToHead(node)
	this.size++
	if this.size > this.capacity {
		removed := this.removeTail()
		delete(this.cache, removed.key)
		this.size--
	}
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next = node
	node.next.pre = node
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
