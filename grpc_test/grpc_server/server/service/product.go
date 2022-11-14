package service

import "time"

var ProdService = &productService{}

type productService struct {
}

func (p *productService) GetProductStockServerStream(request *ProductRequest, stream ProdService_GetProductStockServerStreamServer) error {
	count := 0
	for {
		rep := &ProductResponse{ProdStock: request.ProdId}
		err := stream.Send(rep)
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
		count++
		if count > 10 {
			return nil
		}
	}
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {}
