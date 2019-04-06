package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	SHA256 = "SHA256"
	SHA384 = "SHA384"
	SHA512 = "SHA512"
)

func main() {
	hashStr := flag.String("c", "", "The Hash character")
	hashAlgorithm := flag.String("t", SHA256, "The Hash Algorithm")

	flag.Parse()

	fmt.Println(*hashStr, *hashAlgorithm)

	var hashArray [32]byte
	switch *hashAlgorithm {
	case SHA256:
		hashArray = sha256.Sum256([]byte(*hashStr))
	case SHA512:
		hashArray = sha512.Sum512_256([]byte(*hashStr))
	default:
		fmt.Fprintf(os.Stderr, "The Hash Algorithm is incorrect!")
		os.Exit(1)
	}

	fmt.Printf("%x\n", hashArray)

}
