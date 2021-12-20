/*
Package count_primes
https://leetcode-cn.com/problems/count-primes/

204. 计数质数

统计所有小于非负整数 n 的质数的数量。

提示：
	0 <= n <= 5 * 10^6

*/
package count_primes

// --- 自己

/*
方法一:
	// 超出时限

时间复杂度：
空间复杂度：
*/
func countPrimes(n int) int {
	ans := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			ans++
		}
	}
	return ans
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/*
方法一: 记忆化
	勉强通过

时间复杂度：
空间复杂度：
*/
func countPrimesS2(n int) int {
	primes := make([]int, 0)
	for i := 2; i < n; i++ {
		flag := false
		for _, j := range primes {
			if i%j == 0 {
				flag = true
				break
			}
			if j*j > i {
				break
			}
		}
		if !flag {
			primes = append(primes, i)
		}
	}
	return len(primes)
}

// --- 官方

/*
方法二: 埃氏筛

时间复杂度：
空间复杂度：
*/
func countPrimesO1(n int) int {
	cnt := 0
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return cnt
}
