package public

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxN(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

func minN(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
