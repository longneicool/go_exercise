package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(step int, s []int) {
	reverse(s[0:step])
	reverse(s[step:])
	reverse(s)
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	rotate(3, s)

	fmt.Println(s)
}
