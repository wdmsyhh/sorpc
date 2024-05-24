package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"sorpc/proto/product"
	p "sorpc/service/product/service"
)

func main() {
	grpcServer := grpc.NewServer()
	product.RegisterProductServiceServer(grpcServer, new(p.ProductService))
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Product server start")
	log.Fatal(grpcServer.Serve(lis))
	log.Println("Product server end")
}
