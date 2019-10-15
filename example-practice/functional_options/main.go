package main

import (
	"fmt"
	"time"
)

type Handler interface {
	Do(k, v interface{})
}

type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

// 强制转型
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}

func selfInfo(k, v interface{}) {
	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	// HandlerFunc(w.selfInfo)不是方法的调用，而是转型，因为selfInfo和HandlerFunc是同一种类型，所以可以强制转型。
	// 转型后，因为HandlerFunc实现了Handler接口，所以我们就可以继续使用原来的Each方法了。
	// Each(persons, HandlerFunc(w.selfInfo))

	EachFunc(persons, selfInfo)
}

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

func Connect(addr string, opts ...Option) {
	options := options{
		timeout: 1,
		caching: false,
	}
	for _, o := range opts {
		o.apply(&options)
	}

	//...
	return
}

// db.Connect(addr, db.WithTimeout(newTimeout))
// db.Connect(addr, db.WithCaching(false))
