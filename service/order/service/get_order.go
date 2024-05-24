package service

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"sorpc/proto/order"
	"sorpc/proto/product"
)

type OrderService struct {
	order.UnimplementedOrderServiceServer
}

func (*OrderService) GetOrder(ctx context.Context, req *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := product.NewProductServiceClient(conn)
	response, err := client.GetProduct(context.Background(), &product.GetProductRequest{Id: "aaa"})
	if err != nil {
		log.Fatal(err)
	}

	return &order.GetOrderResponse{
		Id:   response.Id,
		Name: response.Name,
	}, nil
}
