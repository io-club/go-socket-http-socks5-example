package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Conn struct {
}

func NewConn() *Conn {
	rand.Seed(time.Now().UnixMicro())
	return &Conn{}
}

func (c *Conn) Read(buf *[]byte) int {
	s := ""
	if rand.Intn(100)%2 == 0 {
		s = "A"
	} else {
		s = "B\n"
	}
	*buf = append(*buf, []byte(s)...)
	return len(s)
}
func (Conn) write() int { return 1 }

type BufferReader struct {
	conn *Conn
	buf  []byte
	s    []byte
}

func NewReader(conn *Conn) *BufferReader {
	return &BufferReader{conn: conn, buf: make([]byte, 0), s: make([]byte, 0)}
}

func (self *BufferReader) ReadLine() string {
	for {
		self.conn.Read(&self.buf)
		self.s = append(self.s, self.buf...)
		if self.s[len(self.s)-1] == '\n' {
			ret := string(self.s)
			self.s = make([]byte, 0)
			return ret
		}
		self.buf = make([]byte, 0)
	}
}

func main() {
	conn := NewConn()
	reader := NewReader(conn)
	for {
		msg := reader.ReadLine()
		fmt.Print(msg)
		time.Sleep(1 * time.Second)
	}
}
