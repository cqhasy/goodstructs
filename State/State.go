package State

type ElevatorState interface {
	HandleEvent(ctx ElevatorContext, event ElevatorEvent)
	Name() string
}
