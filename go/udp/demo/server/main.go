package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 23333,
	})
	if err != nil {
		fmt.Printf("UDP listen failed: %v\n", err)
	}

	defer listen.Close()

	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("udp read failed: %v\n", err)
			continue
		}
		fmt.Printf("received data: %v, addr: %v, count: %v\n ", string(buf[:n]), addr, n)
		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Println("udp write failed, err", err)
			continue
		}
	}
}
