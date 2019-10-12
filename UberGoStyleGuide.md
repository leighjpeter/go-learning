# Uber Go 代码风格指南

[英文原版](https://github.com/uber-go/guide/blob/master/style.md)

## 目录

- [介绍](#介绍)
- [指南](#指南)
  - [Pointers to Interfaces](#pointers-to-interfaces)
  - [Receivers and Interfaces](#receivers-and-interfaces)
  - [Zero-value Mutexes are Valid](#zero-value-mutexes-are-valid)
  - [Copy Slices and Maps at Boundaries](#copy-slices-and-maps-at-boundaries)
  - [Defer to Clean Up](#defer-to-clean-up)
  - [Channel Size is One or None](#channel-size-is-one-or-none)
  - [Start Enums at One](#start-enums-at-one)
  - [Error Types](#error-types)
  - [Error Wrapping](#error-wrapping)
  - [Handle Type Assertion Failures](#handle-type-assertion-failures)
  - [Don't Panic](#dont-panic)
  - [Use go.uber.org/atomic](#use-gouberorgatomic)
- [性能](#performance)
  - [Prefer strconv over fmt](#prefer-strconv-over-fmt)
  - [Avoid string-to-byte conversion](#avoid-string-to-byte-conversion)
- [风格](#style)
  - [Group Similar Declarations](#group-similar-declarations)
  - [Import Group Ordering](#import-group-ordering)
  - [Package Names](#package-names)
  - [Function Names](#function-names)
  - [Import Aliasing](#import-aliasing)
  - [Function Grouping and Ordering](#function-grouping-and-ordering)
  - [Reduce Nesting](#reduce-nesting)
  - [Unnecessary Else](#unnecessary-else)
  - [Top-level Variable Declarations](#top-level-variable-declarations)
  - [Prefix Unexported Globals with _](#prefix-unexported-globals-with-_)
  - [Embedding in Structs](#embedding-in-structs)
  - [Use Field Names to initialize Structs](#use-field-names-to-initialize-structs)
  - [Local Variable Declarations](#local-variable-declarations)
  - [nil is a valid slice](#nil-is-a-valid-slice)
  - [Reduce Scope of Variables](#reduce-scope-of-variables)
  - [Avoid Naked Parameters](#avoid-naked-parameters)
  - [Use Raw String Literals to Avoid Escaping](#use-raw-string-literals-to-avoid-escaping)
  - [Initializing Struct References](#initializing-struct-references)
  - [Format Strings outside Printf](#format-strings-outside-printf)
  - [Naming Printf-style Functions](#naming-printf-style-functions)
- [格式布局](#patterns)
  - [Test Tables](#test-tables)
  - [Functional Options](#functional-options)

## 介绍
Style是控制代码的约定。术语`样式`有点用词不当，因为这些约定不仅仅涵盖源文件.formatting-gofmt为我们处理这个问题。

本指南的目标是通过详细描述在uber编写go代码的注意事项。这些规则是为了代码的基本可管理，同时仍允许工程师富有成效地使用go语言特性。

本指南最初由[Prashant Varanasi]和[Simon Newton]创建，作为一种让同事们加快使用go的的方法。多年来根据其他人的反馈修改。

本文档记录了我们在Uber遵循的Go代码中的惯用约定。 很多是Go的一般准则，而其他准则则适用于外部资源：

1. [Effective Go](https://golang.org/doc/effective_go.html)
2. [The Go common mistakes guide](https://github.com/golang/go/wiki/CodeReviewComments)

当运行“golint”和“go vet”时，所有代码都应无错误。我们建议将编辑器设置为：

- Run `goimports` on save
- Run `golint` and `go vet` to check for errors

您可以在此处对Go工具的编辑器支持中找到信息:

https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins


## 指南

### Pointers to Interfaces

你几乎不需要一个指向接口的指针。应该将接口作为值传递——其底层数据还是一个指针。

接口是两个字段：

1. 指向某些特定类型信息的指针，可以将其视为“类型”
2. 数据指针：如果存储的数据是指针，则直接存储。如果存储的数据是一个值，则存储指向该值的指针。

如果要接口方法修改基础数据，必须使用指针。

### Receivers and Interfaces

可以对值和指针调用具有值类型接收器的方法

举例：

```go
type S struct {
  data string
}

func (s S) Read() string {
  return s.data
}

func (s *S) Write(str string) {
  s.data = str
}

sVals := map[int]S{1: {"A"}}

// You can only call Read using a value
sVals[1].Read()

// This will not compile:
//  sVals[1].Write("test")

sPtrs := map[int]*S{1: {"A"}}

// You can call both Read and Write using a pointer
sPtrs[1].Read()
sPtrs[1].Write("test")
```

同样，即使该方法具有值接收器，也可以通过指针来满足接口。

```go
type F interface {
  f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

s1Val := S1{}
s1Ptr := &S1{}
s2Val := S2{}
s2Ptr := &S2{}

var i F
i = s1Val
i = s1Ptr
i = s2Ptr

// The following doesn't compile, since s2Val is a value, and there is no value receiver for f.
//   i = s2Val
```

Effective Go 有好的写法关于 [Pointers vs. Values].

[Pointers vs. Values]: https://golang.org/doc/effective_go.html#pointers_vs_values

### Zero-value Mutexes are Valid

sync.Mutex和sync.RWMutex的零值是有效的，因此几乎不需要指向mutex的指针。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
mu := new(sync.Mutex)
mu.Lock()
```
</td><td>

```go
var mu sync.Mutex
mu.Lock()
```
</td></tr>
</tbody></table>


如果通过指针使用结构体，那么mutex可以不是指针字段，或者最好是直接嵌入到结构体中

<table>
<tbody>
<tr><td>

```go
type smap struct {
  sync.Mutex

  data map[string]string
}

func newSMap() *smap {
  return &smap{
    data: make(map[string]string),
  }
}

func (m *smap) Get(k string) string {
  m.Lock()
  defer m.Unlock()

  return m.data[k]
}
```

</td><td>

```go
type SMap struct {
  mu sync.Mutex

  data map[string]string
}

func NewSMap() *SMap {
  return &SMap{
    data: make(map[string]string),
  }
}

func (m *SMap) Get(k string) string {
  m.mu.Lock()
  defer m.mu.Unlock()

  return m.data[k]
}
```

</td></tr>

</tr>
<tr>
<td>为私有类型或者需要实现的nutex接口的类型嵌入</td>
<td>对于可见类型，使用私有锁</td>
</tr>

</tbody></table>

### Copy Slices and Maps at Boundaries

Slices 和 Maps 包含指向底层数据的指针，所以在复制它们时需要特别小心

#### Receiving Slices and Maps

请记住，如果您存储了对Map或Slice的引用，则可以对其进行修改。

<table>
<thead><tr><th>差劲👎</th> <th>优秀👍</th></tr></thead>
<tbody>
<tr>
<td>

```go
func (d *Driver) SetTrips(trips []Trip) {
  d.trips = trips
}

trips := ...
d1.SetTrips(trips)

// Did you mean to modify d1.trips?
trips[0] = ...
```

</td>
<td>

```go
func (d *Driver) SetTrips(trips []Trip) {
  d.trips = make([]Trip, len(trips))
  copy(d.trips, trips)
}

trips := ...
d1.SetTrips(trips)

// We can now modify trips[0] without affecting d1.trips.
trips[0] = ...
```

</td>
</tr>

</tbody>
</table>

#### Returning Slices and Maps

同样的，需要小心修改结构体内部的map或者slice

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
type Stats struct {
  sync.Mutex

  counters map[string]int
}

// Snapshot returns the current stats.
func (s *Stats) Snapshot() map[string]int {
  s.Lock()
  defer s.Unlock()

  return s.counters
}

// snapshot is no longer protected by the lock!
snapshot := stats.Snapshot()
```

</td><td>

```go
type Stats struct {
  sync.Mutex

  counters map[string]int
}

func (s *Stats) Snapshot() map[string]int {
  s.Lock()
  defer s.Unlock()

  result := make(map[string]int, len(s.counters))
  for k, v := range s.counters {
    result[k] = v
  }
  return result
}

// Snapshot is now a copy.
snapshot := stats.Snapshot()
```

</td></tr>
</tbody></table>

### Defer to Clean Up

使用defer清理资源比如文件和锁

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
p.Lock()
if p.count < 10 {
  p.Unlock()
  return p.count
}

p.count++
newCount := p.count
p.Unlock()

return newCount

// easy to miss unlocks due to multiple returns
```

</td><td>

```go
p.Lock()
defer p.Unlock()

if p.count < 10 {
  return p.count
}

p.count++
return p.count

// more readable
```

</td></tr>
</tbody></table>


Defer的开销非常小，只有在能保证方法执行时间是纳秒级的情况下可以不使用。
与使用defer的可读性相比，其开销忽略不计，尤其适用于具有比简单的内存访问更多更大的方法，其中其他计算比defer更大

### Channel Size is One or None

Channel的大小通常应为1或者无缓冲。默认情况下，channel是无缓冲的，大小为0。任何其他大小都必须经过审查，考虑大小是如何确定的，如何防止channel在负载下填充并阻塞写入的原因，以及会引起什么问题

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
// Ought to be enough for anybody!
c := make(chan int, 64)
```

</td><td>

```go
// Size of one
c := make(chan int, 1) // or
// Unbuffered channel, size of zero
c := make(chan int)
```

</td></tr>
</tbody></table>


### Start Enums at One

在Go中引入枚举的标准方法是声明一个自定义类型和带有iota的const组。 由于变量的默认值为0，因此通常应该以非零值开始枚举。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
type Operation int

const (
  Add Operation = iota
  Subtract
  Multiply
)

// Add=0, Subtract=1, Multiply=2
```

</td><td>

```go
type Operation int

const (
  Add Operation = iota + 1
  Subtract
  Multiply
)

// Add=1, Subtract=2, Multiply=3
```

</td></tr>
</tbody></table>

在某些情况下，使用零值是有意义的，例如例子中零值是理想的默认值

```go
type LogOutput int

const (
  LogToStdout LogOutput = iota
  LogToFile
  LogToRemote
)

// LogToStdout=0, LogToFile=1, LogToRemote=2
```

<!-- TODO: section on String methods for enums -->

### Error Types

声明错误有多种形式：

- [`errors.New`] for errors with simple static strings
- [`fmt.Errorf`] for formatted error strings
- Custom types that implement an `Error()` method
- Wrapped errors using [`"pkg/errors".Wrap`]

返回错误时，考虑一下因素以确定最佳选择：


- 是否时简单的错误不需要额外信息？是，则[`errors.New`]是最合适的。

- 客户端是否需要监听和处理这个错误？是，则你需要使用自定义类型，且实现`Error()`方法

- 是否需要传递下游函数的错误？是，则参考[section on error wrapping](#error-wrapping).

- 其他，[`fmt.Errorf`] okey.

  [`errors.New`]: https://golang.org/pkg/errors/#New
  [`fmt.Errorf`]: https://golang.org/pkg/fmt/#Errorf
  [`"pkg/errors".Wrap`]: https://godoc.org/github.com/pkg/errors#Wrap

如果客户端需要检测错误，并且您已经创建了一个简单的错误，使用[`errors.new`]，并为错误使用一个变量

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
// package foo

func Open() error {
  return errors.New("could not open")
}

// package bar

func use() {
  if err := foo.Open(); err != nil {
    if err.Error() == "could not open" {
      // handle
    } else {
      panic("unknown error")
    }
  }
}
```

</td><td>

```go
// package foo

var ErrCouldNotOpen = errors.New("could not open")

func Open() error {
  return ErrCouldNotOpen
}

// package bar

if err := foo.Open(); err != nil {
  if err == foo.ErrCouldNotOpen {
    // handle
  } else {
    panic("unknown error")
  }
}
```

</td></tr>
</tbody></table>

如果您有一个客户端可能需要检测的错误，并且您希望添加更多信息，那么您应该使用自定义类型。


<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
func open(file string) error {
  return fmt.Errorf("file %q not found", file)
}

func use() {
  if err := open(); err != nil {
    if strings.Contains(err.Error(), "not found") {
      // handle
    } else {
      panic("unknown error")
    }
  }
}
```

</td><td>

```go
type errNotFound struct {
  file string
}

func (e errNotFound) Error() string {
  return fmt.Sprintf("file %q not found", e.file)
}

func open(file string) error {
  return errNotFound{file: file}
}

func use() {
  if err := open(); err != nil {
    if _, ok := err.(errNotFound); ok {
      // handle
    } else {
      panic("unknown error")
    }
  }
}
```

</td></tr>
</tbody></table>

请小心直接暴露自定义错误类型，因为它们已成为包的公共api。最好是提供匹配函数来检查错误。

```go
// package foo

type errNotFound struct {
  file string
}

func (e errNotFound) Error() string {
  return fmt.Sprintf("file %q not found", e.file)
}

func IsNotFoundError(err error) bool {
  _, ok := err.(errNotFound)
  return ok
}

func Open(file string) error {
  return errNotFound{file: file}
}

// package bar

if err := foo.Open("foo"); err != nil {
  if foo.IsNotFoundError(err) {
    // handle
  } else {
    panic("unknown error")
  }
}
```

<!-- TODO: Exposing the information to callers with accessor functions. -->

### Error Wrapping

当发生fails时，主要有三种方式传递错误：

- 如果没有要添加的其他上下文，并且您想要维护原始错误类型，则返回原始错误。

- 使用[`"pkg/errors".Wrap`]给error信息增加上下文，[`"pkg/errors".Cause`]可用于提取错误信息

- 如果调用者不需要监听和处理错误当情况，使用 [`fmt.Errorf`]


建议在可能的情况下添加上下文，这样就不会出现诸如“connection rejected”之类的模糊错误，而是出现诸如“call service foo:connection refused”之类的更有用的错误。

将上下文添加到返回的错误时，请避免使用“failed to”这样的短语，以保持上下文的简洁性，这些短语表示明显的错误并随着错误在堆栈中的渗透而堆积：

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
s, err := store.New()
if err != nil {
    return fmt.Errorf(
        "failed to create new store: %s", err)
}
```

</td><td>

```go
s, err := store.New()
if err != nil {
    return fmt.Errorf(
        "new store: %s", err)
}
```

<tr><td>

```
failed to x: failed to y: failed to create new store: the error
```

</td><td>

```
x: y: new store: the error
```

</td></tr>
</tbody></table>


但是，一旦将错误发送到另一个系统，就应该清楚该消息是一个错误（例如，错误标记或日志中的“失败”前缀）。

参考 [Don't just check errors, handle them gracefully].

  [`"pkg/errors".Cause`]: https://godoc.org/github.com/pkg/errors#Cause
  [Don't just check errors, handle them gracefully]: https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

### Handle Type Assertion Failures

The single return value form of a [type assertion] will panic on an incorrect
type. Therefore, always use the "comma ok" idiom.
类型断言的单值返回将会对不正确的类型引发panic，因此使用“comma ok”形式是必要的。

  [type assertion]: https://golang.org/ref/spec#Type_assertions

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
t := i.(string)
```

</td><td>

```go
t, ok := i.(string)
if !ok {
  // handle the error gracefully
}
```

</td></tr>
</tbody></table>

<!-- TODO: There are a few situations where the single assignment form is
fine. -->

### Don't Panic

在生产环境避免panics Panics是[cascading failures]的主要来源. 如果发生错误，函数必须返回错误，且让调用者决定如何处理它I

  [cascading failures]: https://en.wikipedia.org/wiki/Cascading_failure

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
func foo(bar string) {
  if len(bar) == 0 {
    panic("bar must not be empty")
  }
  // ...
}

func main() {
  if len(os.Args) != 2 {
    fmt.Println("USAGE: foo <bar>")
    os.Exit(1)
  }
  foo(os.Args[1])
}
```

</td><td>

```go
func foo(bar string) error {
  if len(bar) == 0
    return errors.New("bar must not be empty")
  }
  // ...
  return nil
}

func main() {
  if len(os.Args) != 2 {
    fmt.Println("USAGE: foo <bar>")
    os.Exit(1)
  }
  if err := foo(os.Args[1]); err != nil {
    panic(err)
  }
}
```

</td></tr>
</tbody></table>

Panic/recover 不是错误处理策略。仅当发生不可恢复的事情（例如取消nil引用），程序才必须panic。例外情况是程序初始化：程序启动时出现的错误，应中止程序可能导致恐慌。

```go
var _statusTemplate = template.Must(template.New("name").Parse("_statusHTML"))
```

即使在测试中，也要优先选择t.Fatal或t.FailNow来代替恐慌，以确保测试标记为失败。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
// func TestFoo(t *testing.T)

f, err := ioutil.TempFile("", "test")
if err != nil {
  panic("failed to set up test")
}
```

</td><td>

```go
// func TestFoo(t *testing.T)

f, err := ioutil.TempFile("", "test")
if err != nil {
  t.Fatal("failed to set up test")
}
```

</td></tr>
</tbody></table>

<!-- TODO: Explain how to use _test packages. -->

### Use go.uber.org/atomic


使用[sync/atomic]包的原子操作对原始类型（`int32`，`int64`等）进行操作，因此很容易忘记使用原子操作来读取或修改变量。

[go.uber.org/atomic] 通过隐藏底层的类型，为这些操作增加类型安全性. 另外，它包括了一个方便的 `atomic.Bool` 类型

  [go.uber.org/atomic]: https://godoc.org/go.uber.org/atomic
  [sync/atomic]: https://golang.org/pkg/sync/atomic/

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
type foo struct {
  running int32  // atomic
}

func (f* foo) start() {
  if atomic.SwapInt32(&f.running, 1) == 1 {
     // already running…
     return
  }
  // start the Foo
}

func (f *foo) isRunning() bool {
  return f.running == 1  // race!
}
```

</td><td>

```go
type foo struct {
  running atomic.Bool
}

func (f *foo) start() {
  if f.running.Swap(true) {
     // already running…
     return
  }
  // start the Foo
}

func (f *foo) isRunning() bool {
  return f.running.Load()
}
```

</td></tr>
</tbody></table>

## Performance

Performance-specific guidelines apply only to the hot path.

### Prefer strconv over fmt

当原始类型和string互转时 `strconv` is faster than
`fmt`.

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
for i := 0; i < b.N; i++ {
  s := fmt.Sprint(rand.Int())
}
```

</td><td>

```go
for i := 0; i < b.N; i++ {
  s := strconv.Itoa(rand.Int())
}
```

</td></tr>
<tr><td>

```
BenchmarkFmtSprint-4    143 ns/op    2 allocs/op
```

</td><td>

```
BenchmarkStrconv-4    64.2 ns/op    1 allocs/op
```

</td></tr>
</tbody></table>

### Avoid string-to-byte conversion

不要重复从固定字符串创建byte slice。相反，请执行一次转换并捕获结果。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
for i := 0; i < b.N; i++ {
  w.Write([]byte("Hello world"))
}
```

</td><td>

```go
data := []byte("Hello world")
for i := 0; i < b.N; i++ {
  w.Write(data)
}
```

</tr>
<tr><td>

```
BenchmarkBad-4   50000000   22.2 ns/op
```

</td><td>

```
BenchmarkGood-4  500000000   3.25 ns/op
```

</td></tr>
</tbody></table>

## Style

### Group Similar Declarations

Go支持对相似的声明进行分组

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
import "a"
import "b"
```

</td><td>

```go
import (
  "a"
  "b"
)
```

</td></tr>
</tbody></table>

同样适用于常量，变量，类型声明。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go

const a = 1
const b = 2



var a = 1
var b = 2



type Area float64
type Volume float64
```

</td><td>

```go
const (
  a = 1
  b = 2
)

var (
  a = 1
  b = 2
)

type (
  Area float64
  Volume float64
)
```

</td></tr>
</tbody></table>

仅与组相关的声明。不要对不相关的声明进行分组。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
type Operation int

const (
  Add Operation = iota + 1
  Subtract
  Multiply
  ENV_VAR = "MY_ENV"
)
```

</td><td>

```go
type Operation int

const (
  Add Operation = iota + 1
  Subtract
  Multiply
)

const ENV_VAR = "MY_ENV"
```

</td></tr>
</tbody></table>

组不限于可以使用的地方。例如，可以在函数内部使用它们。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
func f() string {
  var red = color.New(0xff0000)
  var green = color.New(0x00ff00)
  var blue = color.New(0x0000ff)

  ...
}
```

</td><td>

```go
func f() string {
  var (
    red   = color.New(0xff0000)
    green = color.New(0x00ff00)
    blue  = color.New(0x0000ff)
  )

  ...
}
```

</td></tr>
</tbody></table>

### Import Group Ordering

应该是两个分组：

- 标准库
- 其他

goimports的默认分组

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
import (
  "fmt"
  "os"
  "go.uber.org/atomic"
  "golang.org/x/sync/errgroup"
)
```

</td><td>

```go
import (
  "fmt"
  "os"

  "go.uber.org/atomic"
  "golang.org/x/sync/errgroup"
)
```

</td></tr>
</tbody></table>

### Package Names

包命名, 选择名字参考如下:

- 小写。不要大写和下划线
- 绝大多数时候不需要去重命名导入的包
- 简短精炼。在每个call site，名称必须是完整标识的
- 不要复数形式。如 `net/url`, 而不是 `net/urls`.
- 不要"common", "util", "shared", or "lib"。这些是差劲的，无用的名称。

参考 [Package Names] and [Style guideline for Go packages].

  [Package Names]: https://blog.golang.org/package-names
  [Style guideline for Go packages]: https://rakyll.org/style-packages/

### Function Names

We follow the Go community's convention of using [MixedCaps for function
names]. An exception is made for test functions, which may contain underscores
for the purpose of grouping related test cases, 
我们遵循Go社区关于使用[MixedCaps for function names]的约定。测试功能有一个例外，为了对相关的测试用例进行分组，可能包含下划线e.g.,
`TestMyFunction_WhatIsBeingTested`.

  [MixedCaps for function names]: https://golang.org/doc/effective_go.html#mixed-caps

### Import Aliasing

如果程序包名称与导入路径的最后一个元素不匹配，则必须使用导入别名

```go
import (
  "net/http"

  client "example.com/client-go"
  trace "example.com/trace/v2"
)
```

在所有其他情况下，应避免导入别名，除非import包之间存在直接冲突。



<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
import (
  "fmt"
  "os"


  nettrace "golang.net/x/trace"
)
```

</td><td>

```go
import (
  "fmt"
  "os"
  "runtime/trace"

  nettrace "golang.net/x/trace"
)
```

</td></tr>
</tbody></table>

### Function Grouping and Ordering

- 函数应粗略的按调用顺序排序。 
- 文件中的函数应按接收者分组。

Therefore, exported functions should appear first in a file, after
`struct`, `const`, `var` definitions.
因此，可见函数应该出现在文件前部，在`struct`, `const`, `var`定义之后。

A `newXYZ()`/`NewXYZ()` may appear after the type is defined, but before the
rest of the methods on the receiver.

`newXYZ()`/`NewXYZ()` 应该出现在类型定义之后，但是要在receiver的其余方法之前。

由于功能是按receiver分组的，因此普通实用程序功能应在文件末尾出现。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
func (s *something) Cost() {
  return calcCost(s.weights)
}

type something struct{ ... }

func calcCost(n int[]) int {...}

func (s *something) Stop() {...}

func newSomething() *something {
    return &something{}
}
```

</td><td>

```go
type something struct{ ... }

func newSomething() *something {
    return &something{}
}

func (s *something) Cost() {
  return calcCost(s.weights)
}

func (s *something) Stop() {...}

func calcCost(n int[]) int {...}
```

</td></tr>
</tbody></table>

### Reduce Nesting

代码应通过尽可能先处理错误情况/特殊情况并尽早返回或继续循环来减少嵌套。减少嵌套多个级别的代码量。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
for _, v := range data {
  if v.F1 == 1 {
    v = process(v)
    if err := v.Call(); err == nil {
      v.Send()
    } else {
      return err
    }
  } else {
    log.Printf("Invalid v: %v", v)
  }
}
```

</td><td>

```go
for _, v := range data {
  if v.F1 != 1 {
    log.Printf("Invalid v: %v", v)
    continue
  }

  v = process(v)
  if err := v.Call(); err != nil {
    return err
  }
  v.Send()
}
```

</td></tr>
</tbody></table>

### Unnecessary Else

如果在if的两个分支中都设置了变量，则可以将其替换为单个if

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
var a int
if b {
  a = 100
} else {
  a = 10
}
```

</td><td>

```go
a := 10
if b {
  a = 100
}
```

</td></tr>
</tbody></table>

### Top-level Variable Declarations



顶层，使用标准的`var`关键字。请勿指定类型，除非它与表达式的类型不同。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
var _s string = F()

func F() string { return "A" }
```

</td><td>

```go
var _s = F()
// Since F already states that it returns a string, we don't need to specify
// the type again.

func F() string { return "A" }
```

</td></tr>
</tbody></table>

如果表达式的类型与所需的类型不完全匹配，请指定类型。

```go
type myError struct{}

func (myError) Error() string { return "error" }

func F() myError { return myError{} }

var _e error = F()
// F returns an object of type myError but we want error.
```

### Prefix Unexported Globals with _

Prefix unexported top-level `var`s and `const`s with `_` to make it clear when
they are used that they are global symbols.

在私有的顶级var和const之前加_前缀，以便在使用它们时清楚地表明它们是全局符号。

例外：私有的err value，应以`err`为前缀

基本原则：顶级变量和常量具有包范围。使用通用名称可以轻松地在其他文件中使用错误的值。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
// foo.go

const (
  defaultPort = 8080
  defaultUser = "user"
)

// bar.go

func Bar() {
  defaultPort := 9090
  ...
  fmt.Println("Default port", defaultPort)

  // We will not see a compile error if the first line of
  // Bar() is deleted.
}
```

</td><td>

```go
// foo.go

const (
  _defaultPort = 8080
  _defaultUser = "user"
)
```

</td></tr>
</tbody></table>

### Embedding in Structs

嵌入式类型（例如互斥体）应位于结构的字段列表的顶部，并且必须有一个空行将嵌入式字段与常规字段分隔开

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
type Client struct {
  version int
  http.Client
}
```

</td><td>

```go
type Client struct {
  http.Client

  version int
}
```

</td></tr>
</tbody></table>

### Use Field Names to initialize Structs

初始化结构时，几乎始终指定字段名称。现在由 [`go vet`] 强制执行.

  [`go vet`]: https://golang.org/cmd/vet/

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
k := User{"John", "Doe", true}
```

</td><td>

```go
k := User{
    FirstName: "John",
    LastName: "Doe",
    Admin: true,
}
```

</td></tr>
</tbody></table>

例外：如果有3个或更少的字段，则可以在测试表中省略字段名称。

```go
tests := []struct{
}{
  op Operation
  want string
}{
  {Add, "add"},
  {Subtract, "subtract"},
}
```

### Local Variable Declarations

如果将变量显式设置为某个值，则应使用短变量声明(`:=`)

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
var s = "foo"
```

</td><td>

```go
s := "foo"
```

</td></tr>
</tbody></table>

但是，在某些情况下，使用var关键字时默认值会更清晰。[Declaring Empty Slices], 例如

  [Declaring Empty Slices]: https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
func f(list []int) {
  filtered := []int{}
  for _, v := range list {
    if v > 10 {
      filtered = append(filtered, v)
    }
  }
}
```

</td><td>

```go
func f(list []int) {
  var filtered []int
  for _, v := range list {
    if v > 10 {
      filtered = append(filtered, v)
    }
  }
}
```

</td></tr>
</tbody></table>

### nil is a valid slice

`nil` 是长度为0的有效slice，这意味着,

- 您不应该明确返回长度为零的切片. 而应该返回 `nil` 替代.

  <table>
  <thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
  <tbody>
  <tr><td>

  ```go
  if x == "" {
    return []int{}
  }
  ```

  </td><td>

  ```go
  if x == "" {
    return nil
  }
  ```

  </td></tr>
  </tbody></table>

- 检查slice是否为空, 通常使用 `len(s) == 0`. 不检查  `nil`.

  <table>
  <thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
  <tbody>
  <tr><td>

  ```go
  func isEmpty(s []string) bool {
    return s == nil
  }
  ```

  </td><td>

  ```go
  func isEmpty(s []string) bool {
    return len(s) == 0
  }
  ```

  </td></tr>
  </tbody></table>

- 零值（用var 声明的slice）可以直接用，无需 `make()`.

  <table>
  <thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
  <tbody>
  <tr><td>

  ```go
  nums := []int{}
  // or, nums := make([]int)

  if add1 {
    nums = append(nums, 1)
  }

  if add2 {
    nums = append(nums, 2)
  }
  ```

  </td><td>

  ```go
  var nums []int

  if add1 {
    nums = append(nums, 1)
  }

  if add2 {
    nums = append(nums, 2)
  }
  ```

  </td></tr>
  </tbody></table>

### Reduce Scope of Variables

尽可能减少变量的范围。如果与 [Reduce Nesting](#reduce-nesting) 范围冲突，请勿缩小范围.

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
err := ioutil.WriteFile(name, data, 0644)
if err != nil {
 return err
}
```

</td><td>

```go
if err := ioutil.WriteFile(name, data, 0644); err != nil {
 return err
}
```

</td></tr>
</tbody></table>


如果需要if之外的函数调用结果，那么不应尝试缩小范围

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
if data, err := ioutil.ReadFile(name); err == nil {
  err = cfg.Decode(data)
  if err != nil {
    return err
  }

  fmt.Println(cfg)
  return nil
} else {
  return err
}
```

</td><td>

```go
data, err := ioutil.ReadFile(name)
if err != nil {
   return err
}

if err := cfg.Decode(data); err != nil {
  return err
}

fmt.Println(cfg)
return nil
```

</td></tr>
</tbody></table>

### Avoid Naked Parameters


函数调用中的裸参数可能会损害可读性。当参数名称的含义不明显时，添加 C-style 注释(`/* ... */`)作为参数名称。

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
// func printInfo(name string, isLocal, done bool)

printInfo("foo", true, true)
```

</td><td>

```go
// func printInfo(name string, isLocal, done bool)

printInfo("foo", true /* isLocal */, true /* done */)
```

</td></tr>
</tbody></table>

Better yet, replace naked `bool` types with custom types for more readable and
type-safe code. This allows more than just two states (true/false) for that
parameter in the future.

更好的是，将裸布尔类型替换为自定义类型，以获取更具可读性和类型安全的代码。将来，该参数不仅允许两个状态（true/false）。

```go
type Region int

const (
  UnknownRegion Region = iota
  Local
)

type Status int

const (
  StatusReady = iota + 1
  StatusDone
  // Maybe we will have a StatusInProgress in the future.
)

func printInfo(name string, region Region, status Status)
```

### Use Raw String Literals to Avoid Escaping

Go 支持 [raw string literals](https://golang.org/ref/spec#raw_string_lit)，可以跨越多行并包含引号。使用这些字符串可以避免更难阅读的手工转义的字符串

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
wantError := "unknown name:\"test\""
```

</td><td>

```go
wantError := `unknown error:"test"`
```

</td></tr>
</tbody></table>

### Initializing Struct References

初始化结构体时，使用 `&T{}` 代替 `new(T)` ，以使其与结构初始化一致

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
sval := T{Name: "foo"}

// inconsistent
sptr := new(T)
sptr.Name = "bar"
```

</td><td>

```go
sval := T{Name: "foo"}

sptr := &T{Name: "bar"}
```

</td></tr>
</tbody></table>

### Format Strings outside Printf

如果您在字符串文字之外声明了`Printf`-style 函数的格式字符串，请使其为`const`值

这有助于`go vet`对格式字符串进行静态分析

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
msg := "unexpected values %v, %v\n"
fmt.Printf(msg, 1, 2)
```

</td><td>

```go
const msg = "unexpected values %v, %v\n"
fmt.Printf(msg, 1, 2)
```

</td></tr>
</tbody></table>

### Naming Printf-style Functions

当声明`Printf`-style的函数时，请确保`go vet`可以检测到它并检查格式字符串。

这意味着应尽可能使用预定义的`Printf`-style 函数名称。`go vet`会默认检查这些。参考Printf family]

  [Printf family]: https://golang.org/cmd/vet/#hdr-Printf_family

如果不能使用预定义名称，请以f：`Wrapf`（而不是`Wrap`）结尾选择的名称。可以要求`go vet`检查特定的`Printf`-style名称，但名称必须以f结尾。

```shell
$ go vet -printfuncs=wrapf,statusf
```

参考 [go vet: Printf family check].

  [go vet: Printf family check]: https://kuzminva.wordpress.com/2017/11/07/go-vet-printf-family-check/

## Patterns

### Test Tables

在核心测试逻辑重复时，将表驱动测试与[subtests]一起使用，以避免重复代码

  [subtests]: https://blog.golang.org/subtests

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
// func TestSplitHostPort(t *testing.T)

host, port, err := net.SplitHostPort("192.0.2.0:8000")
require.NoError(t, err)
assert.Equal(t, "192.0.2.0", host)
assert.Equal(t, "8000", port)

host, port, err = net.SplitHostPort("192.0.2.0:http")
require.NoError(t, err)
assert.Equal(t, "192.0.2.0", host)
assert.Equal(t, "http", port)

host, port, err = net.SplitHostPort(":8000")
require.NoError(t, err)
assert.Equal(t, "", host)
assert.Equal(t, "8000", port)

host, port, err = net.SplitHostPort("1:8")
require.NoError(t, err)
assert.Equal(t, "1", host)
assert.Equal(t, "8", port)
```

</td><td>

```go
// func TestSplitHostPort(t *testing.T)

tests := []struct{
  give     string
  wantHost string
  wantPort string
}{
  {
    give:     "192.0.2.0:8000",
    wantHost: "192.0.2.0",
    wantPort: "8000",
  },
  {
    give:     "192.0.2.0:http",
    wantHost: "192.0.2.0",
    wantPort: "http",
  },
  {
    give:     ":8000",
    wantHost: "",
    wantPort: "8000",
  },
  {
    give:     "1:8",
    wantHost: "1",
    wantPort: "8",
  },
}

for _, tt := range tests {
  t.Run(tt.give, func(t *testing.T) {
    host, port, err := net.SplitHostPort(tt.give)
    require.NoError(t, err)
    assert.Equal(t, tt.wantHost, host)
    assert.Equal(t, tt.wantPort, port)
  })
}
```

</td></tr>
</tbody></table>


测试表使得 向错误消息添加上下文，减少重复的逻辑以及添加新的测试用例变得更加容易。

我们遵循惯例，将结构体切片称为`tests`，每一个测试用例称为 `tt`. 此外，我们鼓励使用 `give` 和 `want` 前缀来说明每个测试用例的输入和输出值.

```go
tests := []struct{
  give     string
  wantHost string
  wantPort string
}{
  // ...
}

for _, tt := range tests {
  // ...
}
```

### Functional Options

Functional options 是一种模式，你可以声明一个不透明的`Option`类型，该类型在某些内部结构中记录信息。然后可以接受一些options，并根据option内部结构上的记录的生效。

将此模式用于你需要扩展的构造函数或者其他公共APIs的可选参数，尤其是这些函数已经有3个甚至更多参数的情况下

<table>
<thead><tr><th>差劲👎</th><th>优秀👍</th></tr></thead>
<tbody>
<tr><td>

```go
// package db

func Connect(
  addr string,
  timeout time.Duration,
  caching bool,
) (*Connection, error) {
  // ...
}

// Timeout and caching must always be provided,
// even if the user wants to use the default.

db.Connect(addr, db.DefaultTimeout, db.DefaultCaching)
db.Connect(addr, newTimeout, db.DefaultCaching)
db.Connect(addr, db.DefaultTimeout, false /* caching */)
db.Connect(addr, newTimeout, false /* caching */)
```

</td><td>

```go
type options struct {
  timeout time.Duration
  caching bool
}

// Option overrides behavior of Connect.
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

// Connect creates a connection.
func Connect(
  addr string,
  opts ...Option,
) (*Connection, error) {
  options := options{
    timeout: defaultTimeout,
    caching: defaultCaching,
  }

  for _, o := range opts {
    o.apply(&options)
  }

  // ...
}

// Options must be provided only if needed.

db.Connect(addr)
db.Connect(addr, db.WithTimeout(newTimeout))
db.Connect(addr, db.WithCaching(false))
db.Connect(
  addr,
  db.WithCaching(false),
  db.WithTimeout(newTimeout),
)
```

</td></tr>
</tbody></table>

参考，

- [Self-referential functions and the design of options]
- [Functional options for friendly APIs]

  [Self-referential functions and the design of options]: https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
  [Functional options for friendly APIs]: https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

<!-- TODO: replace this with parameter structs and functional options, when to
use one vs other -->