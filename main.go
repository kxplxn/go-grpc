package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kxplxn/learning_go-grpc/pb"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	req := pb.StartRequest{
		Id:       "526a122b-1808-4f80-acb2-124ca3136f2a",
		DriverId: "007",
		Location: &pb.Location{
			Lat: 51.1293414,
			Lng: -0.1241513,
		},
		PassengerIds: []string{"M", "Q"},
		Time:         timestamppb.Now(),
		Type:         pb.RideType_POOL,
	}
	fmt.Println(&req)

	data, err := proto.Marshal(&req)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	var req2 pb.StartRequest
	if err := proto.Unmarshal(data, &req2); err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(&req2)

	fmt.Printf("proto size: %d\n", len(data))

	jData, err := protojson.Marshal(&req)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Printf("json size: %d\n", len(jData))

	if _, err := os.Stdout.Write(jData); err != nil {
		log.Fatalf("error: %s", err)
	}
}
