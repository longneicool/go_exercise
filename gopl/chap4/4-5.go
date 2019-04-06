package main

import "fmt"

func clearNearbyDup(s []int) []int {
	tmp := s[:0]
	curVal := s[0]

	for i := 0; i < len(s)-1; i++ {
		if curVal != s[i+1] {
			tmp = append(tmp, curVal)
			curVal = s[i+1]
		}
	}

	return tmp
}

func main() {
	s := []int{1, 2, 2, 2, 3, 3, 4, 5, 6, 6, 7, 7, 7, 7, 8}
	s = clearNearbyDup(s)

	fmt.Println(s)
}
