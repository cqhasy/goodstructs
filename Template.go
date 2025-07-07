package main

import "fmt"

//和外观模式有点像啊

// 模板接口（由子类实现具体行为）
type AbstractClass interface {
	Step1()
	Step2()
	Step3()

	// 钩子方法：可选实现
	SkipStep2() bool
}

// 模板结构体，控制算法骨架
type BaseTemplate struct {
	impl AbstractClass
}

func (b *BaseTemplate) TemplateMethod() {
	b.impl.Step1()

	// 钩子控制是否执行 Step2
	if !b.impl.SkipStep2() {
		b.impl.Step2()
	}

	b.impl.Step3()
}

// DefaultHook 提供默认钩子实现
type DefaultHook struct{}

// 默认不跳过 Step2
func (h *DefaultHook) SkipStep2() bool {
	return false
}

// ConcreteClass1：不跳过 Step2
type ConcreteClass1 struct {
	BaseTemplate
	DefaultHook // 使用默认钩子
}

func (c *ConcreteClass1) Step1() {
	fmt.Println("[ConcreteClass1] Step 1")
}
func (c *ConcreteClass1) Step2() {
	fmt.Println("[ConcreteClass1] Step 2")
}
func (c *ConcreteClass1) Step3() {
	fmt.Println("[ConcreteClass1] Step 3")
}

// ConcreteClass2：跳过 Step2（自定义钩子）
type ConcreteClass2 struct {
	BaseTemplate
}

func (c *ConcreteClass2) Step1() {
	fmt.Println("[ConcreteClass2] Step 1")
}
func (c *ConcreteClass2) Step2() {
	fmt.Println("[ConcreteClass2] Step 2")
}
func (c *ConcreteClass2) Step3() {
	fmt.Println("[ConcreteClass2] Step 3")
}
func (c *ConcreteClass2) SkipStep2() bool {
	return true
}
func main() {
	fmt.Println("===> Class 1（执行全部步骤）")
	c1 := &ConcreteClass1{}
	c1.BaseTemplate.impl = c1
	c1.TemplateMethod()

	fmt.Println("\n===> Class 2（跳过 Step2）")
	c2 := &ConcreteClass2{}
	c2.BaseTemplate.impl = c2
	c2.TemplateMethod()
}
