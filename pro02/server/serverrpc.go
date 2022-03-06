package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

// 算数运算结构体
type Arith struct{}

//算术运算响应结构体
type ArithRequest struct {
	A int
	B int
}

// 算术运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

// 乘法运算方法
func (arith *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	res.Quo = req.A / req.B
	res.Rem = req.A / req.B
	return nil
}

// 商运算方法
func (arith *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("division by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	err := rpc.Register(new(Arith)) // 注册rpc服务
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("start rpc server")
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			_, err = fmt.Fprintf(os.Stdout, "%s", "new client in comming\n")
			if err != nil {
				log.Fatal(err)
			}
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
