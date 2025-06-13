package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "api/internal/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHi(ctx, &pb.SayHiRequest{Msg: "我从客户端来"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	stream, err := client.Chat(ctx)
	if err != nil {
		panic(err)
	}

	go func() {
		data := &pb.ChatRequest{
			Msg: "hello",
		}
		if err := stream.Send(data); err != nil {
			panic(err)
		}
		if err := stream.CloseSend(); err != nil {
			panic(err)
		}
	}()

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Println("data : ", data.GetMsg())
	}

}
