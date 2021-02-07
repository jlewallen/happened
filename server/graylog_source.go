package main

import (
	"context"
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog/client"
)

type GraylogSource struct {
	gcl *client.Client
}

func NewGraylogSource() (*GraylogSource, error) {
	return &GraylogSource{
		gcl: nil,
	}, nil
}

func (s *GraylogSource) Initialize(ctx context.Context) error {
	gcl, err := client.NewClientV3("code.conservify.org", "admin", "admin")
	if err != nil {
		return err
	}

	s.gcl = gcl

	return nil
}

func (s *GraylogSource) Tail(ctx context.Context, position Position) (*TailResponse, error) {
	role, ei, err := s.gcl.GetRole("Admin")
	if err != nil {
		return nil, err
	}

	fmt.Println(ei.Response.StatusCode)
	fmt.Println(role.Name)

	return &TailResponse{
		//
	}, nil
}

func (s *GraylogSource) Written() uint64 {
	return uint64(0)
}
