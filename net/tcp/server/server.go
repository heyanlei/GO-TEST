package main

//net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket。
//虽然本包提供了对网络原语的访问，大部分使用者只需要Dial、Listen和Accept函数提供的基本接口；以及相关的Conn和Listener接口。
import (
	"fmt"
	"log"
	"net"
)

/*
type Conn interface {
    // Read从连接中读取数据
    // Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    Read(b []byte) (n int, err error)
    // Write从连接中写入数据
    // Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    Write(b []byte) (n int, err error)
    // Close方法关闭该连接
    // 并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
    Close() error
    // 返回本地网络地址
    LocalAddr() Addr
    // 返回远端网络地址
    RemoteAddr() Addr
    // 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
    // deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
    // deadline对之后的所有I/O操作都起效，而不仅仅是下一次的读或写操作
    // 参数t为零值表示不设置期限
    SetDeadline(t time.Time) error
    // 设定该连接的读操作deadline，参数t为零值表示不设置期限
    SetReadDeadline(t time.Time) error
    // 设定该连接的写操作deadline，参数t为零值表示不设置期限
    // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
    SetWriteDeadline(t time.Time) error
}
Conn接口代表通用的面向流的网络连接。多个线程可能会同时调用同一个Conn的方法。
*/
//Listen函数创建的服务端：
/*
ln, err := net.Listen("tcp", ":8080")
if err != nil {
	// handle error
}
for {
	conn, err := ln.Accept()
	if err != nil {
		// handle error
		continue
	}
	go handleConnection(conn)
}
func handleConnection(c net.Conn){
     defer io.close()
     for{
        io.Copy(c, c)
     }
}
*/
//[]byte和string可以互相转换...string(byteParam)
//[]byte(string)
/*
golang里边 string的概念其实不是以前遇到\0结尾的概念了，他其实就是一块连续的内存，首地址+长度，上面那样赋值，
如果p里边有\0，他不会做处理这个时候，如果再对这个string做其他处理就可能出问题了，
比如strconv.Atoi转成int就有错误，解决办法就是需要自己写一个正规的转换函数：
func byteString(p []byte) string {
        for i := 0; i < len(p); i++ {
                if p[i] == 0 {
                        return string(p[0:i])
                }
        }
        return string(p)
}
*/
func handleWrite(conn net.Conn, content []byte) {

	_, e := conn.Write(content)
	if e != nil {
		fmt.Println("Error to send message because of ", e.Error())
	}

}

func main() {

	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		addr := conn.RemoteAddr()
		fmt.Println(addr.Network(), addr)
		go func(c net.Conn) {
			defer c.Close()
			buf := make([]byte, 1024)
			for {

				reqLen, err := conn.Read(buf)
				if err != nil {
					fmt.Printf("recived error !\n")
					return
				}
				strContent := string(buf[:reqLen])
				fmt.Printf("recived content is %s len is %d\n", strContent, reqLen)
				handleWrite(c, []byte(strContent))
				//清楚buf数据
				for i := range buf {
					buf[i] = 0
				}

			}
			// Echo all incoming data.
			//io.Copy(c, c)
			// Shut down the connection.

		}(conn)
	}
}
