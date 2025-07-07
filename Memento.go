package main

import "fmt"

//很有用

// Memento 表示备忘录
type Memento struct {
	state string
}

func NewMemento(state string) *Memento {
	return &Memento{state: state}
}

func (m *Memento) State() string {
	return m.state
}

// Originator 拥有状态
type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) CreateMemento() *Memento {
	return NewMemento(o.state)
}

func (o *Originator) RestoreMemento(m *Memento) {
	o.state = m.State()
}

// Caretaker 管理所有状态历史（支持多次撤销/重做）
type Caretaker struct {
	history    []*Memento
	current    int
	originator *Originator
}

func NewCaretaker(originator *Originator) *Caretaker {
	return &Caretaker{
		originator: originator,
		history:    []*Memento{},
		current:    -1,
	}
}

// Save 当前状态到历史
func (c *Caretaker) Save() {
	// 删除当前之后的所有 redo 历史
	c.history = c.history[:c.current+1]

	m := c.originator.CreateMemento()
	c.history = append(c.history, m)
	c.current++
}

// Undo 恢复上一个状态
func (c *Caretaker) Undo() {
	if c.current <= 0 {
		fmt.Println("⚠️ 不能再撤销了")
		return
	}
	c.current--
	c.originator.RestoreMemento(c.history[c.current])
}

// Redo 恢复下一个状态
func (c *Caretaker) Redo() {
	if c.current >= len(c.history)-1 {
		fmt.Println("⚠️ 没有更多可重做的操作")
		return
	}
	c.current++
	c.originator.RestoreMemento(c.history[c.current])
}
func main() {
	originator := &Originator{}
	caretaker := NewCaretaker(originator)

	originator.SetState("State A")
	caretaker.Save()

	originator.SetState("State B")
	caretaker.Save()

	originator.SetState("State C")
	caretaker.Save()

	fmt.Println("🟢 当前状态：", originator.GetState())

	// 多次撤销
	caretaker.Undo()
	fmt.Println("↩️ 撤销后：", originator.GetState())

	caretaker.Undo()
	fmt.Println("↩️ 再次撤销：", originator.GetState())

	// 重做
	caretaker.Redo()
	fmt.Println("↪️ 重做一次：", originator.GetState())
}
