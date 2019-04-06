package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Name string
	Body string
	Time int64
}

type UnmatchMessage struct {
	Name string
	Boy  string
	Tim  int64
}

func main() {
	message := Message{"Tom", "Hello", 1294706395881547000}
	b, err := json.Marshal(message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Marshal!")
		os.Exit(1)
	}

	fmt.Printf("%s", b)

	var output Message
	decodeErr := json.Unmarshal(b, &output)
	if decodeErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to Unmarshal json data!err:%s", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", output)

	var unmatchedOutput UnmatchMessage
	message1 := []byte{`{"Name":"Tom","Body":"Hello","Time":1294706395881547000}`}
	decodeErr1 := json.Unmarshal(b, &unmatchedOutput)
	if decodeErr1 != nil {
		fmt.Fprintf(os.Stderr, "Failed to unmarshal json data! err:", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", unmatchedOutput)
}
