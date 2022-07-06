package main

import (
	"fmt"
	"github.com/CoderSamYhc/learngo/day_5/protoes"
	"github.com/golang/protobuf/proto"
)

func main()  {
	helloRq := protoes.HelloRequest{
		Name: *proto.String("lisi"),
		Height: *proto.Int64(182),
	}

	data, err := proto.Marshal(&helloRq)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)

	var list protoes.HelloRequest
	err = proto.Unmarshal(data, &list)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(list.GetName())
	fmt.Println(list.GetHeight())
}
