default: build build/hpn build/hserver

build:
	mkdir build

build/hpn: hpn/hpn.go server/common/*.go
	go build -o $@ hpn/hpn.go

build/hserver: server/*.go server/common/*.go
	go build -o $@ server/*.go

clean:
	rm -rf build

install: default
	cp build/hpn ~/tools/bin
