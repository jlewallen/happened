package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"sync"

	"github.com/armon/circbuf"
)

const (
	BufferSize = int64(4096 * 10)
	Key        = "<todo>"
)

type TcpTextSource struct {
	lock       sync.RWMutex
	conn       net.Conn
	buffer     *circbuf.Buffer
	bufferSize int64
	written    int64
}

func NewTcpTextSource(conn net.Conn) (s *TcpTextSource, err error) {
	bufferSize := BufferSize
	buffer, err := circbuf.NewBuffer(bufferSize)
	if err != nil {
		return nil, err
	}

	return &TcpTextSource{
		conn:       conn,
		buffer:     buffer,
		bufferSize: bufferSize,
		written:    0,
	}, nil
}

func (s *TcpTextSource) Initialize(ctx context.Context) error {
	return nil
}

func (s *TcpTextSource) Tail(ctx context.Context, pos Position) (*TailResponse, error) {
	s.lock.Lock()

	defer s.lock.Unlock()

	position := int64(0)
	if pos.Encoded != nil {
		if p, err := strconv.ParseInt(*pos.Encoded, 10, 64); err != nil {
			return nil, err
		} else {
			position = p
		}
	}

	dropped := int64(0)
	buffered := s.buffer.Bytes()
	data := buffered
	if position > s.written {
		// Way out of bounds, so we can just treat this as fresh from the beginning.
		log.Printf("(%s) written=%d position=%d overflow", Key, s.written, position)
	} else if s.written <= s.bufferSize {
		// Buffer is still filling up, so we're going to return from position to the end of the buffer.
		log.Printf("(%s) written=%d position=%d filling", Key, s.written, position)
		data = buffered[position:]
	} else {
		// Buffer is full so we need to adjust position based on that, but the buffer may not have all we need.
		remaining := s.written - position
		if remaining > s.bufferSize {
			// Buffer has less than we need, so we return the whole thing and tell the client.
			log.Printf("(%s) written=%d position=%d remaining=%d dropped logs", Key, s.written, position, remaining)
			dropped = remaining - s.bufferSize
		} else {
			// Most common scenario, we return the remaining bytes from the end of the buffer.
			data = buffered[s.bufferSize-remaining:]
			log.Printf("(%s) written=%d position=%d remaining=%d", Key, s.written, position, remaining)
		}
	}

	return &TailResponse{
		MoreURL: fmt.Sprintf("position=%d", s.written),
		Dropped: dropped > 0,
		Blocks: []string{
			string(data),
		},
	}, nil
}

func (s *TcpTextSource) Written() int64 {
	return s.written
}

func (s *TcpTextSource) handle(ctx context.Context) {
	if err := s.tail(); err != nil {
		log.Printf("error: %v", err)
	}
}

func (s *TcpTextSource) write(b []byte) error {
	s.lock.Lock()

	defer s.lock.Unlock()

	bytesWrote, err := s.buffer.Write(b)
	if err != nil {
		if err == io.EOF {
			log.Printf("[%s] eof", Key)
			return nil
		}
		log.Printf("[%s] [ERROR] writing: %v", Key, err.Error())
		return err
	}

	s.written += int64(bytesWrote)

	if false {
		log.Printf("[%s] write %d", Key, bytesWrote)
	}

	return nil
}

func (s *TcpTextSource) tail() error {
	defer s.conn.Close()

	buf := make([]byte, 4096)
	for {
		bytesRead, err := s.conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("[%s] eof", Key)
				return nil
			}
			log.Printf("[%s] [ERROR] reading: %v", Key, err.Error())
			return err
		}

		if bytesRead > 0 {
			if false {
				log.Printf("[%s] read %d", Key, bytesRead)
			}

			if err := s.write(buf[:bytesRead]); err != nil {
				return err
			}
		}
	}
}
