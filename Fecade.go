package main

//暴露一个高层接口进行内部的复杂调用，简化调用难度，有点像mvc的controller对service,但本质还是不同

import "fmt"

// StereoSystem 子系统：音响
type StereoSystem struct{}

func (s *StereoSystem) TurnOn()  { fmt.Println("Stereo System is turned on") }
func (s *StereoSystem) TurnOff() { fmt.Println("Stereo System is turned off") }

// Projector 子系统：投影仪
type Projector struct{}

func (p *Projector) TurnOn()  { fmt.Println("Projector is turned on") }
func (p *Projector) TurnOff() { fmt.Println("Projector is turned off") }

// LightsControl 子系统：灯光控制
type LightsControl struct{}

func (l *LightsControl) TurnOn()  { fmt.Println("Lights are turned on") }
func (l *LightsControl) TurnOff() { fmt.Println("Lights are turned off") }

// HomeTheaterFacade 外观类：家庭影院外观
type HomeTheaterFacade struct {
	stereo    *StereoSystem
	projector *Projector
	lights    *LightsControl
}

// NewHomeTheaterFacade 构造函数
func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		stereo:    &StereoSystem{},
		projector: &Projector{},
		lights:    &LightsControl{},
	}
}

// WatchMovie 外观方法：准备观影
func (ht *HomeTheaterFacade) WatchMovie() {
	fmt.Println("Getting ready to watch a movie...")
	ht.lights.TurnOff()
	ht.projector.TurnOn()
	ht.stereo.TurnOn()
}

// EndMovie 外观方法：结束观影
func (ht *HomeTheaterFacade) EndMovie() {
	fmt.Println("Ending the movie...")
	ht.stereo.TurnOff()
	ht.projector.TurnOff()
	ht.lights.TurnOn()
}

func main() {
	homeTheater := NewHomeTheaterFacade()

	// 准备观影
	homeTheater.WatchMovie()

	// 结束观影
	homeTheater.EndMovie()
}
