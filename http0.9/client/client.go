package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 1、与服务端建立连接
	// 这里并没有为客户端制定端口号，也没必要
	conn, _ := net.Dial("tcp", "127.0.0.1:9091")
	// 2、使用 conn 连接进行数据的发送和接收
	msgs := []string{
		"GET /index.html\n",
		"GET /user/edit\n",
		"POST /user/edit\n",
		"POST /render\n",
	}
	for i := 0; i < len(msgs); i++ {
		msg := msgs[i]

		n, _ := conn.Write([]byte(msg))
		fmt.Printf("write: %v\n", n)

		// 从服务端接收回复消息
		var buf [1024]byte
		n_read, _ := conn.Read(buf[:])
		fmt.Printf("收到服务端回复:%v\n", string(buf[:n_read]))
		time.Sleep(1 * time.Second)
	}
}
