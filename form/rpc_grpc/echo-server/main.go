package echo_server

import (
	"form/rpc_grpc/echo"
	"form/rpc_grpc/echo-server/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, server.NewEchoServer())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
