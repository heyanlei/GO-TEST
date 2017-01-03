package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
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

// 客户端以同步的方式向rpc服务器请求
func callRpcBySynchronous() {
	// 连接至服务器
	client, err := net.DialTimeout("tcp", "localhost:8888", 1000*1000*1000*30) // 30秒超时时间
	if err != nil {
		log.Fatal("client\t-", err.Error())
	}

	defer client.Close()

	// 建立rpc通道
	clientRpc := jsonrpc.NewClient(client)

	// 远程服务器返回的对象
	var rpcObj RpcObj
	log.Println("client\t-", "call GetName method")
	// 请求数据，rpcObj对象会被填充
	clientRpc.Call("ServerHandler.GetName", 1, &rpcObj)
	log.Println("client\t-", "recive remote return", rpcObj)

	// 远程返回的对象
	var reply ReplyObj

	// 传给远程服务器的对象参数
	saveObj := RpcObj{2, "对象2"}

	log.Println("client\t-", "call SetName method")
	// 请求数据
	clientRpc.Call("ServerHandler.SaveName", saveObj, &reply)

	log.Println("client\t-", "recive remote return", reply)
}

// 客户端以异步的方式向rpc服务器请求
func callRpcByAsynchronous() {
	// 打开链接
	client, err := net.DialTimeout("tcp", "localhost:8888", 1000*1000*1000*30) // 30秒超时时间
	if err != nil {
		log.Fatal("client\t-", err.Error())
	}
	defer client.Close()

	// 建立rpc通道
	clientRpc := jsonrpc.NewClient(client)

	// 用于阻塞主goroutine
	endChan := make(chan int, 15)

	// 15次请求
	for i := 1; i <= 15; i++ {

		// 传给远程的对象
		saveObj := RpcObj{i, "对象"}

		log.Println("client\t-", "call SetName method")
		// 异步的请求数据
		divCall := clientRpc.Go("ServerHandler.SaveName", saveObj, &ReplyObj{}, nil)

		// 在一个新的goroutine中异步获取远程的返回数据
		go func(num int) {
			reply := <-divCall.Done
			log.Println("client\t-", "recive remote return by Asynchronous", reply.Reply)
			endChan <- num
		}(i)
	}

	// 15个请求全部返回时此方法可以退出了
	for i := 1; i <= 15; i++ {
		_ = <-endChan
	}

}
func main() {
	//callRpcBySynchronous()
	callRpcByAsynchronous()
}
