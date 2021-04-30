/*
https://leetcode-cn.com/problems/flatten-nested-list-iterator/

341. 扁平化嵌套列表迭代器

给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。
列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。

*/
package flatten_nested_list_iterator

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

// --- 自己

/*
方法一:
	!!! 属于一次性加载, 不好

时间复杂度：
空间复杂度：
*/

type NestedIterator struct {
	cache []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	if len(nestedList) == 0 {
		return &NestedIterator{}
	}

	cache := make([]int, 0)
	stack := make([]*NestedInteger, 0)
	for len(nestedList) > 0 {
		stack = append(stack, nestedList[0])
		nestedList = nestedList[1:]
		for len(stack) > 0 {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if n.IsInteger() {
				cache = append(cache, n.GetInteger())
			} else {
				list := n.GetList()
				for i := len(list) - 1; i >= 0; i-- {
					stack = append(stack, list[i])
				}
			}
		}
	}
	// fmt.Println(cache)
	return &NestedIterator{
		cache: cache,
	}
}

func (this *NestedIterator) Next() int {
	num := this.cache[0]
	this.cache = this.cache[1:]
	return num
}

func (this *NestedIterator) HasNext() bool {
	return len(this.cache) != 0
}

// --- 官方

/*
方法一: 栈

时间复杂度：初始化和 next 为 O(1)，hasNext 为均摊 O(1)。
空间复杂度：O(n)。最坏情况下嵌套的整型列表是一条链，我们需要一个 O(n) 大小的栈来存储链上的所有元素。
*/

type NestedIteratorO1 struct {
	// 将列表视作一个队列，栈中直接存储该队列
	stack [][]*NestedInteger
}

func ConstructorO1(nestedList []*NestedInteger) *NestedIteratorO1 {
	return &NestedIteratorO1{[][]*NestedInteger{nestedList}}
}

func (it *NestedIteratorO1) Next() int {
	// 由于保证调用 Next 之前会调用 HasNext，直接返回栈顶列表的队首元素，将其弹出队首并返回
	queue := it.stack[len(it.stack)-1]
	val := queue[0].GetInteger()
	it.stack[len(it.stack)-1] = queue[1:]
	return val
}

func (it *NestedIteratorO1) HasNext() bool {
	for len(it.stack) > 0 {
		queue := it.stack[len(it.stack)-1]
		if len(queue) == 0 { // 当前队列为空，出栈
			it.stack = it.stack[:len(it.stack)-1]
			continue
		}
		nest := queue[0]
		if nest.IsInteger() {
			return true
		}
		// 若队首元素为列表，则将其弹出队列并入栈
		it.stack[len(it.stack)-1] = queue[1:]
		it.stack = append(it.stack, nest.GetList())
	}
	return false
}
