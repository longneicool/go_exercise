package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p Point) Add(q Point) Point {
	fmt.Println("Add")
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	fmt.Println("Sub")
	return Point{p.X - q.X, p.Y - q.Y}
}

func transfer(p Point, q Point, add bool) Point {
	var oper func(Point, Point) Point
	if add {
		oper = Point.Add
	} else {
		oper = Point.Sub
	}

	return oper(p, q)
}

func main() {
	p, q := Point{1, 1}, Point{2, 2}
	fmt.Printf("Add: %v\n", transfer(p, q, true))
	fmt.Printf("Sub: %v\n", transfer(p, q, false))
}
