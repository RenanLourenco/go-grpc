package main

import (
	"context"
	"log"
	"net"

	"github.com/RenanLourenco/go-grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error){
 return &invoicer.CreateResponse{
	Pdf: []byte("test"),
	Docx: []byte("test"),
 }, nil
}

func main(){
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	server := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(server,service)
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("fail to serve: %s", err)
	}
}