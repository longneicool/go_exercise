package main

import (
	"fmt"
	"os"
)

func main() {

	for index, v := range os.Args {
		fmt.Println(index, v)
	}

}
