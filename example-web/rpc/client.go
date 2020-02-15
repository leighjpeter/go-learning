package main

import (
	"fmt"
	"github.com/leighjpeter/go-learning/example-web/rpc/utils"
	"log"
	"net/rpc"
)

func main() {
	var serverAddress = "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress+":8080")
	if err != nil {
		log.Fatal("建立连接失败", err)
	}
	args := &utils.Args{10, 10}
	var reply int

	// Call是同步调用
	err = client.Call("MathService.Mutiply", args, &reply)
	if err != nil {
		log.Fatal("调用 MathService.Mutiply err:", err)
	}
	fmt.Printf("%d * %d=%d\n", args.A, args.B, reply)
	// Go是异步调用
	divideCall := client.Go("MathService.Divide", args, &reply, nil)
	for {
		select {
		case <-divideCall.Done:
			fmt.Printf("%d / %d=%d\n", args.A, args.B, reply)
			return
		}
	}
}
