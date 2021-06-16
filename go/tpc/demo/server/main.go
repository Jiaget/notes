package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		// 数据接收
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}
		recv := string(buf[:n])
		fmt.Printf("received data: %v\n", recv)
		// 数据返回
		conn.Write([]byte("request received!"))
	}
}

func main() {
	// 1. 启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:2333")
	if err != nil {
		fmt.Printf("dial failed, err: %v\n", err)
		return
	}
	for {
		// 监听连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		// 处理连接
		go process(conn)
	}
}
