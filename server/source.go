package main

import (
	"context"
)

type Position struct {
	Encoded *string `json:"encoded"`
}

type SourceMeta struct {
	Name    string `json:"name"`
	Written int64  `json:"written"`
}

type Source interface {
	Initialize(ctx context.Context) error
	Tail(ctx context.Context, position Position) (*TailResponse, error)
	Meta() *SourceMeta
}
