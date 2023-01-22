package main

import (
	"fmt"
	"net"
	"strings"
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
		conn.Close()
	}
}

func process(conn net.Conn) {
	fmt.Printf("%v\n", conn.RemoteAddr())
	// 针对当前连接做发送和接受操作
	msg := ""
	var buf []byte = make([]byte, 1)
	for {
		n, _ := conn.Read(buf)
		// 如果有一次没收到数据，就说明客户端发送完了
		if n == 0 {
			break
		}

		// 没读到 \n 表示这次对话没结束
		if buf[0] != '\n' {
			msg += string(buf)
			continue
		}

		// 读到 \n 表示这次对话结束，解析拿到的数据
		method, uri := func(s string, sep string) (string, string) {
			x := strings.Split(s, sep)
			return x[0], x[1]
		}(msg, " ")
		fmt.Printf("method: %v uri: %v\n", method, uri)

		// 根据用户访问的方法和路径做出不同操作
		res := ""
		switch uri {
		case "/", "/index.html":
			res += "访问了 /index"
		case "/user/edit":
			switch method {
			case "GET":
				res += "该接口不能用 GET 方法访问"
			case "POST":
				res += "添加用户成功"
			}
		case "/render":
			res = `<html>
			<body>Hello World</body>
		  </html>`
		}

		// 返回给客户端
		_, _ = conn.Write([]byte(res))

		// 清空 msg
		msg = ""
	}
}
