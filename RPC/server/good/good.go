package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义类结构
type Goods struct{}
type AddGoodReq struct {
	ID      int
	Title   string
	Price   int
	Content string
}

type AddGoodRes struct {
	Success bool
	Message string
}

// 绑定类方法
func (g *Goods) AddGood(req AddGoodReq, res *AddGoodRes) error {
	fmt.Println(req)

	//TODO:
	*res = AddGoodRes{
		Success: true,
		Message: "Add Success",
	}
	return nil
}

func main() {
	//register rpc
	err := rpc.RegisterName("good", new(Goods))
	if err != nil {
		fmt.Println(err)
	}

	//listen
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
	}

	defer listener.Close()

	for {
		fmt.Println("开始建立连接")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//绑定服务
		rpc.ServeConn(conn)
	}

}
