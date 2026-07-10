# HTTP

## 上传文件

> **https://github.com/zeromicro/zero-examples/tree/main/http/upload**

### 前端

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Document</title>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="http://localhost:8888/upload"
      method="post"
    >
      <input type="file" name="myFile" />
      <input type="submit" value="upload" />
    </form>
  </body>
</html>
```

> **参数说明**

- `enctype="multipart/form-data"` 用于支持文件的二进制传输

-  `<input type="file">` 让用户挑选本地文件

-  `action` 和 `method`，将数据以 `POST` 方式提交到后端地址



### 后端

```yaml
Name: upload-api
Host: 0.0.0.0
Port: 8888
MaxBytes: 1073741824
Timeout: 30000
```

> **yaml文件需要设置上传的最大字节，并且如果为大文件，需要考虑超时时间设置是否合理**

```go
func (l *UploadLogic) Upload(r *http.Request) (resp *types.Response, err error) {
    err = r.ParseMultipartForm(10 << 27)
    if err != nil {
       fmt.Println("错误发生：", err)
       return nil, err
    }
    file, handler, err := r.FormFile("myFile")
    if err != nil {
       fmt.Println(err)
       return nil, err
    }
    defer file.Close()

    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    tempFile, err := os.Create(path.Join(".", handler.Filename))
    if err != nil {
       fmt.Println(err)
       return nil, err
    }
    defer tempFile.Close()
    io.Copy(tempFile, file)

    return &types.Response{
       OK: 0,
    }, nil
}
```

> **API逻辑说明**

- 需要在handler里面将`*http.Request`传递到逻辑层

- `ParseMultipartForm`用于解析 `multipart/form-data` 类型的请求体

  主要作用是：

​	解析表单字段（包括文件上传字段）

​	将请求体拆分成多个部分（parts），并填充到 `r.MultipartForm` 中

​	内存与磁盘分配：最多在内存中保留 `maxFileSize` 字节，超出的部分会写入临时文件

- `FormFile(key string)` 从已解析的 `r.MultipartForm` 中提取名为 `"myFile"` 的文件部分

- `tempFile, err := os.Create(path.Join(".", handler.Filename))`将文件上传到当前目录





# DTM

## 基础概念

- **AP（Application）**：应用程序，发起全局事务请求。
- **TM（Transaction Manager）**：事务协调者，负责协调分布式事务的执行流程。
- **RM（Resource Manager）**：资源管理者，例如订单服务、库存服务等，执行具体业务与补偿逻辑
- **事务分支**：我们把每个服务管理的全局事务组成部分，称为事务分支，例如前面的转账，分为转出和转入两个事务分支
- **分支操作**：每个事务分支，在SAGA、XA、TCC等事务模式下，会有多个操作，例如转出事务分支，包括正向操作TransOut和TransOutCompensate
- **本地事务**：转出事务分支中的正向操作，通常会开启一个事务，对余额进行扣减，我们将数据操作的这个事务，称为本地事务
- **GID**：全局事务ID，用于标记全局事务，必须唯一。该ID可以采用UUID生成，也可以使用业务上的ID，例如订单ID等

| 模式       | 场景                                     | 典型使用                          |
| ---------- | ---------------------------------------- | --------------------------------- |
| 二阶段消息 | 需要确保消息通知或缓存更新与本地事务同步 | 发短信/推送、消息投递、缓存一致性 |
| Saga       | 流程复杂、长链路、不需强一致性           | 电商下单→扣库存→支付等微服务流程  |
| TCC        | 强一致、高隔离资源场景                   | 金融系统、库存冻结、大额支付      |

## 二阶段消息

是dtm首创的事务模式，用于替换本地事务表和事务消息这两种现有的方案。它能够保证本地事务的提交和全局事务提交是“原子的”，适合解决不需要回滚的分布式事务场景，比如发消息、更新缓存等

- **准备阶段**（Prepare） ，应用先发一个准备信号给 DTM Server，表示“我现在要做本地事务了”。DTM Server 接收到后先记录这个请求，但不触发后续通知。
- 执行本地事务， 应用开始执行本地事务（如写数据库、扣款），并在同一个数据库里，记录这次全局事务的 gid（全局事务 ID）。这保证了后续可以根据 gid 判断本地事务是否已成功。
- **提交阶段**（Submit）， 本地事务成功后，应用再向 DTM 提交提交请求。DTM 执行这个操作后，才真正触发下游服务或消息投递。开始调用其他微服务或消息队列，实现完整的全局动作。

**优点：**

- 保证“本地事务 + 消息”业务要么同时成功，否则都不执行。（强触发保障）
- 对开发者透明，业务侵入小。

**缺点：**

- 需要开发一套 Prepare/Submit 调用，系统需定期监控、回查状态。
- 不适用于需要复杂补偿的场景。

## SAGA

核心思想是将长事务拆分为多个短事务，由Saga事务协调器协调，如果每个短事务都成功提交完成，那么全局事务就正常完成，如果某个步骤失败，则根据相反顺序一次调用补偿操作。例如我们要进行一个类似于银行跨行转账的业务，将A中的30元转给B，根据Saga事务的原理，我们将整个全局事务，切分为以下服务：

- **转出**（TransOut）服务，这里转出将会进行操作A-30
- **转出补偿**（TransOutCompensate）服务，回滚上面的转出操作，即A+30
- **转入**（TransIn）服务，转入将会进行B+30
- **转入补偿**（TransInCompensate）服务，回滚上面的转入操作，即B-30

**优点：**

- 流程自然、开发简单，只要为操作实现补偿逻辑即可。
- 对资源无锁设计，支持高并发、长流程，性能佳。

**缺点：**

- 是最终一致性，存在临时中间状态（如库存已扣但未支付）。
- 无事务隔离，可能出现脏写或并发冲突。
- 补偿逻辑可能复杂易错（幂等、防重复）。

## TCC

TCC 是一种“预留-确认-取消”的事务模式，将事务分为三个阶段，高精度控制点 的场景，比如：支付系统、资金冻结、库存预扣等

- **Try** 阶段：尝试执行，完成所有业务检查（一致性）, 预留必须业务资源（准隔离性）
- **Confirm** 阶段：如果所有分支的Try都成功了，则走到Confirm阶段。Confirm真正执行业务，不作任何业务检查，只使用 Try 阶段预留的业务资源
- **Cancel** 阶段：如果所有分支的Try有一个失败了，则走到Cancel阶段。Cancel释放 Try 阶段预留的业务资源。

**优点：**

- 强一致性：通过 Try 预留资源，Confirm/Cancel 明确执行或回滚，避免中间状态。
- 高并发友好，资源锁粒度小；性能优于传统 XA/2PC。

**缺点：**

- 业务侵入性高，每个参与方需实现三个操作（Try/Confirm/Cancel）。
- 实现复杂：需处理悬挂、空回滚、幂等、补偿重试等问题。

## SAGA实践

- 准备RM和DTM数据表

  [GitHub地址](https://github.com/dtm-labs/dtm/tree/main/sqls)

- DTM配置

  [GitHub地址](https://github.com/dtm-labs/dtm/blob/main/conf.sample.yml)

- docker部署参考配置

```yaml
LogLevel: 'info'
Server:
  HttpPort: 36789	#HTTP port
  GrpcPort: 36790	#Grpc port
Store:
  Driver: "mysql" # 数据库类型
  Host: "mysql"  # 使用容器名
  Port: 3306
  User: "root"
  Password: "root"
  Db: "dtm"
MicroService:
  Driver: "dtm-driver-gozero"	#使用go-zero的注册服务发现驱动
  Target: "etcd://etcd:2379/dtmservice"  # 使用容器名，当前dtm的server直接注册到微服务所在的etcd集群中使用
  EndPoint: "dtm:36790"  # 使用容器名，集群中的微服务可以直接通过etcd获得此地址跟dtm交互
```

- go-zero的yaml参考配置

```yaml
#AP使用etcd的配置
Etcd:
  Hosts:
    - etcd:2379
  Key: order.api
#RM使用etcd的配置
OrderRpc:
  Etcd:
    Hosts:
      - etcd:2379
    Key: order.rpc
#RM使用etcd的配置
StockRpc:
  Etcd:
    Hosts:
      - etcd:2379
    Key: stock.rpc
```

SAGA 事务编排使用步骤

```go
// dtm服务在etcd中的调用地址
	var dtmServer = "etcd://localhost:2379/dtmservice"

// 从配置中通过BuildTarget获得对应RM的etcd调用地址
	orderTarget, err := l.svcCtx.Config.OrderRpcConf.BuildTarget()
	stockTarget, err := l.svcCtx.Config.StockRpcConf.BuildTarget()

// 构造对应的请求结构体，用于后面的本地事务的执行或者回滚补偿
	createOrderReq := &order.CreateReq{UserId: req.UserId, GoodsId: req.GoodsId, Num: req.Num}
	deductReq := &stock.DecuctReq{GoodsId: req.GoodsId, Num: req.Num}

// 使用dtmgrpc注册一个全局事务ID，用于标记全局事务
	gid := dtmgrpc.MustGenGid(dtmServer)

// 开启全局事务，并提交
// dtm需要从dtm服务器调用该方法，所以不走强类型，而是走动态的url: busiServer+"/trans.TransSvc/TransOut"
// "/trans.TransSvc/TransOut"对应的是grpc_pb.go文件中Invoke方法所使用的路径
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(
			orderTarget+"/pb.order/create",
			orderTarget+"/pb.order/createRollback",
			createOrderReq,
		).
		Add(
			stockTarget+"/pb.stock/deduct",
			stockTarget+"/pb.stock/deductRollback",
			deductReq,
		)
	err = saga.Submit()
```

服务实现（包含子事务屏障）

```go
// 从 gRPC 请求上下文中提取分布式事务信息，构建一个用于管理子事务屏障的对象
barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
// 获取一个sqlx的数据库连接
db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()

if err != nil {
// 一般数据库不会错误不需要dtm回滚，就让他一直重试
// 返回为status.Error(codes.Internal, err.Error())，会让dtm一直重试，不会回滚，直到
// 返回为status.Error(codes.Aborted, dtmcli.ResultFailure)，会让dtm立刻调用回滚方法
    return nil, status.Error(codes.Internal, err.Error())
}

// 开启本地事务（这里没写，正常应该在本地事务中再开启子事务屏障进行业务处理
// 将业务逻辑包裹在子事务屏障的保护下执行
if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
    order := new(model.Order)
    order.GoodsId = in.GoodsId
    order.Num = in.Num
    order.UserId = in.UserId

    _, err = l.svcCtx.OrderModel.Insert(tx, order)
    if err != nil {
       return fmt.Errorf("创建订单失败 err : %v , order:%+v \n", err, order)
    }

    return nil
}); err != nil {
// 同上，让dtm一直重试，不回滚
    return nil, status.Error(codes.Internal, err.Error())
}

// 补偿操作跟上面差不多，就是业务逻辑修改
```

