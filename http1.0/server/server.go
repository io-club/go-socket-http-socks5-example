package main

import (
	"bufio"
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
	reader := bufio.NewReader(conn)
	// 读取 method, uri, http_version
	first_line, _ := reader.ReadString('\n')
	fmt.Print(first_line)

	// 读取 header
	var headers []string
	for {
		header, _ := reader.ReadString('\n')
		fmt.Print(header)
		if header == "\r\n" {
			break
		}
		headers = append(headers, header)
	}

	// 读取 body
	// body, _ := reader.ReadString('\n')

	// GET 一般没有 BODY
	method, uri := func(s string, sep string) (string, string) {
		x := strings.Split(s, sep)
		return x[0], x[1]
	}(first_line, " ")

	// body, _ := reader.ReadString('\n')
	// fmt.Print("body: \n" + body)
	var ret_body string
	// 根据用户访问的方法和路径做出不同操作
	switch uri {
	case "/", "/index.html":
		ret_body = "访问了 /index"
	case "/user/edit":
		switch method {
		case "GET":
			ret_body = "该接口不能用 GET 方法访问"
		case "POST":
			ret_body = "添加用户成功"
		}
	case "/render":
		ret_body = "<html><body>Hello World</body></html>"
	}
	res := "HTTP/1.0 200 OK\r\n"
	res += "Content-Type: text/plain\r\n"
	res += fmt.Sprintf("Content-Length: %v\r\n", len(ret_body))
	res += "\r\n"
	res += ret_body
	res += "\r\n"
	fmt.Println(res)
	// 返回给客户端
	_, _ = conn.Write([]byte(res))
}
