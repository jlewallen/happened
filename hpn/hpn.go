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
	Echo    bool
}

func applyEcho(o *options, writer io.Writer) io.Writer {
	if o.Echo {
		return io.MultiWriter(writer, os.Stdout)
	}
	return writer
}

func pipeReader(o *options, hs *common.Handshake, r io.Reader) error {
	dialing := fmt.Sprintf("%s:%d", o.Address, o.Port)
	conn, err := net.Dial("tcp", dialing)
	if err != nil {
		return err
	}

	defer conn.Close()

	encoded, err := hs.Encode()
	if err != nil {
		return err
	}

	if _, err := conn.Write(encoded); err != nil {
		return err
	}

	reader := bufio.NewReader(r)

	writing := applyEcho(o, conn)

	if _, err := io.Copy(writing, reader); err != nil {
		return err
	}

	return nil
}

func pipeFile(o *options, hs *common.Handshake, path string) error {
	r, err := os.Open(path)
	if err != nil {
		return err
	}

	defer r.Close()

	return pipeReader(o, hs, r)
}

func main() {
	o := &options{}

	flag.StringVar(&o.Name, "name", "", "stream name")
	flag.StringVar(&o.Address, "address", "127.0.0.1", "happened server address")
	flag.IntVar(&o.Port, "port", 2570, "happened server port")
	flag.BoolVar(&o.Retry, "retry", true, "retry connection attempts")
	flag.BoolVar(&o.Echo, "echo", false, "echo received stream to standard out")

	flag.Parse()

	if flag.NArg() > 0 {
		for _, arg := range flag.Args() {
			if _, err := os.Stat(arg); !os.IsNotExist(err) {
				hs := &common.Handshake{
					Name: arg,
				}

				if err := pipeFile(o, hs, arg); err != nil {
					panic(err)
				}
			} else {
				log.Printf("not a file: %v", arg)
			}
		}
	} else {
		for {
			hs := &common.Handshake{
				Name: o.Name,
			}

			if err := pipeReader(o, hs, os.Stdin); err != nil {
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
}
