package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"github.com/sandves/zaplab/chzap"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "224.0.1.130:10000")
	checkError(err)
	sock, err := net.ListenMulticastUDP("udp", nil, addr)
	checkError(err)
		for {
		handleClient(sock)
		time.Sleep((time.Second))
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [1024]byte
	n, _, err := conn.ReadFromUDP(buf[0:])
	checkError(err)
	str := string(buf[:n])
	var channelZap *chzap.ChZap = chzap.NewChZap(str)
	fmt.Printf("%s\n", channelZap.String())
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
