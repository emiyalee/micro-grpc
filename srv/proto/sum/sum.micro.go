// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sum.proto

/*
Package sum is a generated protocol buffer package.

It is generated from these files:
	sum.proto

It has these top-level messages:
	SumRequest
	SumResponse
*/
package sum

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SumService service

type SumService interface {
	Sum(ctx context.Context, in *SumRequest, opts ...client.CallOption) (*SumResponse, error)
}

type sumService struct {
	c    client.Client
	name string
}

func NewSumService(name string, c client.Client) SumService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "sum"
	}
	return &sumService{
		c:    c,
		name: name,
	}
}

func (c *sumService) Sum(ctx context.Context, in *SumRequest, opts ...client.CallOption) (*SumResponse, error) {
	req := c.c.NewRequest(c.name, "SumService.Sum", in)
	out := new(SumResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SumService service

type SumServiceHandler interface {
	Sum(context.Context, *SumRequest, *SumResponse) error
}

func RegisterSumServiceHandler(s server.Server, hdlr SumServiceHandler, opts ...server.HandlerOption) error {
	type sumService interface {
		Sum(ctx context.Context, in *SumRequest, out *SumResponse) error
	}
	type SumService struct {
		sumService
	}
	h := &sumServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SumService{h}, opts...))
}

type sumServiceHandler struct {
	SumServiceHandler
}

func (h *sumServiceHandler) Sum(ctx context.Context, in *SumRequest, out *SumResponse) error {
	return h.SumServiceHandler.Sum(ctx, in, out)
}
