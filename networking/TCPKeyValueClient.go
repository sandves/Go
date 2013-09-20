package main

import (
	"fmt"
	"net/rpc"
	"os"
	"strconv"
)

type Pair struct {
	Key, Value string
}

func main() {
	client, err := rpc.Dial("tcp", ":12110")
	checkError(err)

	numberOfGoroutines := 5
	callReturned := make(chan bool)

	for i := 0; i < numberOfGoroutines; i++ {
		go func(i int) {
			InsertAndLookup(client, i)
			callReturned <- true
		}(i)
	}

	for i := 0; i < numberOfGoroutines; i++ {
		<-callReturned
	}
	fmt.Println("Finished!")
}

func InsertAndLookup(client *rpc.Client, pairNumber int) {
	var i string = strconv.Itoa(pairNumber)
	pair := Pair{"key" + i, "world" + i}

	var success bool
	err := *client.Call("KeyValue.Insert", pair, &success)
	checkError(err)
	fmt.Printf("Insert of key %s and value %s was successful: %b\n", pair.Key, pair.Value, success)

	var insertedValue string
	err = *client.Call("KeyValue.Lookup", pair.Key, &insertedValue)
	checkError(err)
	if insertedValue == pair.Value {
		fmt.Println("Confirmation that the pair was stored")
	} else {
		fmt.Printf("The value %s was not found with key %s\n", pair.Value, pair.Key)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
