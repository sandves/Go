//use: sortmsgs [file]
package main

import (
	"fmt"
	"msg"
	"os"
	"time"
	"encoding/gob"
	"io"
)

var (
	connectingChan = make(chan interface{}, 5) // Channel from demarshaler to sort
	msgChan        = make(chan msg.StrMsg, 5)  // Message channel
	errChan        = make(chan msg.ErrMsg, 5)  // Error channel
)

func main() {
	go handleMsgs(msgChan)
	go handleErrors(errChan)
	go sort(connectingChan, msgChan, errChan)
	unmarshal(os.Args[1], connectingChan)
	time.Sleep(5 * time.Second)
}

func unmarshal(filepath string, outchan chan<- interface{}) {
	// Retrive every stored element by decoding into an interface{} until EOF is reached.
	// Send each message decoded on the outgoing channel.
	// When all messages are sent, close the outgoing channel.
	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		fmt.Printf("Error opening file...", err)
	}

	decoder := gob.NewDecoder(file)	
	
	var m interface{}
	for decoder.Decode(&m) != io.EOF
	
	if err != nil {
		fmt.Printf("Error decoding file", err)
	}
}

func sort(inchan <-chan interface{}, msgchan chan<- msg.StrMsg, errchan chan<- msg.ErrMsg) {
	// Receive messages from the inchan channel.
	// Forward messages to the msgchan or errchan according to type.
	// Tips: Use type assertion, e.g. switch x.(type) { case:...})
	// When all messages are received, close the msgchan and errchan channels.
	switch
}

func handleMsgs(inchan chan msg.StrMsg) {
	messages := make([]msg.StrMsg, 0)
	for {
		sm, ok := <-inchan
		if ok {
			messages = append(messages, sm)
			continue
		}
		if len(messages) > 0 {
			fmt.Println("Received messages:")
			for i, msg := range messages {
				fmt.Println(i, ":", msg)
			}
		}
		break
	}
}

func handleErrors(inchan chan msg.ErrMsg) {
	errors := make([]msg.ErrMsg, 0)
	for {
		em, ok := <-inchan
		if ok {
			errors = append(errors, em)
			continue
		}
		if len(errors) > 0 {
			fmt.Println("Received errors:")
			for i, err := range errors {
				fmt.Println(i, ":", err)
			}
		}
		break
	}
}
