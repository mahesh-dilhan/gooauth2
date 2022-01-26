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

}
