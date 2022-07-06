package main

import (
	"context"
	"fmt"
	"github.com/CoderSamYhc/learngo/day_5/protoes"
	"google.golang.org/grpc"
)

func main()  {

	// 连接服务器
	conn, err := grpc.Dial("127.0.0.1:9981", grpc.WithInsecure())

	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 获得grpc句柄
	c := protoes.NewHelloServerClient(conn)

	// 远程调用
	r1, err := c.SayHi(context.Background(), &protoes.HelloRequest{Name: "杨虎成"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r1)

	r2, err := c.GetMsg(context.Background(), &protoes.HelloRequest{Name: "lisi"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r2)
}
