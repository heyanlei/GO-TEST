"github.com/toolkits/logger"
标准库的RPC
rpc.Register 用于注册RPC服务, 默认的名字是对象的类型名字(这里是Echo).
如果需要指定特殊的名字, 可以用 rpc.RegisterName 进行注册.
被注册对象的类型所有满足以下规则的方法会被导出到RPC服务接口:
func (t *T) MethodName(argType T1, replyType *T2) error
被注册对应至少要有一个方法满足这个特征, 否则可能会注册失败.
然后 rpc.HandleHTTP 用于指定 RPC 的传输协议, 这里是采用 http 协议作为RPC调用的载体. 
用户也可以用rpc.ServeConn接口, 定制自己的传输协议.
//server
type Echo int

func (t *Echo) Hi(args string, reply *string) error {
    *reply = "echo:" + args
    return nil
}

func main() {
    rpc.Register(new(Echo))
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    http.Serve(l, nil)
}
//client
func main() {
    client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    var args = "hello rpc"
    var reply string
    err = client.Call("Echo.Hi", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}
基于 JSON 的 RPC 调用
Go的标准库还提供了一个"net/rpc/jsonrpc"包, 用于提供基于JSON编码的RPC支持.
服务器部分只需要用rpc.ServeCodec指定json编码协议就可以了:
//server
func main() {
    lis, err := net.Listen("tcp", ":1234")
    if err != nil {
        return err
    }
    defer lis.Close()

    srv := rpc.NewServer()
    if err := srv.RegisterName("Echo", new(Echo)); err != nil {
        return err
    }

    for {
        conn, err := lis.Accept()
        if err != nil {
            log.Fatalf("lis.Accept(): %v\n", err)
        }
        go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
    }
}
//client
客户端部分值需要用 jsonrpc.Dial 代替 rpc.Dial 就可以了:
