package main

//当你 赋值一个切片变量给另一个切片变量 时，实际上复制的是切片结构体里的三个字段：
//
//底层数组的 指针（指向实际元素的内存地址）
//
//切片的长度
//
//切片的容量
//
//这里拷贝的是指针的副本，所以两个切片变量都指向同一个底层数组。

type Person struct {
	Name   string
	Scores []int
}

func (p *Person) Clone() *Person {
	// 深拷贝Scores
	scoresCopy := make([]int, len(p.Scores))
	copy(scoresCopy, p.Scores)
	return &Person{
		Name:   p.Name,
		Scores: scoresCopy,
	}
}
