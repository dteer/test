package main

import (
	"grpc_client/server/service"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

func tlsone() (*grpc.Server, error) {
	// 无证书
	server := grpc.NewServer()
	return server, nil
}

func main() {
	server, _ := tlsone()
	service.RegisterProdServiceServer(server, service.ProdService)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = server.Serve(listener)
}
