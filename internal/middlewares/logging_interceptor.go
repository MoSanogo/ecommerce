package middlewares

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

// LoggingInterceptor is a gRPC server interceptor that logs incoming requests
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Log the request details (for demonstration purposes, we just print to console)
	// In a real application, you might want to use a structured logger
	logMessage := "Received request for method: " + info.FullMethod
	if req != nil {
		logMessage += ", with request: " + fmt.Sprintf("%v", req)
	}
	log.Println(logMessage)

	// Call the handler to proceed with the request
	return handler(ctx, req)
}

func StreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// Log the stream request details (for demonstration purposes, we just print to console)
	log.Printf("Received stream request for method: %v", info.FullMethod)

	// Call the handler to proceed with the stream request
	return handler(srv, ss)
}

func ErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Call the handler to proceed with the request
	resp, err := handler(ctx, req)
	if err != nil {
		// Log the error (for demonstration purposes, we just print to console)
		log.Printf("Error occurred in method %s: %v\n", info.FullMethod, err)
	}
	return resp, err
}
