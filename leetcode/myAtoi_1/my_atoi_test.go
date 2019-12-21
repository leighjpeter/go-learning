package myAtoi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// func TestNew(t *testing.T) {
// 	c, err := New("tesla", 100)
// 	if err != nil {
// 		t.Error("got errors", err)
// 	}

// 	if c == nil {
// 		t.Error("car should be nil")
// 	}
// }

func TestMyAtoi(t *testing.T) {
	var s = map[string]int{
		"42":              42,
		"    -43":         -43,
		"4193 with words": 4193,
		"words and 987":   0,
		"-91283472332":    -2147483648,
	}

	for k, v := range s {
		assert.Equal(t, myAtoi(k), v)
	}

}
