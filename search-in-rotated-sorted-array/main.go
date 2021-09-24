/*
Package search_in_rotated_sorted_array
https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

33. 搜索旋转排序数组

升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。
请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

*/
package search_in_rotated_sorted_array

// --- 自己

/*
方法一: 二分+递归

时间复杂度：
空间复杂度：
*/
func search(nums []int, target int) int {
	return searchRotated(nums, 0, len(nums)-1, target)
}

func searchRotated(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}

	if nums[low] < nums[high] {
		return binarySearch(nums, low, high, target)
	}

	mid := low + (high-low)>>1
	if nums[mid] == target {
		return mid
	}

	var ans int
	ans = searchRotated(nums, low, mid-1, target)
	if ans != -1 {
		return ans
	}

	return searchRotated(nums, mid+1, high, target)
}

func binarySearch(nums []int, low, high, target int) int {
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// --- 官方
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/sou-suo-xuan-zhuan-pai-xu-shu-zu-by-leetcode-solut/

/*
方法一: 二分查找

时间复杂度：O(logn)，其中 n 为 nums 数组的大小。整个算法时间复杂度即为二分查找的时间复杂度 O(logn)。
空间复杂度：O(1) 。我们只需要常数级别的空间存放变量。
*/
func searchO1(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
