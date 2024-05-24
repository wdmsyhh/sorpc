genStub() {
    #product
    protoc --go_out=.  --go-grpc_out=. ./proto/product/*.proto
    protoc --grpc-gateway_out=. --grpc-gateway_opt grpc_api_configuration=./proto/product/api.yaml ./proto/product/*.proto
    #order
    protoc --go_out=.  --go-grpc_out=. ./proto/order/*.proto
    protoc --grpc-gateway_out=. --grpc-gateway_opt grpc_api_configuration=./proto/order/api.yaml ./proto/order/*.proto
}

deleteStub() {
  rm proto/*/*.pb.go
  rm proto/*/*.pb.gw.go
}

#deleteStub

genStub