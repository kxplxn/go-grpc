package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kxplxn/learning_go-grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	addr := "localhost:9292"
	creds := insecure.NewCredentials()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer func() { _ = conn.Close() }()

	log.Printf("info: connected to %s", addr)
	c := pb.NewRidesClient(conn)
	fmt.Println(c)

	req := pb.StartRequest{
		Id:       "47a74960d6204a52b1bece53221eb458",
		DriverId: "007",
		Location: &pb.Location{
			Lat: 51.234125,
			Lng: -1.124131,
		},
		PassengerIds: []string{"M", "Q"},
		Time:         timestamppb.Now(),
		Type:         pb.RideType_POOL,
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "api_key", "s3cr3t")
	resp, err := c.Start(ctx, &req)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(resp)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	respEnd, err := c.End(ctx, &pb.EndRequest{
		Id:       "47a74960d6204a52b1bece53221eb458",
		End:      timestamppb.Now(),
		Distance: 21.32,
	})
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(respEnd)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.Location(ctx)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	lReq := pb.LocationRequest{
		DriverId: "007",
		Location: &pb.Location{
			Lat: 12.4123,
			Lng: -12.4123,
		},
	}
	for i := 0.0; i < 0.03; i += 0.01 {
		lReq.Location.Lat += i
		if err := stream.Send(&lReq); err != nil {
			log.Fatalf("error: %s", err)
		}
	}
	lResp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(lResp)
}
