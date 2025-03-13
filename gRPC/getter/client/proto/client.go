package main

import (
	"bufio"
	"context"
	"fmt"
	"gRPC/client/proto/greeter"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcClient, err := grpc.Dial("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer grpcClient.Close()

	client := greeter.NewGreeterClient(grpcClient)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println("请输入命令 (输入 'quit' 退出):")
		input := scanner.Text()
		if input == "quit" {
			fmt.Println("退出客户端")
			break
		}

		res, err := client.SayHello(context.Background(), &greeter.HelloReq{
			Name: input,
		})
		if err != nil {
			fmt.Println("调用失败:", err)
			continue
		}

		fmt.Printf("服务器响应: %#v\n", res)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取输入错误:", err)
	}
}
