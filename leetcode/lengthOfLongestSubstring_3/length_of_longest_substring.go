package lengthOfLongestSubstring

import (
	"math"
)

/**
* 题目：无重复的字符的最长子串
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
