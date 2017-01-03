package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/*func ResolveUDPAddr(net, addr string) (*UDPAddr, os.Error)
func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err os.Error)
func ListenUDP(net string, laddr *UDPAddr) (c *UDPConn, err os.Error)
func (c *UDPConn) Read(b []byte) (int, error)
func (c *UDPConn) Write(b []byte) (int, error)*/
func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	_, err = conn.Write([]byte(""))
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}
	data := make([]byte, 4)
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err)
		os.Exit(1)
	}
	t := binary.BigEndian.Uint32(data)
	fmt.Println(time.Unix(int64(t), 0).String())
	os.Exit(0)
}
