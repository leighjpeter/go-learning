package memoBsoChannel_test

import (
	"github.com/leighjpeter/go-learning/example-practice/bible4go/ch9/memoBsoChannel"
	"github.com/leighjpeter/go-learning/example-practice/bible4go/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func TestMemo(t *testing.T) {
	m := memoBsoChannel.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memoBsoChannel.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}
