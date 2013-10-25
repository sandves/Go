package main

import (
	"fmt"
	"net/rpc"
	"os"
)

func main() {
	client, err := rpc.Dial("tcp", ":12110")
	checkError(err)

	reply := make(chan string)
	go SubscribeForStatistics(client, 1, reply)

	var str string
	for {
		select {
		case str = <-reply:
			fmt.Println(str)
		}
	}
}

func SubscribeForStatistics(client *rpc.Client, rate int, reply chan string) {

	err := client.Call("Zapserver.Subscribe", rate, &reply)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
