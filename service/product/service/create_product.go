package service

import (
	"context"

	"sorpc/proto/product"
)

func (*ProductService) CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	response := &product.CreateProductResponse{
		Id:   req.Id,
		Name: req.Name,
	}

	return response, nil
}
