package main

import (
	"fmt"

	"github.com/go_exercise/gopl/chap2/popcount"
)

func main() {
	fmt.Println(popcount.PopCountBySingle(17777777777777777))
	fmt.Println(popcount.PopCountByLoop(17777777777777777))
	fmt.Println(popcount.PopCountByShift(17777777777777777))
}
