package main

import (
	"flag"
	"fmt"
	"github.com/sandves/zaplab/chzap"
	"github.com/sandves/zaplab/ztorage"
	"net"
	"os"
	"os/signal"
	"runtime/pprof"
	"strings"
	"time"
)

type ZapServer struct {
	Zaps *ztorage.Zaps
}

var zaps *ZapServer

func main() {

	udpAddr, err := net.ResolveUDPAddr("udp", "224.0.1.130:10000")
	checkError(err)
	sock, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	checkError(err)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":12110")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	zaps = &ZapServer{Zaps: ztorage.NewZapStore()}

	go zaps.handleZaps(sock)
	go handleClient(listener)

	writeMemProfifle()
}

func (zs *ZapServer) handleZaps(conn *net.UDPConn) {
	for {
		var buf [1024]byte
		n, _, err := conn.ReadFromUDP(buf[0:])
		checkError(err)
		str := string(buf[:n])
		strSlice := strings.Split(str, ", ")
		if len(strSlice) == 5 {
			var channelZap *chzap.ChZap = chzap.NewChZap(str)
			zs.Zaps.StoreZap(*channelZap)
		}
	}
}

func handleClient(listener *net.TCPListener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go Subscribe(conn)
	}
}

func (zs *ZapServer) stats() string {
	topTen := zs.zaps.TopTenChannels()
	var topTenStr string
	for i := range topTen {
		topTenStr += fmt.Sprintf("Channel %d: %s\n", (i + 1), topTen[i])
	}
	avgZapDur := zs.Zaps.AverageZapDuration().String()
	topTenStr += fmt.Sprintf("\nAverage zap duration: %s", avgZapDur)
	return topTenStr
}

func Subscribe(conn net.Conn) {
	stats := zaps.stats()
	for _ = range time.Tick(1 * time.Second) {
		_, err := conn.Write([]byte(stats))
		if err != nil {
			conn.Close()
			break
		}
	}
}

//if the memprofile flag was specified, write a heap profile to file
func writeMemProfifle() {
	var memprofile = flag.String("memprofile", "", "write memory profile to this file")
	flag.Parse()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	<-signalChan
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		checkError(err)
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
