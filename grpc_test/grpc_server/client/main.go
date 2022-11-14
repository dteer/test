package main

import (
	"context"
	"fmt"
	"grpc_server/client/service"
	"io"
	"log"

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
	request := &service.ProductRequest{
		ProdId: 123,
	}
	stream, err := proClient.GetProductStockServerStream(context.Background(), request)
	if err != nil {
		log.Fatal("获取流出错", err)
	}
	for {
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端结束数据完成")
				err := stream.CloseSend()
				if err != nil {
					log.Fatal(err)
				}
				break
			}
			log.Fatal(err)
		}
		fmt.Println("客户端收到的流：", recv.ProdStock)
	}

}
