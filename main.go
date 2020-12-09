package main

import (
	_ "bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"

	"github.com/armon/circbuf"
)

type Stream struct {
	key        string
	conn       net.Conn
	lock       sync.Mutex
	buffer     *circbuf.Buffer
	bufferSize int64
	written    int64
}

const (
	BufferSize = int64(4096 * 10)
)

func NewStream(key string, conn net.Conn) (*Stream, error) {
	bufferSize := BufferSize
	buffer, err := circbuf.NewBuffer(bufferSize)
	if err != nil {
		return nil, err
	}
	return &Stream{
		key:        key,
		conn:       conn,
		buffer:     buffer,
		bufferSize: bufferSize,
		written:    0,
	}, nil
}

func (s *Stream) handle() {
	if err := s.tail(); err != nil {
		log.Printf("error: %v", err)
	}
}

func (s *Stream) write(b []byte) error {
	s.lock.Lock()

	defer s.lock.Unlock()

	bytesWrote, err := s.buffer.Write(b)
	if err != nil {
		if err == io.EOF {
			log.Printf("[%s] eof", s.key)
			return nil
		}
		log.Printf("[%s] [ERROR] writing: %v", s.key, err.Error())
		return err
	}

	s.written += int64(bytesWrote)

	if false {
		log.Printf("[%s] write %d", s.key, bytesWrote)
	}

	return nil
}

func (s *Stream) tail() error {
	defer s.conn.Close()

	buf := make([]byte, 4096)
	for {
		bytesRead, err := s.conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("[%s] eof", s.key)
				return nil
			}
			log.Printf("[%s] [ERROR] reading: %v", s.key, err.Error())
			return err
		}

		if bytesRead > 0 {
			if false {
				log.Printf("[%s] read %d", s.key, bytesRead)
			}

			if err := s.write(buf[:bytesRead]); err != nil {
				return err
			}
		}
	}
}

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

func (sm *StreamManager) AddStream(conn net.Conn) (*Stream, error) {
	sm.lock.Lock()

	defer sm.lock.Unlock()

	key := fmt.Sprintf("stream-%d", sm.counter)
	stream, err := NewStream(key, conn)
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

func listen(sm *StreamManager) error {
	address := fmt.Sprintf("0.0.0.0:%d", 2570)
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	defer l.Close()

	log.Printf("listening on " + address)

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		stream, err := sm.AddStream(conn)
		if err != nil {
			return err
		}

		go func() {
			stream.handle()

			if err := sm.RemoveStream(stream); err != nil {
				log.Printf("error %v", err)
			}
		}()
	}

	return nil
}

type handlerFunc func(sm *StreamManager, res http.ResponseWriter, req *http.Request) error

func middleware(sm *StreamManager, handler handlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Printf("[http] %s %s", req.Method, req.URL.Path)
		if err := handler(sm, res, req); err != nil {
			log.Printf("error %v", err)
		}
	}
}

func setupResponse(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Hpn-More-Url, Hpn-Dropped")
	res.Header().Set("Access-Control-Expose-Headers", "Content-Type, Hpn-More-Url, Hpn-Dropped")
}

type StreamResponse struct {
	Key     string `json:"key"`
	URL     string `json:"url"`
	Written int64  `json:"written"`
}

type StreamsResponse struct {
	Streams []*StreamResponse `json:"streams"`
}

func streamsHandler(sm *StreamManager, res http.ResponseWriter, req *http.Request) error {
	sm.lock.RLock()

	defer sm.lock.RUnlock()

	streams := make([]*StreamResponse, 0)

	for key, stream := range sm.streams {
		streams = append(streams, &StreamResponse{
			Key:     key,
			URL:     fmt.Sprintf("/v1/streams/%s", key),
			Written: stream.written,
		})
	}

	response := &StreamsResponse{
		Streams: streams,
	}

	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	setupResponse(res, req)

	res.WriteHeader(http.StatusOK)
	res.Write(data)

	return nil
}

func streamHandler(sm *StreamManager, res http.ResponseWriter, req *http.Request) error {
	vars := mux.Vars(req)
	key := vars["key"]

	position := int64(0)
	maybePosition := req.URL.Query()["position"]
	if len(maybePosition) == 1 {
		if p, err := strconv.ParseInt(maybePosition[0], 10, 64); err != nil {
			return err
		} else {
			position = p
		}
	}

	parsed := false
	if len(req.URL.Query()["parsed"]) > 0 {
		parsed = true
	}

	_ = parsed

	log.Printf("(%s) stream position=%d", key, position)

	sm.lock.RLock()

	defer sm.lock.RUnlock()

	stream := sm.streams[key]
	if stream == nil {
		res.WriteHeader(http.StatusNotFound)
		return nil
	}

	stream.lock.Lock()

	defer stream.lock.Unlock()

	dropped := int64(0)
	buffered := stream.buffer.Bytes()
	data := buffered
	if position > stream.written {
		// Way out of bounds, so we can just treat this as fresh from the beginning.
		log.Printf("(%s) written=%d position=%d overflow", key, stream.written, position)
	} else if stream.written <= stream.bufferSize {
		// Buffer is still filling up, so we're going to return from position to the end of the buffer.
		log.Printf("(%s) written=%d position=%d filling", key, stream.written, position)
		data = buffered[position:]
	} else {
		// Buffer is full so we need to adjust position based on that, but the buffer may not have all we need.
		remaining := stream.written - position
		if remaining > stream.bufferSize {
			// Buffer has less than we need, so we return the whole thing and tell the client.
			log.Printf("(%s) written=%d position=%d remaining=%d dropped logs", key, stream.written, position, remaining)
			dropped = remaining - stream.bufferSize
		} else {
			// Most common scenario, we return the remaining bytes from the end of the buffer.
			data = buffered[stream.bufferSize-remaining:]
			log.Printf("(%s) written=%d position=%d remaining=%d", key, stream.written, position, remaining)
		}
	}

	setupResponse(res, req)

	res.Header().Set("Hpn-Dropped", fmt.Sprintf("%d", dropped))
	res.Header().Set("Hpn-More-Url", fmt.Sprintf("/v1/streams/%s?position=%d", key, stream.written))
	res.WriteHeader(http.StatusOK)
	res.Write(data)

	return nil
}

func web(sm *StreamManager) error {
	mux := mux.NewRouter()
	mux.HandleFunc("/v1/streams", middleware(sm, streamsHandler))
	mux.HandleFunc("/v1/streams/{key}", middleware(sm, streamHandler))

	s := &http.Server{
		Addr:    ":8580",
		Handler: mux,
	}

	log.Printf("[http] listening on %v", 8580)

	s.ListenAndServe()

	return nil
}

func main() {
	sm := NewStreamManager()

	defer sm.Close()

	go listen(sm)

	web(sm)
}
