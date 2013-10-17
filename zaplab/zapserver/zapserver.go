package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"github.com/sandves/zaplab/ztorage"
	"github.com/sandves/zaplab/chzap"
)

var zaps *ztorage.Zaps

func main() {
	addr, err := net.ResolveUDPAddr("udp", "224.0.1.130:10000")
	checkError(err)
	sock, err := net.ListenMulticastUDP("udp", nil, addr)
	checkError(err)
	zaps = ztorage.NewZapStore()
	go computeViewers("NRK1")
	go computeViewers("TV2 Norge")
	go computeZaps()
	for {
		 handleClient(sock)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [1024]byte
	n, _, err := conn.ReadFromUDP(buf[0:])
	checkError(err)
	str := string(buf[:n])
	strSlice := strings.Split(str, ", ")
	if len(strSlice) == 5 {
		var channelZap *chzap.ChZap = chzap.NewChZap(str)
		zaps.StoreZap(*channelZap)
	}
}

func computeViewers(chName string) {
	for _ = range time.Tick(1 * time.Second) {
		numberOfViewers := zaps.ComputeViewers(chName)
		fmt.Printf("%s: %d\n", chName, numberOfViewers)
	}
}

func computeZaps() {
	for _ = range time.Tick(5 * time.Second) {
		numberOfZaps := len(*zaps)
		fmt.Printf("Total number of zaps: %d\n", numberOfZaps)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
