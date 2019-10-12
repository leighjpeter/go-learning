### go内存对齐


```
type Part1 struct {
    a bool
    b int32
    c int8
    d int64
    e byte
}

type Part2 struct {
    e byte
    c int8
    a bool
    b int32
    d int64
}
```



Part1

| 成员变量   | 类型  | 偏移量 | 自身占用 |
| ---------- | ----- | ------ | -------- |
| a          | Bool  | 0      | 1        |
| 字节对齐   | 无    | 1      | 3        |
| b          | Int32 | 4      | 4        |
| c          | Int8  | 8      | 1        |
| 字节对齐   | 无    | 9      | 7        |
| d          | Int64 | 16     | 8        |
| e          | Byte  | 24     | 1        |
| 字节对齐   | 无    | 25     | 7        |
| 总占用大小 |       |        | 32       |

Part1内存布局：axxx|bbbb|cxxx|xxxx|dddddddd|exxx|xxxx

Part2

| 成员变量   | 类型  | 偏移量 | 自身占用 |
| ---------- | ----- | ------ | -------- |
| e          | byte  | 0      | 1        |
| c          | int8  | 1      | 1        |
| a          | bool  | 2      | 1        |
| 字节占用   | 无    | 3      | 1        |
| b          | int32 | 4      | 4        |
| d          | Int64 | 8      | 8        |
| 总占用大小 | -     |        | 16       |

Part2 内存布局：ecax|bbbb|dddd|dddd



**结论：Part1存在许多Padding，浪费。**

**调整结构体内成员变量的字段顺序就能达到缩小结构体占用大小**





### for-loop与json.Unmarshal性能分析

for<for-range<json-iterator<encoding/json

+ for-range 在循环开始前对范围表达式进行求值，

  ```go
  RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .	
  ```

+ 始终是用值拷贝的方式来生成循环变量

+ 官方的 `encoding/json` 标准库，是通过大量反射来实现的。那么 “慢”，也是必然的

+ ` json-iterator/go` 利用`modern-go/reflect2`减少运行时调度开销。类型为 struct 时，只需要反射一次 Name 和 Type，会缓存 struct Encoder 和 Decoder

  ```go
  import "github.com/json-iterator/go"
  var json=jsoniter.ConfigCompatibleWithStandardLibrary
  json.Unmarshal(data,&dst)
  dst,err := json.Marshal(&data)
  ```

  



**总结：**

+ 对性能开销有较高要求，用for
+ 中规中矩 用for-range 大对象慎用
+ 量小、占用小、数量可控：选用 `json.Marshal/Unmarshal` 的方案也可以





### 堆栈

+ 堆(Heap):人为手动进行管理，手动申请，分配，释放。一般涉及的内存大小不定，一般存放较大的对象，分配慢，指令动作也多。
+ 栈(Stack):有编译器进行管理，自动申请，分配，释放。一般不会太大。我们常见的函数参数（不同平台允许存放的数量不同），局部变量等等都会存放在栈上



#### 逃逸分析

+ 是否有在其他地方（非局部）被引用。只要**有可能**被引用了，那么它**一定**分配到堆上。否则分配到栈上

+ 即使没有被外部引用，但对象过大，无法存放在栈区上。依然有可能分配到堆上

```go
案例一： 
type User struct {
	ID     int64
	Name   string
	Avatar string
}

func GetUserInfo() *User {
	return &User{ID: 13746731, Name: "EDDYCJY", Avatar: "https://avatars0.githubusercontent.com/u/13746731"}
}

func main() {
	_ = GetUserInfo()
}

&User literal escapes to heap

原因分析：

很核心的一点就是它有没有被作用域之外所引用。

这是因为 GetUserInfo() 返回的是指针对象，引用被返回到了方法之外了。因此编译器会把该对象分配到堆上，而不是栈上。否则方法结束之后，局部变量就被回收了，岂不是翻车。所以最终分配到堆上是理所当然的

```


  ```
案例二： 
func main() {
    str := new(string)
    *str = "EDDYCJY"
}
  
main new(string) does not escape
  
案例三：
  
func main() {
  str := new(string)
  *str = "EDDYCJY"

	fmt.Println(str)
}

./main.go:9:13: str escapes to heap
./main.go:6:12: new(string) escapes to heap
./main.go:9:13: main ... argument does not escape

原因分析：

通过对其分析，可得知当形参为 interface 类型时，在编译阶段编译器无法确定其具体的类型。因此会产生逃逸，最终分配到堆上

如果你有兴趣追源码的话，可以看下内部的 reflect.TypeOf(arg).Kind() 语句，其会造成堆逃逸，而表象就是 interface 类型会导致该对象分配到堆上
  ```



**总结：**

+ 通过 `go build -gcflags '-m -l'` 就可以看到逃逸分析的过程和结果
+ 静态分配到栈上，性能一定比动态分配到堆上好
+ 到处都用指针传递并不一定是最好的，要用对

























