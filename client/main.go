package main

import (
	"context"
	"log"
	"time"

	"github.com/mahesh-dilhan/gooauth2/pkg/proto/credit"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func main() {
	log.Println("Client running ...")

	rpcCreds := oauth.NewOauthAccess(&oauth2.Token{AccessToken: "client-x-id"})
	trnCreds, err := credentials.NewClientTLSFromFile("./client/cert/public.crt", "localhost")
	if err != nil {
		log.Fatalln(err)
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(trnCreds),
		grpc.WithPerRPCCredentials(rpcCreds),
	}

	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":50051", opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := credit.NewCreditServiceClient(conn)

	request := &credit.CreditRequest{Amount: 1990.01}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Credit(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Response:", response.GetConfirmation())
}
