package main

import (
	"fmt"
)

// Coffee 接口，定义基础行为
type Coffee interface {
	Cost() float64
	Description() string
}

// SimpleCoffee 实现基本咖啡
type SimpleCoffee struct{}

func (s SimpleCoffee) Cost() float64 {
	return 2.0
}

func (s SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// MilkDecorator 装饰器：加牛奶
type MilkDecorator struct {
	base Coffee
}

func NewMilkDecorator(c Coffee) Coffee {
	return MilkDecorator{base: c}
}

func (m MilkDecorator) Cost() float64 {
	return m.base.Cost() + 1.0
}

func (m MilkDecorator) Description() string {
	return m.base.Description() + ", with Milk"
}

// SugarDecorator 装饰器：加糖
type SugarDecorator struct {
	base Coffee
}

func NewSugarDecorator(c Coffee) Coffee {
	return SugarDecorator{base: c}
}

func (s SugarDecorator) Cost() float64 {
	return s.base.Cost() + 0.5
}

func (s SugarDecorator) Description() string {
	return s.base.Description() + ", with Sugar"
}

// 示例入口
func main() {
	var coffee Coffee = SimpleCoffee{}
	fmt.Printf("Cost: $%.2f, Description: %s\n", coffee.Cost(), coffee.Description())

	coffee = NewMilkDecorator(coffee)
	fmt.Printf("Cost: $%.2f, Description: %s\n", coffee.Cost(), coffee.Description())

	coffee = NewSugarDecorator(coffee)
	fmt.Printf("Cost: $%.2f, Description: %s\n", coffee.Cost(), coffee.Description())
}
