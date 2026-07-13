## 句点引入

一个完整引入声明语句形式的引入名`importname`可以是一个句点(`.`)。 这样的引入称为句点引入。使用被句点引入的包中的导出代码要素时，限定标识符的前缀必须省略。

比如在下面这个例子中，`Println`和`Now`函数调用不需要带任何前缀。

```go
package main

import (
	. "fmt"
	. "time"
)

func main() {
	Println("Current time:", Now())
}
```

> 一般来说，句点引入不推荐使用，因为它们会导致较低的代码可读性



## map的坑

go的map内部是哈希表，元素会因 rehash（扩容）而移动到新的内存位置，如果允许 `&m[key]` 拿到一个指针，map 一旦扩容，这个指针就变成悬空指针（指向已失效的内存）,为了避免这种内存安全问题，Go 直接禁止对 map 值取地址,下面是示例代码：
```go

type User struct {
    Name string
    Age  int
}

func main() {
    m := map[int]User{1: {Name: "张三", Age: 20}}

    // ❌ 编译错误：cannot assign to struct field m[1].Name in map
    m[1].Name = "李四"
    
    // ✅ 正确的替代写法
    temp:=m[1]
	temp.Name = "李四"
	m[1]=temp
}
```

> **可寻址（addressable）**：能通过 `&` 拿到它内存地址的操作数。
>
> **不可寻址（not addressable）**：不能通过 `&` 拿到地址的操作数。

