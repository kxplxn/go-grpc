package main

import (
	"context"
	"log"
	"net/http"

	"github.com/kxplxn/learning_go-grpc/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"localhost:9292",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial server: %s", err)
	}

	mux := runtime.NewServeMux()
	err = pb.RegisterRidesHandler(ctx, mux, conn)
	if err != nil {
		log.Fatalf("failed to register handler: %s", err)
	}

	addr := ":8080"
	log.Printf("gateway server starting on %s", addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
