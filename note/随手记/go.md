## 内存逃逸

### 什么是逃逸

在函数内部创建的对象或者变量，当函数的生命周期结束后仍然被引用或持有，则该对象或者变量被称之为逃逸对象

### 影响

- 对象分配到堆上，而不是栈上。
- 堆内存由 Runtime 分配，由 GC 管理生命周期。
- 相比栈分配，堆分配开销更大（需要 Runtime 参与）。
- 增加 GC 扫描和回收压力，频繁逃逸可能影响程序性能。
- 对象生命周期变长，不再随着函数返回立即释放，而是等待 GC 回收。
- 如果多个 goroutine 共享该对象，可能产生数据竞争（Data Race），但这是共享对象导致的，不是逃逸本身导致的。
- 逃逸本身不会导致内存泄漏，但如果堆对象长期被引用，GC 无法回收，就会造成 Go 意义上的内存泄漏。

### 常见原因

1. 返回局部变量的指针

   局部变量在函数返回后仍然需要使用，因此必须分配到堆上。

   ```go
   func NewUser() *User {
       u := User{}
       return &u
   }
   ```

2. 局部变量被闭包引用

   闭包可能在函数返回后仍然执行，因此引用的局部变量需要继续存活。

   ```go
   func f() func() int {
       a := 1
       return func() int {
           return a
       }
   }
   ```

3. 被 goroutine 引用

   goroutine 的执行时间可能超过当前函数的生命周期。

   ```go
   func f() {
       a := 1
       go func() {
           fmt.Println(a)
       }()
   }
   ```

4. 作为 interface 参数或返回值（部分情况）

   如果编译器无法证明对象不会逃逸，就会保守地分配到堆上。

   ```go
   func f() interface{} {
       a := 10
       return a
   }
   ```

   > 注意：并不是所有 interface 都会逃逸。

5. 存入生命周期更长的对象

   例如存入全局变量、堆对象或其他长期存在的数据结构。

   ```go
   var users []*User

   func f() {
       u := User{}
       users = append(users, &u)
   }
   ```

   这里 u 必须逃逸，因为 users 在函数返回后仍然存在。

6. 编译器无法证明不会逃逸

   ```go
   func g(x *User)

   func f() {
       u := User{}
       g(&u)
   }
   ```

   如果编译器无法确定 g 是否保存了 u 的地址，就会保守地让 u 逃逸到堆上。

### 逃逸分析命令

```bash
go build -gcflags -m main.go
```

出现`moved to heap`或者`escapes to heap`说明存在逃逸
