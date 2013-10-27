package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveIPAddr("tcp4", ":12110")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	go SubscribeAndPrintStatistics(client, 1, reply)

	var str string
	for {
		select {
		case str = <-reply:
			fmt.Println(str)
		}
	}
}

func SubscribeAndPrintStatistics(conn *net.Conn) {
	var stats [512]byte
	_, err := conn.Read(stats[0:])
	checkError(err)
	fmt.Println(stats)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
