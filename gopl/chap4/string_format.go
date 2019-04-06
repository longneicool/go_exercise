package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{x: 1, y: 2}
	fmt.Printf("%%v: %v\n", p)   // {1 2}
	fmt.Printf("%%+v: %+v\n", p) //{x:1 y:2} will include the field names of the struct
	fmt.Printf("%%#v: %#v\n", p) //main.point{x:1 y:2}

	fmt.Printf("%%T: %T\n", p)    // main.point %T print the type of a value
	fmt.Printf("%%T: %T\n", true) // bool
	fmt.Printf("%%t: %t\n", p)

	fmt.Printf("%%d: %d\n", 123) //123
	fmt.Printf("%%b: %b \n", 14) // 1110 print binary representation
	fmt.Printf("%%c: %c\n", 33)  // ! print the character
	fmt.Printf("%%x: %x\n", 44)  // 2c provides the hex encoding

	fmt.Printf("%%f: %f\n", 78.9)

	fmt.Printf("%%e: %e\n", 123400000.0) // 只能ｆｏｒｍａｔ float.不能ｆｏｒｍａｔ　ｉｎｔ类型
	fmt.Printf("%%E: %E\n", 123400000.0)

	fmt.Printf("%%s: %s\n", "\"string\"")
	fmt.Printf("%%q: %q\n", "\"string\"") // "\"string\"" To double-quote strings as in Go source
	fmt.Printf("%%x: %x\n", "hex this")   // 其实就是将其中的character转换为Ｂｙｔｅ

	fmt.Printf("%%p: %p\n", &p) // print the pointer.

	fmt.Printf("%%6d: |%6d|\n", 12)           //|    12| right justified.
	fmt.Printf("|%6.2f|%-6.2f|\n", 1.2, 3.45) //|  1.20|3.45  | right justified

	fmt.Printf("|%6s|\n", "foo") // |   foo|

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// you can use fotmat+print to io.Writers other than os.Stdout with Fprintf
	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
