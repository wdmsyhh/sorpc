package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"sorpc/proto/order"
	o "sorpc/service/order/service"
)

func main() {
	grpcServer := grpc.NewServer()
	order.RegisterOrderServiceServer(grpcServer, new(o.OrderService))
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Order server start")
	log.Fatal(grpcServer.Serve(lis))
	log.Println("Order server end")
}
