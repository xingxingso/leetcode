/*
Package candy
https://leetcode-cn.com/problems/candy/

135. 分发糖果

老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
你需要按照以下要求，帮助老师给这些孩子分发糖果：
	- 每个孩子至少分配到 1 个糖果。
	- 评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？

*/
package candy

// --- 自己

/*
方法一: 模拟
	第1位只发1颗，后面依次调整

时间复杂度：
空间复杂度：
*/
func candy(ratings []int) int {
	if len(ratings) <= 0 {
		return 0
	}
	last := 1
	ans := last
	pickup := make([]int, len(ratings))
	pickup[0] = last
	for i := 1; i < len(ratings); i++ {
		cur := 1
		if ratings[i] > ratings[i-1] {
			cur = last + 1
		} else if ratings[i] == ratings[i-1] {
			cur = 1
		} else {
			if pickup[i-1] == 1 {
				j := i - 1
				for ; j >= 1; j-- {
					if ratings[j-1] > ratings[j] && pickup[j-1] == pickup[j]+1 { // 差值在1的时候补发1颗
						pickup[j]++
						ans++
					} else {
						break
					}
				}
				if j >= 0 {
					pickup[j]++
					ans++
				}
				cur = 1
			} else {
				cur = 1
			}
		}
		last = cur
		pickup[i] = last
		ans += last
		//fmt.Println(pickup)
	}
	return ans
}

/*
方法二: 贪心

时间复杂度：
空间复杂度：
*/
func candyO1(ratings []int) int {
	ans := len(ratings)
	pick := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			pick[i] = pick[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			pick[i-1] = max(pick[i]+1, pick[i-1])
		}
	}
	for i := 0; i < len(pick); i++ {
		ans += pick[i]
	}
	//fmt.Println(pick)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
