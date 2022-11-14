package service

import (
	"context"
	"fmt"
	"io"
)

var ProdService = &productService{}

type productService struct {
}

func (p *productService) GetProductStock(context context.Context, request *ProductRequest) (*ProductResponse, error) {
	// 实现具体的业务逻辑
	stock := p.GetStockById(request.ProdId)
	return &ProductResponse{ProdStock: stock}, nil
}

func (p *productService) UpdateProductStockClientStream(stream ProdService_UpdateProductStockClientStreamServer) error {
	count := 0
	for {
		// 接收客户端的消息
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println("服务接收到的流：", recv.ProdId)
		count++
		if count > 10 {
			resp := &ProductResponse{ProdStock: recv.ProdId}
			err := stream.SendAndClose(resp)
			if err != nil {
				return err
			}
		}
	}

}

func (p *productService) GetStockById(id int32) int32 {
	return 100
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {}
