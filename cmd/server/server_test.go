package main

import (
	"context"
	"fmt"
	"github.com/kxplxn/learning_go-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	req := pb.EndRequest{
		Id:       "end",
		End:      timestamppb.Now(),
		Distance: 3.14,
	}
	var srv Rides
	resp, err := srv.End(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Id != req.Id {
		t.Fatalf("bad response id: got %v, want %v", resp.Id, req.Id)
	}
}

func TestServerE2E(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	srv := createServer()
	go func() { _ = srv.Serve(lis) }()

	port := lis.Addr().(*net.TCPAddr).Port
	addr := fmt.Sprintf("localhost:%d", port)

	creds := insecure.NewCredentials()
	conn, err := grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	c := pb.NewRidesClient(conn)

	req := pb.EndRequest{
		Id:       "end",
		End:      timestamppb.Now(),
		Distance: 3.14,
	}
	resp, err := c.End(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Id != req.Id {
		t.Fatalf("bad response: got %s, want %s", resp.Id, req.Id)
	}
}
