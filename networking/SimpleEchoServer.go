/* SimpleEchoServer
 */
 package main

 import (
	 "net"
	 "os"
	 "fmt"
 )

 func main() {
	 service := ":1201"
	 udpAddr, err := net.ResolveUDPAddr("udp", service)
	 checkError(err)

	 conn, err := net.ListenUDP("udp", udpAddr)
	 checkError(err)

	 for {
		 handleClient(conn)
	 }
 }

 func handleClient(conn *net.UDPConn) {
	 var buf[512]byte
	 n, addr, err := conn.ReadFromUDP(buf[0:])
	 if err != nil {
	 return
	 }
	 fmt.Println(string(buf[:0]))
	 _, err2 := conn.WriteToUDP(buf[0:n], addr)
	 if err2 != nil {
		 return
	 }
 }

 func checkError(err error) {
	 if err != nil {
		 fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		 os.Exit(1)
	 }
 }
