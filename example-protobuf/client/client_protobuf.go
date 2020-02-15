package main

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	stProto "github.com/leighjpeter/go-learning/example-protobuf/proto"
	"net"
	"os"
)

func main() {
	conIP := "localhost:7000"
	var conn net.Conn
	var err error

	//建立连接
	conn, err = net.Dial("tcp", conIP)
	if err != nil {
		fmt.Println("connect fail")
		return
	}
	fmt.Println("connect", conIP, "success")
	defer conn.Close()
	//发送消息
	cnt := 1
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		stSend := &stProto.UserInfo{
			Message: sender.Text(),
			Length:  *proto.Int(len(sender.Text())),
			Cnt:     *proto.Int(cnt),
		}
		//protobuf编码
		pData, err := proto.Marshal(stSend)
		if err != nil {
			panic(err)
		}

		//发送
		conn.Write(pData)
		if sender.Text() == "stop" {
			return
		}
		cnt++
	}
}
