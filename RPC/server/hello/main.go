package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义类结构
type World struct{}

// 绑定模板方法
func (this *World) HelloWorld(req string, rep *string) error {
	*rep += "Hello "+req 
	return nil
}

func main() {

	//注册rpc
	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println(err)
	}
	//监听

	lisener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println(err)
	}

	defer lisener.Close()

	//建立链接
	for {
		//接受连接
		fmt.Println("开始建立连接")
		con, err := lisener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//绑定服务
		rpc.ServeConn(con)

	}

}
