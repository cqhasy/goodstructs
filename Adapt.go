package main

import "fmt"

//本质就是外包了一层，用来新老接口的对接和复用

// LegacyRectangle 已存在的类，带有Display方法
type LegacyRectangle struct{}

func (r *LegacyRectangle) Display(x1, y1, x2, y2 int) {
	fmt.Printf("LegacyRectangle: Point1(%d, %d), Point2(%d, %d)\n", x1, y1, x2, y2)
}

// FShape 统一接口
type Shape interface {
	Draw(x, y, width, height int)
}

// RectangleAdapter 适配器，实现Shape接口
type RectangleAdapter struct {
	legacyRectangle *LegacyRectangle
}

func NewRectangleAdapter(lr *LegacyRectangle) *RectangleAdapter {
	return &RectangleAdapter{legacyRectangle: lr}
}

func (r *RectangleAdapter) Draw(x, y, width, height int) {
	x1 := x
	y1 := y
	x2 := x + width
	y2 := y + height
	r.legacyRectangle.Display(x1, y1, x2, y2)
}

func main() {
	legacyRectangle := &LegacyRectangle{}
	shapeAdapter := NewRectangleAdapter(legacyRectangle)

	shapeAdapter.Draw(10, 20, 50, 30)
}
