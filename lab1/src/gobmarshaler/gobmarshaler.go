// Usage: gobMarshaler filepath [Msg/Err] Sender Content [Msg/Err] Sender Content ...
package main

import (
	"encoding/gob"
	"fmt"
	"msg"
	"os"
)

func main() {
	file, _ := os.Create(os.Args[1])
	defer file.Close()
	encoder := gob.NewEncoder(file)
	var m interface{}
	for i := 0; i*3 < len(os.Args)-4; i++ {
		switch {
		case os.Args[i*3+2] == "Msg" && i*3+3 < len(os.Args):
			m = msg.StrMsg{os.Args[i*3+3], os.Args[i*3+4]}
			fmt.Println("Message from " + os.Args[i*3+3] + " saying " + os.Args[i*3+4])
		case os.Args[i*3+2] == "Err" && i*3+3 < len(os.Args):
			m = msg.ErrMsg{os.Args[i*3+3], os.Args[i*3+4]}
			fmt.Println("Error from " + os.Args[i*3+3] + " saying " + os.Args[i*3+4])
		default:
			fmt.Println("Wrong input.\n Usage: gobMarshaler filepath [Msg/Err] Sender Content [Msg/Err] Sender Content ...")
			return
		}
		if err := encoder.Encode(&m); err != nil {
			fmt.Println("Error!", err)
			return
		}
	}
	fmt.Println("All messages stored")
}
