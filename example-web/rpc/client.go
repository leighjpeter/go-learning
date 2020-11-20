package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/leighjpeter/go-learning/example-web/rpc/utils"
)

// func main() {
// 	var serverAddress = "localhost"
// 	client, err := rpc.DialHTTP("tcp", serverAddress+":8080")
// 	if err != nil {
// 		log.Fatal("建立连接失败", err)
// 	}
// 	args := &utils.Args{10, 10}
// 	var reply int

// 	// Call是同步调用
// 	err = client.Call("MathService.Mutiply", args, &reply)
// 	if err != nil {
// 		log.Fatal("调用 MathService.Mutiply err:", err)
// 	}
// 	fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)
// 	// Go是异步调用
// 	divideCall := client.Go("MathService.Divide", args, &reply, nil)
// 	for {
// 		select {
// 		case <-divideCall.Done:
// 			fmt.Printf("%d / %d = %d\n", args.A, args.B, reply)
// 			return
// 		}
// 	}
// }

type MathServiceClient struct {
	Client *rpc.Client
}

// 判断是否实现了MathServiceInterface
var _ utils.MathServiceInterface = (*MathServiceClient)(nil)

func DialMathDervice(network, address string) (*MathServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &MathServiceClient{Client: c}, nil
}
func (m *MathServiceClient) Mutiply(args *utils.Args, reply *int) error {
	return m.Client.Call(utils.MathServiceName+".Mutiply", args, reply)
}

func (m *MathServiceClient) Divide(args *utils.Args, reply *int) error {
	return m.Client.Call(utils.MathServiceName+".Divide", args, reply)
}

func main() {
	client, err := DialMathDervice("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &utils.Args{10, 10}
	var reply int
	err = client.Mutiply(args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)
	err = client.Divide(args, &reply)
	fmt.Printf("%d / %d = %d\n", args.A, args.B, reply)
}
