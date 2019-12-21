# panic and recover

### 问题

+ 为什么`panic`会中止运行

+ 为什么`defer+recover`不会中止运行

+ 不设置`defer` 行不

+ 设置 `defer` + `recover` 组合后就能无忧无虑了吗

- 为什么起个goroutine就不行



#### 数据结构

```go
type _panic struct{
    argp      unsafe.Pointer // 指向defer延迟调用的参数的指针
    arg       interface{}  // panic的原因
    link      *_panic  // 指向上一个调用的_panic
    recovered bool // 是否已经被处理
    aborted   bool // 是否被中止
}
```



`panic` 实际上就是处理当前`goroutine(g)`上所挂载的_panic链表(所以无法对其他goroutine的异常事件响应)，然后对其所属的defer链表和recover进行检测并处理，最后调用退出的命令。

