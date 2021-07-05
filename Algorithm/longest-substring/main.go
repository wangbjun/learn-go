package main

import "fmt"

func main() {
	var s = "abcabcbb"
	fmt.Printf("%d\n", lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过，同时也维护了一个窗口内的字符
	m := make(map[byte]int)
	n := len(s)
	// rk右指针，ans是最大长度
	rk, ans := 0, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk < n && m[s[rk]] == 0 {
			// 不断地移动右指针
			m[s[rk]] = 1
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

//
//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//	max := 1
//	for i := 0; i < len(s); i++ {
//		for j := i + 1; j < len(s); j++ {
//			subStr := s[i : j+1]
//			if !hasRepeatChar(subStr) && len(subStr) > max {
//				max = len(subStr)
//				println(subStr)
//			}
//		}
//	}
//	return max
//}
//
//func hasRepeatChar(s string) bool {
//	for i := 0; i < len(s); i++ {
//		for j := i + 1; j < len(s); j++ {
//			if s[i] == s[j] {
//				return true
//			}
//		}
//	}
//	return false
//}
