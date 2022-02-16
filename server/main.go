package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/mahesh-dilhan/gooauth2/pkg/proto/credit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	credit.UnimplementedCreditServiceServer
}

func main() {
	log.Println("Server running ...")
	cert, err := tls.LoadX509KeyPair("./server/cert/public.crt", "./server/cert/private.key")
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}
	opts := []grpc.ServerOption{
		// Intercept request to check the token.
		grpc.UnaryInterceptor(validateToken),
		// Enable TLS for all incoming connections.
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	srv := grpc.NewServer(opts...)
	credit.RegisterCreditServiceServer(srv, &server{})
	log.Fatalln(srv.Serve(lis))
}
func (s *server) Credit(ctx context.Context, request *credit.CreditRequest) (*credit.CreditResponse, error) {
	log.Println(fmt.Sprintf("Request: %g", request.GetAmount()))

	return &credit.CreditResponse{Confirmation: fmt.Sprintf("Credited %g", request.GetAmount())}, nil
}
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// If you have more than one client then you will have to update this line.
	return token == "client-x-id"
}
func validateToken(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
}
