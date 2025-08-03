package main // 注意：可执行文件应该用 package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"form/rpc_grpc/echo" // 确认这是正确的本地路径
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure" // 添加这行
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := echo.NewEchoClient(conn)
	ctx := context.Background()
	in := &echo.EchoMsg{
		Name: "Jack",
		Addr: &echo.Address{
			City:     "Beijing",
			Province: "100000",
		},
		Birthday: timestamppb.New(time.Now()),
		Data:     []byte("今天天气不错"),
		Gender:   echo.Gender_MALE,
		Hobbies:  []string{"羽毛球", "写代码", "看博客"},
	}

	res, err := client.UnaryEcho(ctx, in)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(res)

	stream, err := client.ClientStreamEcho(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for i := 0; i < 5; i++ {
		err := stream.Send(in)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}
	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(result)
}
