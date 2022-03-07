package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"pro03/pb"
	"time"
)

const (
	address = "127.0.0.1:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewArithClient(conn)
	// 1s 的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Multiply(ctx, &pb.ArithRequest{A: 9, B: 2})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Pro)
	r, err = c.Divide(ctx, &pb.ArithRequest{A: 9, B: 2})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Quo, r.Rem)
}
