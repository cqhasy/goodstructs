package State

import "fmt"

type ElevatorContext interface {
	SetState(state ElevatorState)
	GetState() ElevatorState
	Log(msg string)
}

// Elevator 实现 ElevatorContext
type Elevator struct {
	state ElevatorState
}

func NewElevator() *Elevator {
	return &Elevator{
		state: &CloseState{}, // 初始状态为关门
	}
}

func (e *Elevator) Handle(event ElevatorEvent) {
	e.state.HandleEvent(e, event)
}

func (e *Elevator) SetState(state ElevatorState) {
	e.state = state
}

func (e *Elevator) GetState() ElevatorState {
	return e.state
}

func (e *Elevator) Log(msg string) {
	fmt.Println(msg)
}
