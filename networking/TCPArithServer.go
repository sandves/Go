/**
* TCPArithServer
*/

package main

import (
	"fmt"
	"net/rpc"
	"net"
	"os"
	"sync"
)

type Pair struct {
	Key, Value string
}

type KeyValue struct {
	lock *sync.Mutex
	Map map[string]string
}

func (kv *KeyValue) Insert(input Pari, reply *bool) error {
	kv.lock.Lock()
	kv.Map[input.Key]=input.Value
	kv.lock.Unlock()
	*reply = true
	fmt.Printf("inserted Key %v´n", input.Key, input.Value)
	return nil
}

func (kv* KeyValue) Lookup(input string, reply *string) error {
	kv.lock.Lock()
	if v, ok := kv.Map[input]; ok {
		*reply = v
		fmt.Println(v)
	} else {
		*reply = "not found"
	}
	kv.lock.Unlock()
	return nil
}

func main() {
	kv := &KeyValue{new(sync.Mutex), make(map[string]string)}
	rpc.Register(kv)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":12110")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	//This works:
	rcp.Accept(listener)

	/* and so does this:
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}*/
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
