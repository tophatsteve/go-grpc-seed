package sample

import (
	"golang.org/x/net/context"
)

type Service interface {
	Generate(ctx context.Context, in *Empty) (*GenerateResponse, error)
}

type server struct{}

func NewService() Service {
	return server{}
}

func (s server) Generate(ctx context.Context, in *Empty) (*GenerateResponse, error) {
	resp := GenerateResponse{
		Value: generateMessage(),
	}
	return &resp, nil
}

func generateMessage() string {
	return "Success"
}
