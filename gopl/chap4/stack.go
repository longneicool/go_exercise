package main

func main() {
	stack := make([]string, 0)

	stack = append(stack, "s")
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

}
