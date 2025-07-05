package main

import "fmt"

type Drawer interface {
	Draw()
}
type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing a circle")
}

type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("Drawing a rectangle")
}

type ShapeFactory interface {
	CreateShape() Drawer
}
type CircleFactory struct{}

func (cf *CircleFactory) CreateShape() Drawer {
	return &Circle{}
}

type RectangleFactory struct{}

func (rf *RectangleFactory) CreateShape() Drawer {
	return &Rectangle{}
}
func drawer(f ShapeFactory) {
	d := f.CreateShape()
	d.Draw()
}
func main() {
	drawer(&CircleFactory{})
	drawer(&RectangleFactory{})
}
