package main

import (
	"context"
)

type Position struct {
	Encoded *string `json:"encoded"`
}

type Source interface {
	Initialize(ctx context.Context) error
	Tail(ctx context.Context, position Position) (*TailResponse, error)
	Written() int64
}
