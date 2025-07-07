package main

import "fmt"

// 定义泛型策略接口
type MathOperation[T any] interface {
	Operate(a, b T) T
}
type Addition[T int | float64] struct{}

func (Addition[T]) Operate(a, b T) T {
	return a + b
}

type Subtraction[T int | float64] struct{}

func (Subtraction[T]) Operate(a, b T) T {
	return a - b
}

type Multiplication[T int | float64] struct{}

func (Multiplication[T]) Operate(a, b T) T {
	return a * b
}

// 策略注册表
type OperationRegistry[T int | float64] struct {
	registry map[string]MathOperation[T]
}

func NewOperationRegistry[T int | float64]() *OperationRegistry[T] {
	return &OperationRegistry[T]{registry: make(map[string]MathOperation[T])}
}

func (r *OperationRegistry[T]) Register(name string, op MathOperation[T]) {
	r.registry[name] = op
}

func (r *OperationRegistry[T]) Get(name string) (MathOperation[T], bool) {
	op, ok := r.registry[name]
	return op, ok
}

// 上下文 Calculator
type Calculator[T int | float64] struct {
	op MathOperation[T]
}

func (c *Calculator[T]) SetOperation(op MathOperation[T]) {
	c.op = op
}

func (c *Calculator[T]) PerformOperation(a, b T) T {
	if c.op == nil {
		panic("no operation set")
	}
	return c.op.Operate(a, b)
}
func main() {
	registry := NewOperationRegistry[int]()
	registry.Register("add", Addition[int]{})
	registry.Register("sub", Subtraction[int]{})
	registry.Register("mul", Multiplication[int]{})

	calc := &Calculator[int]{}

	// 假设用户选择的策略为 "sub"
	operationName := "sub"
	if op, ok := registry.Get(operationName); ok {
		calc.SetOperation(op)
		result := calc.PerformOperation(10, 3)
		fmt.Printf("Result of '%s': %d\n", operationName, result)
	} else {
		fmt.Println("Unknown operation:", operationName)
	}
}
