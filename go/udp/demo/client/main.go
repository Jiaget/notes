package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 23333,
	})

	if err != nil {
		fmt.Printf("dial udp err: %v\n", err)
	}

	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		_, err := socket.Write([]byte(input))
		if err != nil {
			fmt.Printf("write into udp err: %v\n", err)
		}
		// 接收响应
		var buf [1024]byte
		n, addr, err := socket.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("receive from udp failed, err:%v\n", err)
			return
		}
		fmt.Printf("read from %v, messge %v\n", addr, string(buf[:n]))

	}
}
