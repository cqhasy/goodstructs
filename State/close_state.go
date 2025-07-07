package State

type CloseState struct{}

func (s *CloseState) Name() string {
	return "CloseState"
}

func (s *CloseState) HandleEvent(ctx ElevatorContext, event ElevatorEvent) {
	switch event {
	case EventOpenDoors:
		ctx.Log("正在开门...")
		ctx.SetState(&OpenState{})
	case EventCloseDoors:
		ctx.Log("门已经关闭。")
	case EventMove:
		ctx.Log("开始移动...")
		ctx.SetState(&MoveState{})
	case EventStop:
		ctx.Log("电梯停止。")
	default:
		ctx.Log("未知操作。")
	}
}
