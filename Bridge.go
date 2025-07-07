package main

import "fmt"

//桥接模式，接口作为字段

type Color interface {
	printColor()
}
type red struct{}
type green struct{}

func (r *red) printColor() {
	fmt.Println("red")
}
func (g *green) printColor() {
	fmt.Println("green")
}

type object interface {
	printArt()
}
type circle struct {
	color Color
}

func newCircle(color Color) *circle {
	return &circle{color}
}
func (c *circle) printArt() {
	fmt.Print("这是一个圆，颜色为")
	c.color.printColor()
}

type square struct {
	color Color
}

func newSquare(color Color) *square {
	return &square{color}
}
func (s *square) printArt() {
	fmt.Print("这是一个方块，颜色为")
	s.color.printColor()
}
