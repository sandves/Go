package main

import (
	"fmt"
	"time"
)

type Pair struct {
	Key, Value string
}

func main() {
	pair := Pair{"foo", "bar"}
	pair2 := Pair{"hello", "world"}
	pair3 := Pair{"go", "goroutine"}

	var c chan string = make(chan string)

	go InsertAndLookup(pair, c, "1")
	go InsertAndLookup(pair2, c, "2")
	go InsertAndLookup(pair3, c, "3")

	for {
		select {
		case msg := <- c:
			fmt.Println("Insert and store", msg)
		case <- time.After(time.Second * 2):
			fmt.Println("Timeout")
			return
		}
	}
}

func InsertAndLookup(pair Pair, c chan string, client string) {
	var success bool
	err := Insert(pair, &success)
	checkError(err)

	var value string
	err = Lookup(pair.Key, &value)
	checkError(err)

	if success == true && value == "found" {
		c <- "succeeded for client " + client
	} else {
		c <- "failed for client " + client
	}
}

func Insert(pair Pair, success *bool) error {
	*success = true
	return nil
}

func Lookup(key string, value *string) error {
	*value = "found"
	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error")
	}
}
