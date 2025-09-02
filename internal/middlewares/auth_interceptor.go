package middlewares

import (
	"context"
	"ecommerce-grpc-api/internal/jwt"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type ContextKey string

const (
	UserIDKey ContextKey = "userID"
	RolesKey  ContextKey = "roles"
)

// AuthInterceptor is a gRPC server interceptor that checks for JWT authentication

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//Skip authentication for specific methods
	if info.FullMethod == "/ecommerce.AuthService/Login" || info.FullMethod == "/ecommerce.AuthService/Register" {
		return handler(ctx, req)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}

	token := strings.TrimPrefix(authHeader[0], "Bearer ")
	if token == "" {
		return nil, status.Error(codes.Unauthenticated, "empty token")
	}

	claims, err := jwt.ValidateToken(token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("invalid token: %v", err))
	}

	// Store user ID and roles in context
	ctx = context.WithValue(ctx, UserIDKey, claims.UserID)
	ctx = context.WithValue(ctx, RolesKey, claims.Roles)

	// Call the handler with the updated context
	return handler(ctx, req)
}
