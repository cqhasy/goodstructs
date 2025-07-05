package main

import (
	"fmt"
	"sync"
)

// 私有结构体，包外无法直接创建
type singleton struct{}

// 私有实例，不能被包外访问
var instance *singleton
var once sync.Once

// Singleton 公共接口，只暴露需要的方法
type Singleton interface {
	ShowMessage()
}

// GetInstance 获取单例实例，返回的是接口类型，限制使用范围
func GetInstance() Singleton {
	if instance == nil {
		once.Do(func() {
			instance = &singleton{}
		})
	}
	return instance
}

// ShowMessage 接口实现方法
func (s *singleton) ShowMessage() {
	fmt.Println("Hello, I am a fully encapsulated Singleton!")
}
