package service

import (
	"context"

	productv1 "ecommerce-grpc-api/gen/go/product/v1"
)

type ProductService interface {
	GetProduct(ctx context.Context, req *productv1.GetProductRequest) (*productv1.GetProductResponse, error)
	CreateProduct(ctx context.Context, req *productv1.CreateProductRequest) (*productv1.CreateProductResponse, error)
	UpdateProduct(ctx context.Context, req *productv1.UpdateProductRequest) (*productv1.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, req *productv1.DeleteProductRequest) (*productv1.DeleteProductResponse, error)
	ListProducts(ctx context.Context, req *productv1.ListProductsRequest) (*productv1.ListProductsResponse, error)
}

type ProductServiceImpl struct {
	productv1.UnimplementedProductServiceServer
}

func NewProductService() *ProductServiceImpl {
	return &ProductServiceImpl{}
}

func (s *ProductServiceImpl) GetProduct(ctx context.Context, req *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	return &productv1.GetProductResponse{}, nil
}

func (s *ProductServiceImpl) CreateProduct(ctx context.Context, req *productv1.CreateProductRequest) (*productv1.CreateProductResponse, error) {
	return &productv1.CreateProductResponse{}, nil
}

func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, req *productv1.UpdateProductRequest) (*productv1.UpdateProductResponse, error) {
	return &productv1.UpdateProductResponse{}, nil
}

func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, req *productv1.DeleteProductRequest) (*productv1.DeleteProductResponse, error) {
	return &productv1.DeleteProductResponse{}, nil
}

func (s *ProductServiceImpl) ListProducts(ctx context.Context, req *productv1.ListProductsRequest) (*productv1.ListProductsResponse, error) {
	return &productv1.ListProductsResponse{}, nil
}
