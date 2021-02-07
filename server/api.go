package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

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
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Hpn-More-Url, Hpn-Dropped, Hpn-Summary")
	res.Header().Set("Access-Control-Expose-Headers", "Content-Type, Hpn-More-Url, Hpn-Dropped, Hpn-Summary")
}

type StreamResponse struct {
	Key     string `json:"key"`
	URL     string `json:"url"`
	Written int64  `json:"written"`
}

type StreamsResponse struct {
	Streams []*StreamResponse `json:"streams"`
}

type StreamsSummary struct {
	Streams int `json:"streams"`
}

type TailResponse struct {
	Blocks   []string        `json:"blocks"`
	Messages []string        `json:"messages"`
	MoreURL  string          `json:"more"`
	Dropped  bool            `json:"dropped"`
	Summary  *StreamsSummary `json:"summary"`
}

func streamsIndexHandler(sm *StreamManager, res http.ResponseWriter, req *http.Request) error {
	sm.lock.RLock()

	defer sm.lock.RUnlock()

	streams := make([]*StreamResponse, 0)

	for key, stream := range sm.streams {
		streams = append(streams, &StreamResponse{
			Key:     key,
			URL:     fmt.Sprintf("/v1/streams/%s", key),
			Written: stream.source.Written(),
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
	ctx := context.Background()
	vars := mux.Vars(req)
	key := vars["key"]

	position := Position{}

	maybePosition := req.URL.Query()["position"]
	if len(maybePosition) == 1 {
		position.Encoded = &maybePosition[0]
	}

	log.Printf("(%s) stream position=%v", key, position)

	stream := sm.Lookup(key)
	if stream == nil {
		res.WriteHeader(http.StatusNotFound)
		return nil
	}

	response, err := stream.source.Tail(ctx, position)
	if err != nil {
		log.Printf("error: %v", err)
		res.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	response.Summary = &StreamsSummary{
		Streams: len(sm.streams),
	}

	response.MoreURL = fmt.Sprintf("/v1/streams/%s?%s", key, response.MoreURL)

	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("error: %v", err)
		res.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	setupResponse(res, req)

	res.Header().Set("Content-Type", "application/json")

	res.Write(data)

	log.Printf("done")

	return nil
}

func web(sm *StreamManager) error {
	mux := mux.NewRouter()
	mux.HandleFunc("/v1/streams", middleware(sm, streamsIndexHandler))
	mux.HandleFunc("/v1/streams/{key}", middleware(sm, streamHandler))

	s := &http.Server{
		Addr:    ":8580",
		Handler: mux,
	}

	log.Printf("[http] listening on %v", 8580)

	s.ListenAndServe()

	return nil
}

func listen(sm *StreamManager) error {
	ctx := context.Background()

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

		source, err := NewTcpTextSource(conn)
		if err != nil {
			return err
		}

		stream, err := sm.AddStream(source)
		if err != nil {
			return err
		}

		go func() {
			source.handle(ctx)

			if err := sm.RemoveStream(stream); err != nil {
				log.Printf("remove error: %v", err)
			}
		}()
	}

	return nil
}
