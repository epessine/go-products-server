package grpc

import (
	"context"

	"github.com/epessine/go-products-server/application/grpc/pb"
	"github.com/epessine/go-products-server/application/usecase"
)

type ProductService struct {
	UseCase usecase.Product
	pb.UnimplementedProductServiceServer
}

func (p *ProductService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := p.UseCase.CreateProduct(in.Name, in.Description, in.Price)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          int32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price: product.Price,
		},
	}, nil
}

func (p *ProductService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products, err := p.UseCase.FindProducts()
	if err != nil {
		return nil, err
	}

	var responseProducts []*pb.Product

	for _, product := range products {
        responseProducts = append(responseProducts, &pb.Product{
			Id:          int32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price: product.Price,
		})
    }

	return &pb.FindProductsResponse{
		Products: responseProducts,
	}, nil
}

func NewProductService(usecase usecase.Product) *ProductService {
	return &ProductService{
		UseCase: usecase,
	}
}
