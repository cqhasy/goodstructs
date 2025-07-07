package main

import (
	"goodstruct/State"
)

func main() {
	e := State.NewElevator()

	e.Handle(State.EventOpenDoors)  // 开门
	e.Handle(State.EventMove)       // 移动失败
	e.Handle(State.EventCloseDoors) // 关门
	e.Handle(State.EventMove)       // 移动
	e.Handle(State.EventStop)       // 停止
	e.Handle(State.EventOpenDoors)  // 开门
}
