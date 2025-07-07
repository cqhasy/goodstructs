package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type FColor string

const (
	Red    FColor = "Red"
	Green  FColor = "Green"
	Blue   FColor = "Blue"
	Yellow FColor = "Yellow"
)

// Shape接口
type FShape interface {
	Draw(x, y int)
	UpdateAccessTime()
	LastAccessTime() time.Time
}

// Circle具体享元
type FCircle struct {
	color       FColor
	lastAccess  time.Time
	accessMutex sync.Mutex
}

func (c *FCircle) Draw(x, y int) {
	c.UpdateAccessTime()
	fmt.Printf("Drawing a %s circle at (%d,%d)\n", c.color, x, y)
}

func (c *FCircle) UpdateAccessTime() {
	c.accessMutex.Lock()
	defer c.accessMutex.Unlock()
	c.lastAccess = time.Now()
}

func (c *FCircle) LastAccessTime() time.Time {
	c.accessMutex.Lock()
	defer c.accessMutex.Unlock()
	return c.lastAccess
}

type FShapeFactory struct {
	circleMap      map[FColor]*FCircle
	mu             sync.RWMutex
	expiryDuration time.Duration
	stopCleaner    chan struct{}
}

var factory *FShapeFactory
var Fonce sync.Once

func GetShapeFactory() *FShapeFactory {
	Fonce.Do(func() {
		factory = &FShapeFactory{
			circleMap:      make(map[FColor]*FCircle),
			expiryDuration: 5 * time.Second, // 比如5秒没被访问就清理掉
			stopCleaner:    make(chan struct{}),
		}
		// 启动后台清理协程
		go factory.cleanupLoop()
	})
	return factory
}

func (f *FShapeFactory) GetCircle(color FColor) FShape {
	f.mu.RLock()
	circle, exists := f.circleMap[color]
	f.mu.RUnlock()

	if exists {
		circle.UpdateAccessTime()
		return circle
	}

	f.mu.Lock()
	// double-check，防止写锁切换期间被创建了
	circle, exists = f.circleMap[color]
	if !exists {
		circle = &FCircle{color: color}
		circle.UpdateAccessTime()
		f.circleMap[color] = circle
	}
	f.mu.Unlock()

	return circle
}

// 后台定时清理未访问享元对象
func (f *FShapeFactory) cleanupLoop() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			f.cleanExpired()
		case <-f.stopCleaner:
			return
		}
	}
}

func (f *FShapeFactory) cleanExpired() {
	f.mu.Lock()
	defer f.mu.Unlock()

	now := time.Now()
	for color, circle := range f.circleMap {
		if now.Sub(circle.LastAccessTime()) > f.expiryDuration {
			fmt.Printf("Cleaning up %s circle due to inactivity\n", color)
			delete(f.circleMap, color)
		}
	}
}

// 关闭工厂清理协程
func (f *FShapeFactory) Stop() {
	close(f.stopCleaner)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	colors := []FColor{Red, Green, Blue, Yellow}

	factory = GetShapeFactory()
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 20; i++ {
				randomColor := colors[rand.Intn(len(colors))]
				circle := factory.GetCircle(randomColor)
				circle.Draw(rand.Intn(100), rand.Intn(100))
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			}
		}()
	}

	// 运行一段时间后停止后台清理
	time.Sleep(10 * time.Second)
	factory.Stop()
}
