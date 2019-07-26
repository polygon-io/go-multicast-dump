package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"flag"
)

var maddr = flag.String("a", "", "Multicast address to broadcast to")

const (
	maxDatagramSize = 8192
)

func main() {
	flag.Parse()
	serveMulticastUDP(*maddr, msgHandler)
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}

func serveMulticastUDP(a string, h func(*net.UDPAddr, int, []byte)) {
	fmt.Println("Listening to:", a)
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	l.SetReadBuffer(maxDatagramSize)
	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		h(src, n, b)
	}
}