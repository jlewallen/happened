package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/jlewallen/happened/server/common"
)

type options struct {
	Name    string
	Address string
	Port    int
	Retry   bool
}

func pipe(o *options) error {
	reader := bufio.NewReader(os.Stdin)

	dialing := fmt.Sprintf("%s:%d", o.Address, o.Port)
	conn, err := net.Dial("tcp", dialing)
	if err != nil {
		return err
	}

	defer conn.Close()

	hs := common.Handshake{
		Name: o.Name,
	}

	encoded, err := hs.Encode()
	if err != nil {
		return err
	}

	if _, err := conn.Write(encoded); err != nil {
		return err
	}

	if _, err := io.Copy(conn, reader); err != nil {
		return err
	}

	return nil
}

func main() {
	o := &options{}

	flag.StringVar(&o.Name, "name", "", "stream name")
	flag.StringVar(&o.Address, "address", "127.0.0.1", "happened server address")
	flag.IntVar(&o.Port, "port", 2570, "happened server port")
	flag.BoolVar(&o.Retry, "retry", true, "retry connection attempts")

	flag.Parse()

	for {
		if err := pipe(o); err != nil {
			log.Printf("error %v", err)
		} else {
			break
		}

		if !o.Retry {
			break
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
