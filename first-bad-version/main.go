/*
Package first_bad_version
https://leetcode-cn.com/problems/first-bad-version/

278. 第一个错误的版本

你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。
由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。

假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。

你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。
你应该尽量减少对调用 API 的次数。

提示：
	1 <= bad <= n <= 2^31 - 1

*/
package first_bad_version

//* Forward declaration of isBadVersion API.
//* @param   version   your guess about first bad version
//* @return 	 	      true if current version is bad
//*			          false if current version is good
//* func isBadVersion(version int) bool;

// --- 自己

/*
方法一: 二分查找

时间复杂度：
空间复杂度：
*/
func firstBadVersion(n int) int {
	l, r := 1, n
	for l <= r {
		//fmt.Println(l, r)
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

var bad int

func isBadVersion(version int) bool {
	if version < bad {
		return false
	}
	return true
}
