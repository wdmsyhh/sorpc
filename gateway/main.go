package main

import (
	"context"
	"log"
	"net/http"

	"sorpc/proto/order"
	"sorpc/proto/product"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	startGRPCGateway()
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers: true, //枚举字段的值使用数字
			UseProtoNames:  true, //编译器生成默认使用的驼峰式 JSON 标签。如果您想使用原始文件中使用的确切大小写，请设置UseProtoNames: true
		},
	}))
	err := product.RegisterProductServiceHandlerFromEndpoint(c, mux, ":8081", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("RegisterProductServiceHandlerFromEndpoint failed: %v", err)
	}
	err = order.RegisterOrderServiceHandlerFromEndpoint(c, mux, ":8082", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("RegisterOrderServiceHandlerFromEndpoint failed: %v", err)
	}
	log.Println("Gateway start")
	err = http.ListenAndServe(":8080", mux) // grpc gateway 的端口
	if err != nil {
		log.Fatalf("cann't listen and serve: %v", err)
	}
	log.Println("Gateway end")
}
