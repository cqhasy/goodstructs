package main

import "fmt"

type Computer struct {
	CPU     string
	RAM     string
	Storage string
}

// 建造者结构体
type ComputerBuilder struct {
	computer Computer
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{}
}

func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	b.computer.CPU = cpu
	return b
}
func (b *ComputerBuilder) SetRAM(ram string) *ComputerBuilder {
	b.computer.RAM = ram
	return b
}
func (b *ComputerBuilder) SetStorage(storage string) *ComputerBuilder {
	b.computer.Storage = storage
	return b
}

func (b *ComputerBuilder) Build() Computer {
	return b.computer
}

func main() {
	pc := NewComputerBuilder().
		SetCPU("Intel i9").
		SetRAM("32GB").
		SetStorage("1TB SSD").
		Build()
	fmt.Printf("%+v\n", pc)
}
