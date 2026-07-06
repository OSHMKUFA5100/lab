# Go 练习区

本目录是 Go 的学习 / 练手代码，共用一个 module（`go.mod` 在本目录）。

## 约定

- **每个叶子目录 = 一个独立小程序**：`package main` + `main.go`。
- 多个小程序并存互不冲突（不同目录、各自的 `main`）。
- 某个程序依赖太特殊时，可在它自己的目录单独 `go mod init`，脱离本 module。

## 主题目录

| 目录 | 放什么 |
|---|---|
| `basics/` | 语法基础、内置类型（slice / map / chan 行为等） |
| `concurrency/` | 并发：goroutine / channel / select / context |
| `stdlib/` | 标准库练习（net/http、encoding/json、io…） |
| `patterns/` | 设计模式 |
| `algorithms/` | 数据结构与算法 |
| `web/` | net/http、Gin 等框架 |
| `testing/` | 练习 Go 测试机制（`*_test.go`、`go test`、table-driven、testify） |
| `project/` | 完整小项目（多文件） |

> 主题目录可随时增删。写新程序时挑最贴的主题，在里面建一个新子目录放 `main.go` 即可。