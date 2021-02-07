package main

import (
	_ "context"
)

type Stream struct {
	key    string
	source Source
}

func NewStream(key string, source Source) (*Stream, error) {
	return &Stream{
		key:    key,
		source: source,
	}, nil
}
