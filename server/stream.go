package main

import (
	_ "github.com/jlewallen/happened/server/common"
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
