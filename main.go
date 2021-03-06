package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"flag"
)

var maddr = flag.String("a", "", "Multicast address to broadcast to")
var intf = flag.String("i", "", "Multicast interface to listen on")

const (
	maxDatagramSize = 8192
)

func main() {
	flag.Parse()
	serveMulticastUDP(*maddr, msgHandler)
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println("\n" + hex.Dump(b[:n]))
}

func serveMulticastUDP(a string, h func(*net.UDPAddr, int, []byte)) {
	fmt.Println("Listening to:", a)
	addr, err := net.ResolveUDPAddr("udp", a)
	
	if err != nil {
		log.Fatal(err)
	}
	iface, err := net.InterfaceByName( *intf )
	if err != nil {
		log.Println("Error using interface:", *intf, "Error:", err)
		iface = nil
	}
	l, err := net.ListenMulticastUDP("udp", iface, addr)
	l.SetReadBuffer(maxDatagramSize)
	b := make([]byte, maxDatagramSize)
	for {
		bytesRead, src, err := l.ReadFromUDP(b)
		if (err != nil) {
			log.Fatal("ReadFromUDP failed:", err)
		}
		if (bytesRead == 0) { continue; }
		h(src, bytesRead, b)
	}
}