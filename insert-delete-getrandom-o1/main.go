/*
https://leetcode-cn.com/problems/insert-delete-getrandom-o1/

380. O(1) 时间插入、删除和获取随机元素

设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构。
	1. insert(val)：当元素 val 不存在时，向集合中插入该项。
	2. remove(val)：元素 val 存在时，从集合中移除该项。
	3. getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。

*/
package insert_delete_getrandom_o1

import (
	"math/rand"
	"time"
)

// type RandomizedSet struct {}
// /** Initialize your data structure here. */
// func Constructor() RandomizedSet {}
// /** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
// func (this *RandomizedSet) Insert(val int) bool {}
// /** Removes a value from the set. Returns true if the set contained the specified element. */
// func (this *RandomizedSet) Remove(val int) bool {}
// /** Get a random element from the set. */
// func (this *RandomizedSet) GetRandom() int {}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

// --- 自己

/*
方法一:
	labuladong
	对于getRandom方法，如果想「等概率」且「在 O(1) 的时间」取出元素，一定要满足：底层用数组实现，且数组必须是紧凑的。
	所以，如果我们想在 O(1) 的时间删除数组中的某一个元素val，可以先把这个元素交换到数组的尾部，然后再pop掉。

时间复杂度：
空间复杂度：
*/
type RandomizedSet struct {
	hash map[int]int
	arr  []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		hash: map[int]int{},
		arr:  []int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hash[val]; ok {
		return false
	}
	this.hash[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.hash[val]
	if !ok {
		return false
	}
	delete(this.hash, val)
	end := len(this.arr) - 1
	// 交换到数组尾部
	if index != end {
		this.arr[index], this.arr[end] = this.arr[end], this.arr[index]
		// !!! 同时要交换对应 hash 中 的下标
		this.hash[this.arr[index]] = index
	}
	this.arr = this.arr[:end]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(this.arr))
	return this.arr[index]
}
