package main

import (
	"errors"
	"github.com/leighjpeter/go-learning/example-web/rpc/utils"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type MathService struct {
}

func (m *MathService) Mutiply(args *utils.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (m *MathService) Divide(args *utils.Args, reply *int) error {
	if args.B == 0 {
		return errors.New("除数不能为0")
	}
	*reply = args.A / args.B
	return nil
}

func main() {
	math := new(MathService)
	rpc.Register(math)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("启动服务监听失败", err)
	}
	defer listener.Close()
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("启动失败：", err)
	}
}
