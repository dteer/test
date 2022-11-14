package main

import (
	"context"
	"fmt"
	"grpc_client/client/service"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func tlsone() (*grpc.ClientConn, error) {
	// 无认证状态
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	return conn, err
}

func main() {

	conn, err := tlsone()
	if err != nil {
		log.Fatal(err)
	}
	// 退出时关闭链接
	defer conn.Close()

	proClient := service.NewProdServiceClient(conn)

	stream, err := proClient.UpdateProductStockClientStream(context.Background())
	if err != nil {
		log.Fatal("获取流错误", err)
	}
	rsp := make(chan struct{}, 1)
	go prodRequest(stream, rsp)
	select {
	case <-rsp:
		recv, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatal(err)
		}
		stock := recv.ProdStock
		fmt.Println("客户端收到响应：", stock)
	}
}

func prodRequest(stream service.ProdService_UpdateProductStockClientStreamClient, rsp chan struct{}) {
	count := 0
	for {
		request := &service.ProductRequest{
			ProdId: 123,
		}
		err := stream.Send(request)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
		count++
		if count > 10 {
			rsp <- struct{}{}
			break
		}
	}
}
