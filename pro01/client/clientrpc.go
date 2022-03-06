package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatal(err)
	}
	req := ArithRequest{9, 2}
	res := ArithResponse{}
	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	req = ArithRequest{9, 2}
	res = ArithResponse{}
	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
