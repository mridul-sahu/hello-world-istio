package main

import (
	"bitbucket.org/crispyapp/services/deployment/hello/protos"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func main() {
	ctxHeader := metadata.NewOutgoingContext(
		context.Background(),
		metadata.New(map[string]string{"request-service": "hello-0"}))

	conn, err := grpc.DialContext(ctxHeader, "35.200.128.155:80", grpc.WithInsecure(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Printf("Dial failed: %v", err)
		panic("User Client not available")
	}
	log.Printf("Dial did not fail %v", conn)

	defer conn.Close()

	client := protos.NewHelloServiceClient(conn)

	res, err := client.Hello(ctxHeader, &empty.Empty{})

	if err != nil {
		log.Panicf("Failed! %v", err)
	}

	log.Printf(res.Message)
}
