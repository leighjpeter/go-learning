package memo_test

import (
	"github.com/leighjpeter/go-learning/example-practice/bible4go/ch9/memo"
	"github.com/leighjpeter/go-learning/example-practice/bible4go/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func TestMemo(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
