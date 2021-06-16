package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:2333")
	if err != nil {
		fmt.Printf("dial failed: %v\n", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入，回车结束
		input = strings.TrimSpace(input)
		// 退出设置
		if strings.ToUpper(input) == "Q" || strings.ToUpper(input) == "QUIT" {
			return
		}

		// 发送消息
		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Printf("send failed, err%v\n", err)
			return
		}

		var buf [1024]byte
		n, err := conn.Read((buf[:]))
		if err != nil {
			fmt.Printf("read failed, err:%v", err)
			return
		}
		fmt.Println("received response:", string(buf[:n]))
	}

}
