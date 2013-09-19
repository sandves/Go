package main

import "fmt"
import "net"

func main() {
	serverAddr, err := net.ResolveTCPAddr("tcp", "152.94.1.255:12110")
	fmt.Println(serverAddr)
	fmt.Println(err)
}
