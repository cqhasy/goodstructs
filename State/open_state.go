package State

type OpenState struct{}

func (s *OpenState) Name() string {
	return "OpenState"
}

func (s *OpenState) HandleEvent(ctx ElevatorContext, event ElevatorEvent) {
	switch event {
	case EventOpenDoors:
		ctx.Log("门已经打开。")
	case EventCloseDoors:
		ctx.Log("正在关门...")
		ctx.SetState(&CloseState{})
	case EventMove:
		ctx.Log("门开着不能移动。")
	case EventStop:
		ctx.Log("电梯在开门状态下已停止。")
	default:
		ctx.Log("未知操作。")
	}
}
