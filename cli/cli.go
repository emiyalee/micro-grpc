package main

import (
	"context"
	"log"

	pb "cli/proto/sum"

	"google.golang.org/grpc"
)

func main() {
	var err error
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	c := pb.NewSumServiceClient(conn)
	rsp, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 100,
		B: 1,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(rsp.Sum)
}
