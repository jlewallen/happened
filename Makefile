default: build build/hpn build/hserver

build:
	mkdir build

build/hpn: hpn/*.go
	go build -o $@ $^

build/hserver: server/*.go
	go build -o $@ $^

clean:
	rm -rf build
