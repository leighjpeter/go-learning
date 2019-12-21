package lengthOfLongestSubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var s = map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
		"aab":      2,
		"dvdf":     3,
		"   ":      1,
		"au":       2,
		"asjrgapa": 6,
	}

	for k, v := range s {
		assert.Equal(t, lengthOfLongestSubstring(k), v)
	}
}
