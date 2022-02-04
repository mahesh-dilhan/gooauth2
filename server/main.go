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
