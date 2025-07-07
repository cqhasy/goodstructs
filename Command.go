package main

import "fmt"

//很有用啊

type TV struct{}

func (tv *TV) On() {
	fmt.Println("电视已打开")
}

func (tv *TV) Off() {
	fmt.Println("电视已关闭")
}

type Command interface {
	Execute()
	Undo()
}

type OnCommand struct {
	tv *TV
}

func (c *OnCommand) Execute() {
	c.tv.On()
}

func (c *OnCommand) Undo() {
	c.tv.Off()
}

type OffCommand struct {
	tv *TV
}

func (c *OffCommand) Execute() {
	c.tv.Off()
}

func (c *OffCommand) Undo() {
	c.tv.On()
}

type RemoteControl struct {
	history []Command
}

func (r *RemoteControl) Press(cmd Command) {
	cmd.Execute()
	r.history = append(r.history, cmd)
}

func (r *RemoteControl) PressUndo() {
	if len(r.history) == 0 {
		fmt.Println("没有命令可撤销")
		return
	}
	last := r.history[len(r.history)-1]
	last.Undo()
	r.history = r.history[:len(r.history)-1]
}
func main() {
	tv := &TV{}
	onCmd := &OnCommand{tv}
	offCmd := &OffCommand{tv}

	remote := &RemoteControl{}
	remote.Press(onCmd)  // 输出：电视已打开
	remote.Press(offCmd) // 输出：电视已关闭

	remote.PressUndo() // 撤销：电视已打开
	remote.PressUndo() // 撤销：电视已关闭
}
