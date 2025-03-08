package main

import (
	"fmt"
	"net/rpc"
)

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

func main() {
	//建立服务器
	client, err := rpc.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
	}

	defer client.Close()

	//调用远程函数
	var result AddGoodRes
	err = client.Call("good.AddGood", AddGoodReq{
		ID:      1,
		Title:   "Luxury Goods",
		Price:   1000,
		Content: "这是奢侈品",
	}, &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
