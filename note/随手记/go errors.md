## `errors.AsType`

> *用泛型把* `errors.As` *的“先声明变量再传指针”简化为“直接返回* `(E, bool)`*”，是* `errors.As` *的语法糖升级*

**使用场景：**

- 与 `errors.As` 完全相同（按类型提取错误），但写起来更简洁
- 由于类型参数 `E` 必须是 `error` 接口的实现类型，**能在编译期避免 `errors.As` 那种“传错指针类型导致运行时 panic”的问题**

```go
// 传统 errors.As：需要先声明变量
var mysqlErr *mysql.MySQLError
if errors.As(err, &mysqlErr) {
    return true, mysqlErr.Number
}

// errors.AsType（Go 1.26+）：无需预先声明变量
if mysqlErr, ok := errors.AsType[*mysql.MySQLError](err); ok {
    return true, mysqlErr.Number
}
```

