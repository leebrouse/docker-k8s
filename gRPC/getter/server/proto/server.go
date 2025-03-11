package main

import (
	"context"
	"fmt"
	"gRPC/server/proto/greeter"
	"net"

	"google.golang.org/grpc"
)

type Hello struct{}

func (this Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req)
	return &greeter.HelloRes{
		Message: "Hello" + req.Name,
	}, nil
}

func main() {
	fmt.Println("Start gRPC.....")
	//init a grpc object
	grpcServer := grpc.NewServer()
	//regist service
	greeter.RegisterGreeterServer(grpcServer, &Hello{})
	//set listener
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	grpcServer.Serve(listener)
}
