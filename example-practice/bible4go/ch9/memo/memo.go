// 并发的非阻塞缓存 -- 基于sync.Mutex

package memo

import (
	"sync"
)

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// key的第一次请求
		// 这个goroutine负责调用函数f，写入map并且广播给其他goroutine已经ready
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // 广播通知其他goroutine已经就绪
	} else {
		memo.mu.Unlock()
		<-e.ready // 等待就绪状态
	}

	return e.res.value, e.res.err
}
