package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

//Requests per second:   207.88 [#/sec] (mean)
// func newBuf() []byte {
// 	return make([]byte, 10<<20)
// }
//Requests per second:    2143.27 [#/sec] (mean)
var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 10<<20)
	},
}

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	http.HandleFunc("/example2", func(w http.ResponseWriter, r *http.Request) {
		// b := newBuf()
		b := bufPool.Get().([]byte)
		for idx := range b {
			b[idx] = 0
		}

		fmt.Fprintf(w, "done,%v", r.URL.Path[1:])
		bufPool.Put(b)
	})

	http.ListenAndServe(":8080", nil)
}
