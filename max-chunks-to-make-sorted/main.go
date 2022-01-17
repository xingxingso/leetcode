/*
Package max_chunks_to_make_sorted
https://leetcode-cn.com/problems/max-chunks-to-make-sorted/

769. 最多能完成排序的块

数组arr是[0, 1, ..., arr.length - 1]的一种排列，我们将这个数组分割成几个“块”，并将这些块分别进行排序。之后再连接起来，使得连接的结果和按升序排序后的原数组相同。

我们最多能将数组分成多少块？

注意:
	arr 的长度在 [1, 10] 之间。
	arr[i]是 [0, 1, ..., arr.length - 1]的一种排列。

*/
package max_chunks_to_make_sorted

// --- 他人

/*
方法一:

时间复杂度：
空间复杂度：
*/
func maxChunksToSorted(arr []int) int {
	chunks, cur := 0, 0
	for i := 0; i < len(arr); i++ {
		cur = max(cur, arr[i])
		if cur == i {
			chunks++
		}
	}
	return chunks
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
