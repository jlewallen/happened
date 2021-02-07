package main

import (
	"fmt"
	"log"
	"sync"
)

type StreamManager struct {
	lock    sync.RWMutex
	counter int32
	streams map[string]*Stream
}

func NewStreamManager() *StreamManager {
	return &StreamManager{
		streams: make(map[string]*Stream),
	}
}

func (sm *StreamManager) Close() error {
	return nil
}

func (sm *StreamManager) AddStream(source Source) (*Stream, error) {
	sm.lock.Lock()

	defer sm.lock.Unlock()

	key := fmt.Sprintf("stream-%d", sm.counter)
	stream, err := NewStream(key, source)
	if err != nil {
		return nil, err
	}

	sm.streams[stream.key] = stream
	sm.counter += 1

	log.Printf("[sm] added %s", stream.key)

	return stream, nil
}

func (sm *StreamManager) RemoveStream(s *Stream) error {
	sm.lock.Lock()

	defer sm.lock.Unlock()

	delete(sm.streams, s.key)

	log.Printf("[sm] removed %s", s.key)

	return nil
}

func (sm *StreamManager) Lookup(key string) (s *Stream) {
	sm.lock.RLock()

	defer sm.lock.RUnlock()

	return sm.streams[key]
}
