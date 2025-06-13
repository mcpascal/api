package main

import (
	"api/internal/proto/hello"
	"context"
	"fmt"
	"io"
	"net"

	"google.golang.org/grpc"
)

var _ hello.GreeterServer = &GreeterServer{}

type GreeterServer struct {
	hello.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHi(ctx context.Context, req *hello.SayHiRequest) (*hello.SayHiReply, error) {
	fmt.Println(req.GetMsg())
	return &hello.SayHiReply{Msg: "Hello " + req.GetMsg()}, nil
}

func (s *GreeterServer) Chat(stream hello.Greeter_ChatServer) error {
	for {
		req, err := stream.Recv()
		fmt.Println("req:", req.GetMsg())
		if err == io.EOF {

			if err := stream.Send(&hello.ChatReply{Msg: "hi" + string(req.GetMsg())}); err != nil {
				return err
			}
		}
	}
}

func main() {
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &GreeterServer{})
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
