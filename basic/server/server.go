package main

import (
	"fmt"
	"net"
)

func main() {
	// 建立 tcp 服务
	listen, err := net.Listen("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}

	for {
		// 等待客户端建立连接
		conn, _ := listen.Accept()
		// 启动一个单独的 goroutine 去处理连接
		process(conn)
		// 处理完关闭连接
		defer conn.Close()
	}
}

func process(conn net.Conn) {
	fmt.Printf("%v\n", conn.RemoteAddr())
	// 针对当前连接做发送和接受操作

	// 准备一个缓冲区，或者说容器，设置大小为 10 个字节
	// 服务器并不知道客户端要发送多少消息，只能说不断尝试接受收
	var buf []byte = make([]byte, 10)
	for {
		n, _ := conn.Read(buf)

		recv := string(buf[:n])
		fmt.Printf("收到的数据：%v\n", recv)

		// 将接受到的数据返回给客户端
		_, _ = conn.Write([]byte("ok"))

		// 如果有一次没收到数据，就说明客户端发送完了
		if n == 0 {
			break
		}
	}
}
