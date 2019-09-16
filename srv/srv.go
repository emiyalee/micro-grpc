package main

import (
	"context"
	"log"

	pb "srv/proto/sum"

	"github.com/micro/go-micro"
)

type sumService struct{}

func (s *sumService) Sum(ctx context.Context, req *pb.SumRequest, rsp *pb.SumResponse) error {
	rsp.Sum = req.A + req.B
	return nil
}

func main() {
	var err error
	service := micro.NewService(
		micro.Name("sum"),
	)

	service.Init()

	pb.RegisterSumServiceHandler(service.Server(), &sumService{})

	if err = service.Run(); err != nil {
		log.Fatalf("error := %s\n", err)
	}
}
