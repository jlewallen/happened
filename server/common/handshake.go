package common

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
)

const (
	Prefix = "HPN"
)

type Handshake struct {
	Name       string `json:"name"`
	BufferSize int64  `json:"bufferSize"`
}

func (h *Handshake) Encode() ([]byte, error) {
	jsonBytes, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}

	data := make([]byte, 0)

	data = append(data, []byte(fmt.Sprintf("HPN:%d", len(jsonBytes)))...)

	data = append(data, jsonBytes...)

	return data, nil
}

func TryParseHandshake(r io.Reader) (reader io.Reader, h *Handshake, err error) {
	var buffer bytes.Buffer

	ourReader := bufio.NewReader(io.TeeReader(r, &buffer))

	headerCheck := make([]byte, 32)
	if _, err := io.ReadFull(ourReader, headerCheck); err != nil {
		return nil, nil, err
	}

	re := regexp.MustCompile(fmt.Sprintf("%s:(\\d+)", Prefix))

	m := re.FindAllStringSubmatch(string(headerCheck), -1)
	if len(m) > 0 {
		skipBytes := int64(len(m[0][0]))

		jsonReader := io.MultiReader(bytes.NewReader(headerCheck[skipBytes:]), ourReader)

		h = &Handshake{}

		decoder := json.NewDecoder(jsonReader)
		if err := decoder.Decode(h); err != nil {
			return nil, nil, err
		}

		reader = io.MultiReader(bytes.NewReader(buffer.Bytes()), r)

		skipBytes += decoder.InputOffset()
		log.Printf("skip %d bytes for handshake", skipBytes)
		io.CopyN(ioutil.Discard, reader, skipBytes)
	} else {
		reader = io.MultiReader(bytes.NewReader(headerCheck), r)
	}

	return
}
