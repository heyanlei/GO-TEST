package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
	"sync"
)

// 需要传输的对象
type RpcObj struct {
	Id   int    `json:"id"` // struct标签， 如果指定，jsonrpc包会在序列化json时，将该聚合字段命名为指定的字符串
	Name string `json:"name"`
}

// 需要传输的对象
type ReplyObj struct {
	Ok  bool   `json:"ok"`
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

// server端的rpc处理器
/*
ServerHandler结构可以不需要什么字段，只需要有符合net/rpcserver端处理器约定的方法即可。

符合约定的方法必须具备两个参数和一个error类型的返回值

第一个参数 为client端调用rpc时交给服务器的数据，可以是指针也可以是实体。net/rpc/jsonrpc的json处理器会将客户端传递的json数据解析为正确的struct对象。

第二个参数 为server端返回给client端的数据,必须为指针类型。net/rpc/jsonrpc的json处理器会将这个对象正确序列化为json字符串，最终返回给client端。

ServerHandler结构需要注册给net/rpc的HTTP处理器，HTTP处理器绑定后，会通过反射得到其暴露的方法，在处理请求时，根据JSON-RPC协议中的method字段动态的调用其指定的方法
*/
type ServerHandler struct{}
type IpaddrStr struct {
	Ip string
}

// server端暴露的rpc方法
func (serverHandler ServerHandler) GetName(id int, returnObj *RpcObj) error {
	log.Println("server\t-", "recive GetName call, id:", id)
	returnObj.Id = id
	returnObj.Name = "名称1"
	return nil
}

// server端暴露的rpc方法
func (serverHandler ServerHandler) SaveName(rpcObj RpcObj, returnObj *ReplyObj) error {
	log.Println("server\t-", "recive SaveName call, RpcObj:", rpcObj)
	returnObj.Ok = true
	returnObj.Id = rpcObj.Id
	returnObj.Msg = "存储成功"
	return nil
}
func main() {
	mutex := sync.Mutex{}
	//	reciveHandler := &IpaddrStr{}
	// 新建Server
	server := rpc.NewServer()

	// 开始监听,使用端口 8888
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("server\t-", "listen error:", err.Error())
	}
	defer listener.Close()

	log.Println("server\t-", "start listion on port 8888")

	// 新建处理器
	serverHandler := &ServerHandler{}

	// 注册处理器
	server.Register(serverHandler)

	// 等待并处理链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go func() {
			// 应用排斥锁
			mutex.Lock()
			// 记录ip地址
			strTmp := strings.Split(conn.RemoteAddr().String(), "]:")
			fmt.Println(strTmp[1])
			server.ServeCodec(jsonrpc.NewServerCodec(conn))
			// 解锁
			mutex.Unlock()
			// 在goroutine中处理请求
			// 绑定rpc的编码器，使用http connection新建一个jsonrpc编码器，并将该编码器绑定给http处理器
		}()

	}
}

/*
在使用net/rpc/jsonrpc时遇到这样一个问题：

有多个client与一个server进行rpc调用，而这些client又处于不同的内网，在server端需要获取client端的公网IP。

按照net/rpc的实现，在服务端处理器的自定义方法中只能获取被反序列化的数据，其他请求相关信息如client的IP只能在主goroutine的net.Listener.Accept中的Conn对象取得。

按源码中的示例，每接收一个TCP请求都会在一个新的goroutine中处理，但是处理器的自定义方法都运行在不同的goroutine中，这些回调的方法没有暴露任何能获取conn的字段、方法。

我是这样解决的，在server端rpc处理器struct中放一个聚合字段，用于存储ip地址的。

处理器被注册与rpc server，全局只有一个，在每次接受到tcp请求后，开启一个goroutine，然后在goroutine内部立即加上排斥锁，
然后再把请求的conn绑定给rpc server处理器，这样，即能保证handler字段的线程安全，又能及时的相应client的请求。
*/
