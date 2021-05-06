/*
https://leetcode-cn.com/problems/lfu-cache/

460. LFU 缓存

请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
实现 LFUCache 类：

	LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
	int get(int key) - 如果键存在于缓存中，则获取键的值，否则返回 -1。
	void put(int key, int value) - 如果键已存在，则变更其值；如果键不存在，请插入键值对。
	当缓存达到其容量时，则应该在插入新项之前，使最不经常使用的项无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，
	应该去除 最近最久未使用 的键。

注意「项的使用次数」就是自插入该项以来对其调用 get 和 put 函数的次数之和。使用次数会在对应项被移除后置为 0 。
为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

提示：
	0 <= capacity, key, value <= 10^4
	最多调用 10^5 次 get 和 put 方法

进阶：
	你可以为这两种操作设计时间复杂度为 O(1) 的实现吗？

*/
package lfu_cache

import "fmt"

//type LFUCache struct{}
//func Constructor(capacity int) LFUCache       {}
//func (this *LFUCache) Get(key int) int        {}
//func (this *LFUCache) Put(key int, value int) {}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 *

// --- 官方

/*
方法二: 双哈希表

时间复杂度：get 时间复杂度 O(1)，put 时间复杂度 O(1)。由于两个操作从头至尾都只利用了哈希表的插入删除还有链表的插入删除，
	且它们的时间复杂度均为 O(1)，所以保证了两个操作的时间复杂度均为 O(1)。
空间复杂度：O(capacity)，其中 capacity 为 LFU 的缓存容量。哈希表中不会存放超过缓存容量的键值对。
*/
type LFUCache struct {
	cap       int
	minFreq   int
	keyTable  map[int]*LinkedList
	freqTable map[int]*Node
}

type Node struct {
	head, tail *LinkedList
}

type LinkedList struct {
	key, val, freq int
	pre, next      *LinkedList
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:       capacity,
		keyTable:  map[int]*LinkedList{},
		freqTable: map[int]*Node{},
	}
}

func (this *LFUCache) Get(key int) int {
	if this.cap == 0 {
		return -1
	}

	list, ok := this.keyTable[key]
	if !ok {
		return -1
	}
	this.removeList(list)
	node := this.freqTable[list.freq]
	if node.head.next.val == -1 {
		//this.freqTable[list.freq] = nil
		delete(this.freqTable, list.freq)
		if list.freq == this.minFreq {
			this.minFreq++
		}
	}
	list.freq++
	this.addFreqNode(list.freq, list)
	return list.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	list, ok := this.keyTable[key]
	if ok {
		list.val = value
		this.Get(key)
		return
	}
	list = &LinkedList{
		key:  key,
		val:  value,
		freq: 1,
	}
	this.keyTable[key] = list
	this.addFreqNode(1, list)

	// 超出容量
	if len(this.keyTable) > this.cap {
		this.removeMinFreqKey(this.minFreq)
	}
	this.minFreq = 1
}

func (this *LFUCache) removeList(list *LinkedList) {
	list.pre.next = list.next
	list.next.pre = list.pre
}

func (this *LFUCache) addFreqNode(freq int, list *LinkedList) {
	node := this.freqTable[freq]
	if node == nil {
		node = this.initNode()
		this.freqTable[freq] = node
	}
	list.next = node.head.next
	list.pre = node.head
	node.head.next = list
	list.next.pre = list
}

func (this *LFUCache) removeMinFreqKey(freq int) {
	node := this.freqTable[freq]
	list := node.tail.pre
	this.removeList(list)
	// node 中 list 为空 head 直接于 tail 相连
	if node.head.next.val == -1 {
		//this.freqTable[freq] = nil
		delete(this.freqTable, list.freq)
		//node = nil
	}
	delete(this.keyTable, list.key)
}

func (this *LFUCache) initNode() *Node {
	head := &LinkedList{val: -1}
	tail := &LinkedList{val: -1}
	head.next = tail
	tail.pre = head
	return &Node{
		head: head,
		tail: tail,
	}
}

func printLFUCache(cache *LFUCache) {
	fmt.Printf("cap:%d\n", cache.cap)
	fmt.Printf("minFreq:%d\n", cache.minFreq)
	fmt.Println("keyTable:")
	for k, v := range cache.keyTable {
		fmt.Printf("key:%d,value:%v\n", k, v)
	}
	for freq, node := range cache.freqTable {
		fmt.Printf("freq:%d\n", freq)
		if node == nil {
			continue
		}
		head := node.head
		for head != nil {
			fmt.Printf("node: key:%d,value:%d,freq:%d\n", head.key, head.val, head.freq)
			head = head.next
		}
	}
	fmt.Println("printLFUCache Finish")
}
