/*
Package can_place_flowers
https://leetcode-cn.com/problems/can-place-flowers/

605. 种花问题

假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。
可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。
另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false。

提示：
    1 <= flowerbed.length <= 2 * 10^4
    flowerbed[i] 为 0 或 1
    flowerbed 中不存在相邻的两朵花
    0 <= n <= flowerbed.length

*/
package can_place_flowers

// --- 自己

/*
方法一: 贪心

时间复杂度：
空间复杂度：
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed) && count <= n; i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if i > 0 && flowerbed[i-1] == 1 {
			continue
		}
		if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
			continue
		}
		flowerbed[i] = 1
		count++
	}
	return count >= n
}

// --- 官方

/*
方法一: 贪心

时间复杂度：
空间复杂度：
*/
func canPlaceFlowersO1(flowerbed []int, n int) bool {
	count := 0
	m := len(flowerbed)
	prev := -1
	for i := 0; i < m; i++ {
		if flowerbed[i] == 1 {
			if prev < 0 {
				count += i / 2
			} else {
				count += (i - prev - 2) / 2
			}
			if count >= n {
				return true
			}
			prev = i
		}
	}
	if prev < 0 {
		count += (m + 1) / 2
	} else {
		count += (m - prev - 1) / 2
	}
	return count >= n
}

// --- 他人

/*
方法一: 贪心
    @mesmerizeBy 只判断后面的位置有没有种花就够了

时间复杂度：
空间复杂度：
*/
func canPlaceFlowersP1(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			i++
		} else if i+1 == len(flowerbed) || flowerbed[i+1] == 0 {
			n--
			i++
		}
	}
	return n <= 0
}
