### URL设计规范

URL为统一资源定位器 ,接口属于服务端资源，通常一个完整的URL组成由以下几个部分构成：

```
URI = scheme "://" host  ":"  port "/" path [ "?" query ][ "#" fragment ]
```

- **scheme**: 指底层用的协议，如http、https、ftp
- **host**: 服务器的IP地址或者域名
- **port**: 端口，http默认为80端口
- **path**: 访问资源的路径，就是各种web 框架中定义的route路由
- **query**: 查询字符串，为发送给服务器的参数，在这里更多发送数据分页、排序等参数。
- **fragment**: 锚点，定位到页面的资源



**通常一个RESTful API的path组成**如下：

```undefined
/{version}/{resources}/{resource_id}
```

- **version**：API版本号，有些版本号放置在头信息中也可以，通过控制版本号有利于应用迭代。
- **resources**：资源，RESTful API推荐用小写英文单词的复数形式。
- **resource_id**：资源的id，访问或操作该资源。



有时候可能资源级别较大，其下还可细分很多子资源也可以灵活设计URL的path，例如：

```undefined
/{version}/{resources}/{resource_id}/{subresources}/{subresource_id}
```



当增删改查无法满足业务要求，可以在URL末尾加上action，例如

```
/{version}/{resources}/{resource_id}/action
```

其中action就是对资源的操作。



**RESTful API的URL具体设计的规范**如下：

1. 不用大写字母，所有单词使用英文且小写。
2. 连字符用中杠`"-"`而不用下杠`"_"`
3. 正确使用 `"/"`表示层级关系,URL的层级不要过深，并且越靠前的层级应该相对越稳定
4. 结尾不要包含正斜杠分隔符`"/"`
5. URL中不出现动词，用请求方式表示动作
6. 资源表示用复数不要用单数
7. 不要使用文件扩展名



**[参考网址](https://www.cnblogs.com/bigsai/p/14099154.html)**



---



### HTTP方法

- **GET**

  GET请求会向数据库发索取数据的请求，从而来获取资源，该请求就像数据库的select操作一样，**只是用来查询数据，不会影响资源的内容**。无论进行多少次操作，结果都是一样的

- **POST**

  POST请求向服务器发送数据，但是该请求**会改变数据的内容(新添)**，就像数据库的`insert`操作一样，会创建新的内容，且POST请求的请求参数都是请求体中，其大小是没有限制的

- **PUT**

  PUT请求是向服务器端发送数据的， 与POST请求不同的是，PUT请求**侧重于数据的修改 ,就像数据库中update一样**，而POST请求侧重于数据的增加

- **DELETE**

  用来**删除资源**,和数据库中`delete`相对应

  

### 参数填写位置

- **Params（Query 参数）**：用于 GET 请求或其他将参数**附加在 URL 中**的情况。

- **Body（请求体）**：用于 POST、PUT 等请求方法，将参数**包含在请求体**中。



### 常见的请求文件类型及其适用场景

- **application/json**

 用途：用于传输结构化数据，如对象和数组。

 场景：客户端与服务器之间交换数据时，尤其是在前后端分离的应用中。

- **multipart/form-data**

 用途：用于上传文件或包含文件的表单数据。

 场景：用户上传图片、视频、文档等文件。

- **application/x-www-form-urlencoded**

 用途：用于提交简单的表单数据。

 场景：传统的表单提交，数据量较小且不包含文件。

- **application/octet-stream**

 用途：用于传输二进制数据。

 场景：下载文件或通过 API 提供文件流。