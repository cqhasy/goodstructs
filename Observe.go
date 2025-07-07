package main

import (
	"fmt"
	"sync"
)

// Subject 接口
type Subject interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	SetState(state int)
	GetState() int
}

// Observer 接口
type Observer interface {
	Start() // 启动监听协程
	Stop()  // 停止监听协程
	Channel() chan int
	Name() string
}

// ConcreteSubject 实现：线程安全
type ConcreteSubject struct {
	observers map[Observer]struct{}
	state     int
	mutex     sync.Mutex
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make(map[Observer]struct{}),
	}
}

func (s *ConcreteSubject) AddObserver(observer Observer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.observers[observer] = struct{}{}
}

func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.observers, observer)
}

func (s *ConcreteSubject) notifyObservers() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for observer := range s.observers {
		observer.Channel() <- s.state // 通过 channel 异步通知
	}
}

func (s *ConcreteSubject) SetState(state int) {
	s.mutex.Lock()
	s.state = state
	s.mutex.Unlock()

	s.notifyObservers()
}

func (s *ConcreteSubject) GetState() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.state
}

// ConcreteObserver：使用 channel 接收通知
type ConcreteObserver struct {
	name string
	ch   chan int
	quit chan struct{}
	wg   sync.WaitGroup
}

func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{
		name: name,
		ch:   make(chan int),
		quit: make(chan struct{}),
	}
}

func (o *ConcreteObserver) Start() {
	o.wg.Add(1)
	go func() {
		defer o.wg.Done()
		for {
			select {
			case state := <-o.ch:
				fmt.Printf("%s 收到新状态: %d\n", o.name, state)
			case <-o.quit:
				fmt.Printf("%s 停止监听\n", o.name)
				return
			}
		}
	}()
}

func (o *ConcreteObserver) Stop() {
	close(o.quit)
	o.wg.Wait()
}

func (o *ConcreteObserver) Channel() chan int {
	return o.ch
}

func (o *ConcreteObserver) Name() string {
	return o.name
}

// main 演示
func main() {
	subject := NewConcreteSubject()

	observer1 := NewConcreteObserver("观察者1")
	observer2 := NewConcreteObserver("观察者2")

	observer1.Start()
	observer2.Start()

	subject.AddObserver(observer1)
	subject.AddObserver(observer2)

	subject.SetState(100)
	//time.Sleep(500 * time.Millisecond)

	//subject.RemoveObserver(observer1)
	defer observer1.Stop()

	subject.SetState(200)
	//time.Sleep(500 * time.Millisecond)

	observer2.Stop()
}
