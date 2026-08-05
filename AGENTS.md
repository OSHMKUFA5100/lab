# AGENTS.md

## 仓库结构

按语言分层的个人练习仓库，每种语言一个顶层目录，各自独立管理依赖。

| 目录 | 说明 |
|---|---|
| `go/` | Go module `lab`（Go 1.26.2），go.mod 在该目录内 |
| `note/` | 个人笔记（非代码），`note/项目/` 被 gitignore |
| `dart/` | 预留，尚未建立 |

## Go 约定

- **Module 在 `go/` 目录下，不在仓库根**——所有 go 命令需在 `go/` 内执行
- 每个叶子目录 = 一个独立小程序（`package main` + `main.go`）
- 多个小程序可并存，不同目录、各自独立
- 主题目录（basics/concurrency/stdlib/patterns/algorithms/web/testing/project）可随时增删
- 测试文件用 `*_test.go`，运行单个测试：`cd go && go test ./testing/ -run TestSliceAppend`

## 开发命令

```
# Go——所有命令在 go/ 目录下执行
cd go
go run ./<主题>/<程序名>/       # 运行单个程序
go test ./...                  # 运行全部测试
go test ./testing/ -v          # 运行 testing 包测试
go mod tidy                    # 整理依赖
```

## 注意事项

- 仓库无 CI、无 pre-commit hooks、无代码检查——每个程序自行保证可编译即可
- 笔记和练习代码严格分离：程序相关说明写在程序目录的 README.md，系统性知识放 `note/`
- `.gitignore` 已覆盖 Go 二进制、Dart/Flutter 构建产物、IDE 配置、操作系统文件
