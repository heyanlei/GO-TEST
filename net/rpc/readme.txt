JSON-RPC是一个轻量级的远程调用协议，简单易用。

请求数据体:

{
    "method": "getName",
    "params": ["1"],
    "id": 1
}
method: 远端的方法名

params: 远程方法接收的参数列表

id: 本次请求的标识码，远程返回时数据的标识码应与本次请求的标识码相同

返回数据体:

{
    "result": {"id": 1, "name": "name1"},
    "error": null,
    "id": 1
}
result: 远程方法返回值

error: 错误信息

id: 调用时所传来的id

net/rpc包实现了最基本的rpc调用，它默认通过HTTP协议传输gob数据来实现远程调用。

服务端实现了一个HTTP server,接收客户端的请求，在收到调用请求后，会反序列化客户端传来的gob数据，获取要调用的方法名，并通过反射来调用我们自己实现的处理方法，这个处理方法传入固定的两个参数，并返回一个error对象，参数分别为客户端的请求内容以及要返回给客户端的数据体的指针。

net/rpc/jsonrpc

net/rpc/jsonrpc包实现了JSON-RPC协议，即实现了net/rpc包的ClientCodec接口与ServerCodec，增加了对json数据的序列化与反序列化。


Go JSON-RPC远程调用
客户端与服务端双方传输数据，其中数据结构必须得让双方都能处理。

首先定义rpc所传输的数据的结构，client端与server端都得用到。

// 需要传输的对象
type RpcObj struct {
    Id   int `json:"id"` // struct标签， 如果指定，jsonrpc包会在序列化json时，将该聚合字段命名为指定的字符串
    Name string `json:"name"`
}

// 需要传输的对象
type ReplyObj struct {
    Ok  bool `json:"ok"`
    Id  int `json:"id"`
    Msg string `json:"msg"`
}
RpcObj 为传输的数据

ReplyObj 为服务端返回的数据

这两个结构体均可以在client和server端双向传递