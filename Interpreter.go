package main

import "fmt"

//感觉和组合模式有点像，构造一棵树，针对一个接口进行多种方法的实现
//表达了一个语言的“语法规则”
//
//构建了“抽象语法树”
//
//提供了解析执行的“解释逻辑”

type intInterpreter interface {
	interpret(...int) int
}
type num int
type AddInterpreter struct{}

func (addInterpreter *AddInterpreter) interpret(args ...int) int {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

type SubInterpreter struct {
	target int
}

func (subInterpreter *SubInterpreter) interpret(args ...int) int {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return subInterpreter.target - sum
}
func (a *num) interpret(args ...int) int {
	return (int)(*a)
}
func main() {
	// 构建终结符表达式：数字
	n1 := num(1)
	n2 := num(2)
	n3 := num(3)
	n10 := num(10)

	// 构建加法表达式：1 + 2 + 3
	add := &AddInterpreter{}
	sum := add.interpret(n1.interpret(), n2.interpret(), n3.interpret(), (&SubInterpreter{target: n10.interpret()}).interpret(n2.interpret()))

	// 构建减法表达式：10 - (1 + 2 + 3)
	sub := &SubInterpreter{target: n10.interpret()}
	result := sub.interpret(sum)

	fmt.Println("Result:", result) // Output: Result: 4
}
