package server

import (
	"context"

	pb "github.com/longkhan786/dhareja-platform/agroservice/gen"
)

// AgroService implements the gRPC service defined in the proto
type AgroService struct {
	pb.UnimplementedAgroTradersServiceServer
}

// Fake product database
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
	return nil, nil // or: return nil, status.Error(codes.NotFound, "product not found")
}

func (s *AgroService) ListProducts(ctx context.Context, _ *pb.Empty) (*pb.ProductList, error) {
	return &pb.ProductList{Products: products}, nil
}
