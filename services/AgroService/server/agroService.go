package server

import (
	"context"

	pb "github.com/longkhan786/dhareja-platform/agroservice/gen"
)

type AgroService struct {
	pb.UnimplementedAgroTradersServiceServer
}

var products = []*pb.Product{
	{Id: "1", Name: "Wheat", Quantity: 100},
	{Id: "2", Name: "Rice", Quantity: 75},
	{Id: "3", Name: "Corn", Quantity: 40},
}

func (s *AgroService) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.Product, error) {
	for _, p := range products {
		if p.Id == req.Id {
			return p, nil
		}
	}
	return nil, nil
}

func (s *AgroService) ListProducts(ctx context.Context, _ *pb.Empty) (*pb.ProductList, error) {
	return &pb.ProductList{Products: products}, nil
}
