/*
Package lru_cache
https://leetcode-cn.com/problems/lru-cache/

146. LRU 缓存

请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。

实现 LRUCacheO1 类：
	- LRUCacheO1(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
	- int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
	- void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；
	如果不存在，则向缓存中插入该组 key-value 。
	如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。

函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

提示：
	1 <= capacity <= 3000
	0 <= key <= 10000
	0 <= value <= 10^5
	最多调用 2 * 10^5 次 get 和 put

*/
package lru_cache

import "fmt"

//type LRUCacheO1 struct{}
//func Constructor(capacity int) LRUCacheO1       {}
//func (this *LRUCacheO1) Get(key int) int        {}
//func (this *LRUCacheO1) Put(key int, value int) {}

/**
 * Your LRUCacheO1 object will be instantiated and called as such:
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
	cache      map[int]*doubleListS1
	head, tail *doubleListS1
}

type doubleListS1 struct {
	key, val  int
	pre, next *doubleListS1
}

func ConstructorS1(capacity int) LRUCacheS1 {
	head := &doubleListS1{}
	tail := &doubleListS1{}
	head.next = tail
	tail.pre = head

	return LRUCacheS1{
		cap:   capacity,
		len:   0,
		cache: make(map[int]*doubleListS1, 0),
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
	node := &doubleListS1{
		key: key,
		val: value,
	}
	this.cache[key] = node
	this.len++
	this.addToHead(node)
	this.removeTail()
}

func (this *LRUCacheS1) moveToHead(node *doubleListS1) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCacheS1) removeNode(node *doubleListS1) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCacheS1) addToHead(node *doubleListS1) {
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
	// remove from doubleListS2
	//this.tail.pre = node.pre
	//node.pre.next = this.tail
	this.removeNode(node)
	// remove from cache
	delete(this.cache, node.key)
	this.len--
}

/*
方法一: 哈希表 + 双向链表

时间复杂度：
空间复杂度：
*/

type LRUCache struct {
	cap, len int
	h        map[int]*doubleListS2
	head     *doubleListS2 // 靠近头节点更新
	tail     *doubleListS2
}

type doubleListS2 struct {
	key, val  int
	pre, next *doubleListS2
}

func Constructor(capacity int) LRUCache {
	head, tail := doubleListS2{}, doubleListS2{}
	head.next = &tail
	tail.pre = &head

	return LRUCache{
		cap:  capacity,
		len:  0,
		h:    make(map[int]*doubleListS2),
		head: &head,
		tail: &tail,
	}
}
func (l *LRUCache) Get(key int) int {
	if node, ok := l.h[key]; ok {
		l.cutNode(node)
		l.moveToHead(node)
		return node.val
	}

	return -1
}

func (l *LRUCache) Put(key int, value int) {
	node, ok := l.h[key]
	if ok {
		node.val = value
		l.cutNode(node)
		l.moveToHead(node)
		return
	}

	// 溢出问题
	if l.len >= l.cap {
		cutNode := l.tail.pre
		delete(l.h, cutNode.key)
		l.cutNode(cutNode)
		l.len--
	}

	newNode := &doubleListS2{key: key, val: value}
	l.h[key] = newNode
	l.moveToHead(newNode)
	l.len++
}

func (l *LRUCache) cutNode(node *doubleListS2) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (l *LRUCache) moveToHead(node *doubleListS2) {
	node.next = l.head.next
	l.head.next.pre = node
	l.head.next = node
	node.pre = l.head
}

// --- 官方

/*
方法一: 哈希表 + 双向链表

时间复杂度：对于 put 和 get 都是 O(1)。
空间复杂度：O(capacity)，因为哈希表和双向链表最多存储 capacity+1 个元素。
*/

type LRUCacheO1 struct {
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

func ConstructorO1(capacity int) LRUCacheO1 {
	l := LRUCacheO1{
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCacheO1) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCacheO1) Put(key int, value int) {
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

func (this *LRUCacheO1) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCacheO1) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next = node
	node.next.pre = node
}

func (this *LRUCacheO1) removeTail() *DLinkedNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}

func (this *LRUCacheO1) removeNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func printList(node *doubleListS1) {
	for node != nil {
		fmt.Printf("%d,", node.val)
		node = node.next
	}
	fmt.Println()
}
