package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//建立服务器
	client, err := rpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println(err)
	}

	defer client.Close()

	//调用远程函数
	var result string
	err = client.Call("hello.HelloWorld", "client", &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
