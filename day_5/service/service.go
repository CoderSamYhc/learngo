package main

import (
	"context"
	"fmt"
	"github.com/CoderSamYhc/learngo/day_5/protoes"
	"google.golang.org/grpc"
	"net"
)

type service struct {

}

func (s *service) SayHi(ctx context.Context, in *protoes.HelloRequest) (*protoes.HelloReplay, error) {
	return &protoes.HelloReplay{
		Message: "Hi" + in.Name,
	}, nil
}

func (s *service) GetMsg(ctx context.Context, in *protoes.HelloRequest) (*protoes.HelloMessage, error) {
	return &protoes.HelloMessage{
		Msg: "Service is running...",
	}, nil
}

func main() {

	// 监听网络
	ln, err := net.Listen("tcp", "127.0.0.1:9981")

	if err != nil {
		fmt.Println(err)
		return
	}

	// grpc 句柄
	srv := grpc.NewServer()

	protoes.RegisterHelloServerServer(srv, &service{})

	// 监听服务
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println(err)
		return
	}


}