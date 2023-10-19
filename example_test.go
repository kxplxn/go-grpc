package main

import (
	"fmt"
	"time"
)

func ExampleInvoice() {
	// TODO: Change to protobuf invoice
	inv := Invoice{
		ID:       "2023-0123",
		Time:     time.Date(2023, time.January, 7, 13, 45, 0, 0, time.UTC),
		Customer: "Wile E. Coyote",
		Items: []LineItem{
			{SKU: "hammer-20", Amount: 1, Price: 249},
			{SKU: "nail-9", Amount: 100, Price: 1},
			{SKU: "glue5", Amount: 1, Price: 799},
		},
	}
	fmt.Printf("%v\n", &inv)
	// TODO: encode to []byte using protobuf
	data, err := []byte(nil), error(nil)
	if err == nil {
		fmt.Printf("size: %d\n", len(data))
	} else {
		fmt.Printf("ERROR: %s\n", err)
	}
}
