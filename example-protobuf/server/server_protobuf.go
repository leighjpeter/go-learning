package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	stProto "github.com/leighjpeter/go-learning/example-protobuf/proto"
	"net"
	"os"
)

func readMessage(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096, 4096)
	for {
		cnt, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		stReceive := &stProto.UserInfo{}
		pData := buf[:cnt]
		//protobuf解码
		err = proto.Unmarshal(pData, stReceive)
		if err != nil {
			panic(err)
		}

		fmt.Println("receive", conn.RemoteAddr(), stReceive)
		if stReceive.Message == "stop" {
			os.Exit(1)
		}
	}
}
func main() {
	//
	listener, err := net.Listen("tcp", "localhost:7000")
	if err != nil {
		fmt.Println("listen fail")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("new connect", conn.RemoteAddr())
		go readMessage(conn)
	}
}
