package main

import (
	"github.com/jlewallen/happened/server/common"
)

type Stream struct {
	key    string
	name   string
	source Source
}

func NewStream(key string, source Source) (*Stream, error) {
	return &Stream{
		key:    key,
		name:   key,
		source: source,
	}, nil
}

func (s *Stream) Configure(h *common.Handshake) error {
	if len(h.Name) > 0 {
		s.name = h.Name
	}
	return nil
}
