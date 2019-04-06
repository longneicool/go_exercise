package main

import (
	"crypto/sha256"
	"fmt"
)
｀
func main() {
	// []byte("x")　其实就是将"x"看作是数组，然后转换为ｓｌｉｃｅ
	// 比如 s := []byte("xyz")　是其中分别为[120 121 122]的ｓｌｉｃｅ
	c1 := sha256.Sum256([]byte("x"))

	fmt.Printf("%x", c1)
}
