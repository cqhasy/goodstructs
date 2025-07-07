package main

import "fmt"

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
type IterableCollection[T any] interface {
	CreateIterator() Iterator[T]
}
type ConcreteCollection[T any] struct {
	items []T
}

func NewConcreteCollection[T any]() *ConcreteCollection[T] {
	return &ConcreteCollection[T]{}
}

func (c *ConcreteCollection[T]) AddItem(item T) {
	c.items = append(c.items, item)
}

func (c *ConcreteCollection[T]) CreateIterator() Iterator[T] {
	return &ConcreteIterator[T]{items: c.items}
}

type ConcreteIterator[T any] struct {
	items    []T
	position int
}

func (it *ConcreteIterator[T]) HasNext() bool {
	return it.position < len(it.items)
}

func (it *ConcreteIterator[T]) Next() T {
	if !it.HasNext() {
		panic("No more elements")
	}
	item := it.items[it.position]
	it.position++
	return item
}
func main() {
	collection := NewConcreteCollection[string]()
	collection.AddItem("Item 1")
	collection.AddItem("Item 2")
	collection.AddItem("Item 3")

	iterator := collection.CreateIterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
