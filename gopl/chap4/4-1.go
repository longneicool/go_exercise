package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/go_exercise/gopl/chap2/popcount"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("c1: %d c2: %d", popcount.PopCountFromByteSlice(c1[0:]), popcount.PopCountFromByteSlice(c2[0:]))
}
