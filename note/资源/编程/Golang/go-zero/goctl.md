## API

### goctl代码生成

可以在doc/genCode.md查看

```powershell
# 生成多个types类型的话需要关闭这个实验性功能
goctl env -w GOCTL_EXPERIMENTAL=off

# 代码生成，会自动处理引入的子api文件
goctl api go -api background.api -dir .  --home D:/zhangyuxin/template/1.8.3
goctl api go -api login.api -dir .  --home ../../../deploy/goctl
goctl api go -api main.api -dir .  --home ../../deploy/goctl
```

### 参数校验

| 接收规则 | 说明                                       | 示例                            |
| -------- | ------------------------------------------ | ------------------------------- |
| optional | 当前字段是可选参数，允许为零值(zero value) | `json:"foo,optional"`           |
| options  | 当前参数仅可接收的枚举值                   | `json:"gender,options=foo|bar"` |
| default  | 当前参数默认值                             | `json:"gender,default=male"`    |
| range    | 当前参数数值有效范围，支持开/闭区间        | `json:"age,range=[0:120]"`      |

#### **`optional`** **与** **`omitempty`** **的区别**

两者都写在 json/form tag 的 options 里，但作用域完全不同，不可互相替代

##### `optional` —— goctl 专用，控制 Swagger 与请求校验

- **Swagger 生成**：标了 `optional` 的字段不进 `required` 数组，Apifox/Swagger UI 显示为可选；未标的字段进 `required` 数组，显示为必填
- **请求校验**：goctl 生成的 logic 不会对标了 `optional` 的字段做必填校验
- **判定逻辑**：goctl 遍历 tag options，找到 `optional` 关键字即判定为可选

##### `omitempty` —— Go 标准库指令，控制 JSON 序列化

- **JSON 序列化**：字段为零值（0、""、nil、false）时，响应 JSON 中**不输出**该字段
- **Swagger 生成**：**goctl 完全不识别**，字段仍按未标 `optional` 处理，即判定为 required
- **请求校验**：无影响

## RPC

### goctl代码生成

```powershell
# 主proto文件代码生成
goctl rpc protoc user.proto  --go_out=.   --go-grpc_out=.   --zrpc_out=.   --proto_path=. -m  --home ../../../deploy/goctl

# 主proto文件中引入的proto代码生成
protoc.exe -I ./ --go_out=. --go-grpc_out=.   ./user_model.proto

```

当`proto`文件里面定义了多个 service的时候，需要额外注意，还要再main函数中手动注册一下

> 每个 proto 文件应该聚焦于一个核心概念/实体。一个文件里定义太多 service 通常是设计需要重新考虑的信号

注册示例

```go
import (
	ebookServer "olive/services/product/rpc/internal/server/ebookservice"
)

var configFile = flag.String("f", "etc/product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// 新服务
		productrpcmodel.RegisterEbookServiceServer(grpcServer, ebookServer.NewEbookServiceServer(ctx))
	})
	defer s.Stop()

	fmt.Printf("Starting product RPC server at %s...\n", c.ListenOn)
	s.Start()
}
```



### proto代码生成错误：File not found错误

- 错误描述

  执行 protoc 报 google/protobuf/timestamp.proto: File not found. 的错提示File not found 

- 解决方式

  重新下载[protoc](https://github.com/protocolbuffers/protobuf/releases)，解压后，将整个include目录复制到protoc执行文件同一级目录下





## model

可以在doc/genCode.md查看

**Model 层**：只做 **“业务循环”** + **“SQL 查询”** + **“结果封装”**，不关心字符串到时间的转换和 Protobuf 类型。

### goctl代码生成

```powershell
# mysql根据远程数据库表生成
goctl model mysql datasource --url="olive:xxvkikaDDx01&&99@tcp(127.0.0.1:3306)/ebming" --table="user_invitation_record" --dir="." -c --home ../../../../deploy/goctl/1.8.3

# mysql根据ddl生成
goctl model mysql ddl --src h5.sql --dir . -c --home D:/zhangyuxin/template/1.8.3

# pgsql根据远程数据库表生成
goctl model pg datasource -url "postgres://labman:tN5bWz2KpL9Fy4@127.0.0.1:15432/sern" -table user -dir ./user --cache --home ../../goctl_template/1.8.3

# pgsql根据远程数据库表生成,带模式（schema）
goctl model pg datasource -url "postgres://admin:zhangyuxin@127.0.0.1:25432/test?sslmode=disable" -s "users" -table "profile" -dir ./go-zero/model/users/ --cache --home ../../deploy/goctl
```

