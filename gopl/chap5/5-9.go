package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	output := f("foo")
	return strings.Replace(s, "foo", output, -1)
}

func f1(s string) string {
	return "Hello World"
}

func main() {
	output := expand("foofoofoofooooofoofofofffooo", f1)
	fmt.Println(output)
}
