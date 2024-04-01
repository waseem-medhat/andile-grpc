package main

import (
	"context"
	"log"
	"net"

	"github.com/wipdev-tech/andile-grpc/internal/invoicer"
	"google.golang.org/grpc"
)

type invoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s invoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{}, nil
}

func main() {
	serverRegistrar := grpc.NewServer()
	service := &invoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, service)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(serverRegistrar.Serve(l))
}
