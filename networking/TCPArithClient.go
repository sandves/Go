pakcage main

import (
	"net/rpc"
	"fmt"
	"log"
	"os"
)

type Pair struct {
	Key, Value string
}

func main() {
	if len(os.Args) !=2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}
	service := os.Args[1]

	client, err := rpc.Dial("tcp" service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//Synchronous call
	pair := Pair{"hello", "world"}
	var success bool
	err = client.Call("KeyValue.Insert", pair, &reply)
	if err != nil {
		log.Fatal("KeyValue error:", err)
	}
	fmt.Printf("Insert of key %s and value %s was successful: %b", pair.Key, pair.Value, success)

	var insertedValue string
	err = client.Call("KeyValue.Lookup", pair.Key, &insertedValue)
	if err != nil {
		log.Fatal("KeyValue error: ", err)
	}
	if insertedValue == pair.Value {
		fmt.Printf("Confirmation that the pair was stored")
	} else {
		fmt.Printf("The value %s was not found with key %s", pair.Value, pair.Key)
	}
}
