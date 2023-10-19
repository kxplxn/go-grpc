package main

import (
	"testing"
	"time"

	"github.com/kxplxn/learning_go-grpc/pb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestInvoice(t *testing.T) {
	inv := pb.Invoice{
		Id: "2023-0123",
		Time: timestamppb.New(
			time.Date(2023, time.January, 7, 13, 45, 0, 0, time.UTC),
		),
		Customer: "Wile E. Coyote",
		Items: []*pb.LineItem{
			{Sku: "hammer-20", Amount: 1, Price: 249},
			{Sku: "nail-9", Amount: 100, Price: 1},
			{Sku: "glue5", Amount: 1, Price: 799},
		},
	}
	t.Logf("%v\n", &inv)

	data, err := proto.Marshal(&inv)
	if err == nil {
		t.Logf("size: %d\n", len(data))
	} else {
		t.Errorf("ERROR: %s\n", err)
	}
}
