package handler

import (
	"context"

	proto "toz/pkg/proto"
	service "toz/pkg/url/service"
)

type GRPCServer struct {
	services *service.ServiceUrl
}

func NewGrpcServer(services *service.ServiceUrl) *GRPCServer {
	return &GRPCServer{services: services}
}

func (s *GRPCServer) GetFullURL(ctx context.Context, req *proto.GetFullURLRequest) (*proto.GetFullURLResponse, error) {
	if answer, err := s.services.GetFull(req.ShortUrl); err != nil {
		return &proto.GetFullURLResponse{
			FullUrl: "",
		}, err
	} else {
		return &proto.GetFullURLResponse{
			FullUrl: answer,
		}, nil
	}
}

func (s *GRPCServer) CreateShotrUrl(ctx context.Context, req *proto.CreateShotrUrlRequest) (*proto.CreateShotrUrlResponse, error) {
	if answer, err := s.services.Create(req.FullUrl); err != nil {
		return &proto.CreateShotrUrlResponse{
			ShortUrl: "",
		}, err
	} else {
		return &proto.CreateShotrUrlResponse{
			ShortUrl: answer,
		}, nil
	}
}
