/*
https://leetcode-cn.com/problems/ip-to-cidr/

751. IP 到 CIDR

给给定一个起始 IP 地址 ip 和一个我们需要包含的 IP 的数量 n，返回用列表（最小可能的长度）表示的 CIDR块的范围。
CIDR 块是包含 IP 的字符串，后接斜杠和固定长度。例如：“123.45.67.89/20”。固定长度 “20” 表示在特定的范围中公共前缀位的长度。

注：
	1. ip 是有效的 IPv4 地址。
	2. 每一个隐含地址 ip + x (其中 x < n) 都是有效的 IPv4 地址。
	3. n 为整数，范围为 [1, 1000]。

*/
package ip_to_cidr

import (
	"fmt"
	"strconv"
	"strings"
)

// --- 官方

/*
方法一:

时间复杂度：O(N)。其中 N 表示的是 nums 的长度
空间复杂度：O(1)。
*/
func ipToCIDR(ip string, n int) []string {
	start := ipToInt(ip)
	ans := make([]string, 0)
	for n > 0 {
		mask := max(33-bitLength(start&-start), 33-bitLength(n))
		ans = append(ans, intToIp(start)+"/"+strconv.Itoa(mask))
		start += 1 << (32 - mask)
		n -= 1 << (32 - mask)
	}
	return ans
}

func ipToInt(ip string) int {
	ans := 0
	for _, v := range strings.Split(ip, ".") {
		num, _ := strconv.Atoi(v)
		ans = 256*ans + num
	}
	return ans
}

func intToIp(x int) string {
	return fmt.Sprintf("%d.%d.%d.%d", x>>24, (x>>16)%256, (x>>8)%256, x%256)
}

func bitLength(x int) int {
	if x == 0 {
		return 1
	}
	ans := 0
	for x > 0 {
		x >>= 1
		ans++
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
