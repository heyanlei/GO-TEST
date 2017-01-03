package main

import (
	"fmt"
	"net"
	"strconv"
)

//Dial函数和服务端建立连接：
/*
在网络network上连接地址address，并返回一个Conn接口。可用的网络类型有：
"tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
其中laddr是本地地址，通常设置为nil。 raddr是一个服务的远程地址, net是一个字符串，可以根据你的需要设置为"tcp4", "tcp6"或"tcp"中的一个。
*/
func handleWrite(conn net.Conn, done chan string) {
	for i := 10; i > 0; i-- {
		_, e := conn.Write([]byte("hello " + strconv.Itoa(i) + "\r\n"))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
	}
	done <- "Sent"
}
func handleRead(conn net.Conn, done chan string) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}
	fmt.Println(string(buf[:reqLen-1]))
	done <- "Read"
}
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		// handle error
		fmt.Println("error")
	}
	defer conn.Close()
	//done := make(chan string)
	//	go handleWrite(conn, done)
	//	go handleRead(conn, done)
	//fmt.Println(<-done)
	//fmt.Println(<-done)
	for {
		_, e := conn.Write([]byte("hello"))
		if e != nil {
			fmt.Println("Error to read message because of ", e)
			return
		}
		buf := make([]byte, 1024)
		reqLen, readErr := conn.Read(buf)
		if readErr != nil {
			fmt.Println("Error to read message because of ", readErr)
			return
		}
		fmt.Println(reqLen)
		fmt.Println(string(buf[:reqLen]))
	}

	/**/
}
