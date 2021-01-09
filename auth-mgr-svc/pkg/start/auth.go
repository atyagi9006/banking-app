package start

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("---> unary Interceptor: ", info.FullMethod)
	return handler(ctx, req)
}
