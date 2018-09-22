// Interface assignment https://www.udemy.com/go-the-complete-developers-guide/learn/v4/t/practice/9944/introduction

package main

import (
	"fmt"
)

func main() {

	s := square{
		sideLength: 8,
	}
	t := triangle{
		height: 10,
		base:   5,
	}

	s.printArea()
	t.printArea()
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	printArea()
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) printArea() {
	fmt.Println(s.getArea())
}

func (t triangle) printArea() {
	fmt.Println(t.getArea())
}
