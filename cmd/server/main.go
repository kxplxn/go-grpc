package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"

	"github.com/kxplxn/learning_go-grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "no metadata")
	}
	log.Printf("info: api_key - %s", md["api_key"])

	return &pb.StartResponse{Id: req.Id}, nil
}

func (r *Rides) End(
	_ context.Context, req *pb.EndRequest,
) (*pb.EndResponse, error) {
	return &pb.EndResponse{Id: req.Id}, nil
}

func (c *Rides) Location(stream pb.Rides_LocationServer) error {
	count := int64(0)
	driverID := ""

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "error: can't read - %s", err)
		}

		// update database...

		driverID = req.DriverId
		count++
	}

	resp := pb.LocationResponse{
		DriverId: driverID,
		Count:    count,
	}
	return stream.SendAndClose(&resp)
}
