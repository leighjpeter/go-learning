package lengthOfLongestSubstring

import (
	"math"
)

/**
* 题目：无重复的字符的最长子串
* abcabcbb
 */
func lengthOfLongestSubstring(str string) int {
	cnt := 0
	temp := make(map[byte]int)
	start, end := 0, 0
	sLen := len(str)

	for start < sLen && end < sLen {
		tmp := str[end]
		if _, ok := temp[tmp]; !ok {
			temp[str[end]] = 1
			end++
			cnt = int(math.Max(float64(cnt), float64(end-start)))
		} else {
			delete(temp, str[start])
			start++
		}
	}
	return cnt
}

// 解2
func LenMax(str string) int {
	rs := 0
	s := make(map[byte]int)
	for i, j := 0, 0; i < len(str); i++ {
		if v, ok := s[str[i]]; ok {
			j = max(v, j)
		}
		rs = max(rs, i-j)
		s[str[i]] = i

	}
	return rs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
