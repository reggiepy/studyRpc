package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"pro03/pb"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedArithServer
} // 服务对象

func (s *server) Multiply(ctx context.Context, req *pb.ArithRequest) (*pb.ArithReply, error) {
	arithReply := pb.ArithReply{}
	arithReply.Pro = req.A * req.B
	return &arithReply, nil
}

func (s *server) Divide(ctx context.Context, req *pb.ArithRequest) (*pb.ArithReply, error) {
	arithReply := pb.ArithReply{}
	if req.B == 0 {
		return &arithReply, fmt.Errorf("division by zero")
	}
	arithReply.Quo = req.A / req.B
	arithReply.Rem = req.A % req.B
	return &arithReply, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	// 注册反射服务，这个服务时cli使用的，跟服务本身没有关系
	pb.RegisterArithServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("start server Fatal: %v", err)
	}
}
