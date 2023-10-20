package main

import (
	"context"
	"log"
	"net"

	"github.com/kxplxn/learning_go-grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := ":9292"

	// create a TCP listener on given port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error: can't listen - %s", err)
	}

	srv := grpc.NewServer()
	var u Rides
	pb.RegisterRidesServer(srv, &u)
	reflection.Register(srv)

	log.Printf("info: server ready on %s", addr)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error: can't serve - %s", err)
	}
}

type Rides struct {
	pb.UnimplementedRidesServer
}

func (r *Rides) Start(
	ctx context.Context, req *pb.StartRequest,
) (*pb.StartResponse, error) {
	// TODO: validate req
	resp := pb.StartResponse{
		Id: req.Id,
	}

	// TODO: work
	return &resp, nil
}

func (r *Rides) End(
	ctx context.Context, req *pb.EndRequest,
) (*pb.EndResponse, error) {
	return &pb.EndResponse{Id: req.Id}, nil
}
