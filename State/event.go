package State

type ElevatorEvent int

const (
	EventOpenDoors ElevatorEvent = iota
	EventCloseDoors
	EventMove
	EventStop
)
