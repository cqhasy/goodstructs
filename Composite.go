package main

import (
	"fmt"
)

// FileSystemComponent 是组件接口，表示文件系统的通用行为
type FileSystemComponent interface {
	DisplayInfo(indent string)
}

// File 是叶子节点，表示文件
type File struct {
	Name string
}

func (f *File) DisplayInfo(indent string) {
	fmt.Printf("%sFile: %s\n", indent, f.Name)
}

// Directory 是复合节点，表示目录，可包含文件和子目录
type Directory struct {
	Name       string
	Components []FileSystemComponent
}

func NewDirectory(name string) *Directory {
	return &Directory{
		Name:       name,
		Components: make([]FileSystemComponent, 0),
	}
}

func (d *Directory) Add(component FileSystemComponent) {
	d.Components = append(d.Components, component)
}

func (d *Directory) DisplayInfo(indent string) {
	fmt.Printf("%sDirectory: %s\n", indent, d.Name)
	for _, component := range d.Components {
		component.DisplayInfo(indent + "  ")
	}
}

func main() {
	file1 := &File{Name: "file1.txt"}
	file2 := &File{Name: "file2.txt"}

	subDir := NewDirectory("Subdirectory")
	subDir.Add(file1)
	subDir.Add(file2)

	root := NewDirectory("Root")
	root.Add(subDir)
	root.Add(&File{Name: "readme.md"})

	// 展示文件系统结构
	root.DisplayInfo("")
}
