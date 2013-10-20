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

var zaps *ztorage.Zaps

func main() {

	addr, err := net.ResolveUDPAddr("udp", "224.0.1.130:10000")
	checkError(err)
	sock, err := net.ListenMulticastUDP("udp", nil, addr)
	checkError(err)
	zaps = ztorage.NewZapStore()

	go handleClient(sock)
	go topTenChannels()
	//go computeViewers("NRK1")
	//go computeViewers("TV2 Norge")
	//go computeZaps()

	writeMemProfifle()
}

func handleClient(conn *net.UDPConn) {
	for {
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
}

func topTenChannels() {
	for _ = range time.Tick(1 * time.Second) {
		fmt.Println()
		topTen := zaps.TopTenChannels()
		for i := range topTen {
			fmt.Printf("Channel %d: %s\n", (i + 1), topTen[i])
		}
		fmt.Println()
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

/*func computeViewers(chName string) {
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
}*/

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
