package main

import (
	"fmt"
	"sync"
)

type Image interface {
	Display(user string)
}

type RealImage struct {
	filename string
}

func (r *RealImage) Display(user string) {
	fmt.Println("Displaying image:", r.filename)
}
func NewRealImage(filename string) Image {
	fmt.Println("Loading image from disk:", filename)
	return &RealImage{filename: filename}
}

type ImageFactory func(filename string) Image

type ProxyImage struct {
	filename    string
	realImage   Image
	displayed   bool
	permissions map[string]bool // 用户权限
	logs        []string
	mu          sync.Mutex
	factory     ImageFactory
}

func NewProxyImage(filename string, allowedUsers []string, factory ImageFactory) *ProxyImage {
	perms := make(map[string]bool)
	for _, u := range allowedUsers {
		perms[u] = true
	}
	return &ProxyImage{
		filename:    filename,
		permissions: perms,
		logs:        []string{},
		factory:     factory,
	}
}

func (p *ProxyImage) Display(user string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 权限控制
	if !p.permissions[user] {
		p.logs = append(p.logs, fmt.Sprintf("Unauthorized access by user: %s", user))
		fmt.Println("Access denied for user:", user)
		return
	}

	// 缓存控制
	if p.realImage == nil {
		p.realImage = p.factory(p.filename)
	}

	// 日志记录
	p.logs = append(p.logs, fmt.Sprintf("User %s displayed image: %s", user, p.filename))

	// 实际显示
	p.realImage.Display(user)
}

func (p *ProxyImage) ShowLogs() {
	fmt.Println("---- Logs ----")
	for _, log := range p.logs {
		fmt.Println(log)
	}
}
func main() {
	proxy := NewProxyImage("secure.jpg", []string{"admin", "user1"}, NewRealImage)

	proxy.Display("guest") // 无权限
	proxy.Display("user1") // 成功
	proxy.Display("admin") // 成功
	proxy.Display("user1") // 缓存生效，无需重新加载

	proxy.ShowLogs()
}
