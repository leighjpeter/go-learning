package main

import "strconv"

func main() {
	s := "babad"
	println(longestPalindrome(s))

	println(isPalindrome(12321))
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for x := range str {
		if str[x:x+1] != str[len(str)-x-1:len(str)-x] {
			return false
		}
	}
	return true
}

// 暴力法
func longestPalindrome(s string) string {
	var maxL int
	sLen := len(s)
	var res string
	for i := 0; i < sLen-1; i++ {
		for j := i + 1; j <= sLen; j++ {
			if ValidPalindrome(s[i:j]) && len(s[i:j]) > maxL {
				maxL = len(s[i:j])
				res = s[i:j]
			}
		}
	}
	return res
}

func ValidPalindrome(s string) bool {
	sLen := len(s)
	for x := range s {
		if s[x:x+1] != s[sLen-x-1:sLen-x] {
			return false
		}
	}
	return true
}
