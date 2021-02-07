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
	Name string `json:"name"`
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
	var buf bytes.Buffer

	length := int64(32)

	headerReader := bufio.NewReader(io.TeeReader(io.LimitReader(r, length), &buf))

	data := make([]byte, length)
	if _, err := io.ReadFull(headerReader, data); err != nil {
		return nil, nil, err
	}

	re := regexp.MustCompile(fmt.Sprintf("%s:(\\d+)", Prefix))

	reader = io.MultiReader(bytes.NewReader(buf.Bytes()), r)

	m := re.FindAllStringSubmatch(string(data), -1)
	if len(m) > 0 {
		skipBytes := int64(len(m[0][0]))

		jsonReader := io.MultiReader(bytes.NewReader(buf.Bytes()[skipBytes:]), r)

		decoder := json.NewDecoder(jsonReader)

		h = &Handshake{}
		if err := decoder.Decode(h); err != nil {
			return nil, nil, err
		}

		skipBytes += decoder.InputOffset()

		log.Printf("skip %d bytes for handshake", skipBytes)

		io.CopyN(ioutil.Discard, reader, skipBytes)
	}

	return
}
