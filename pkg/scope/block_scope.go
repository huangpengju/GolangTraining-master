package scope

import "fmt"

var cc = 0

// 函数内部变量的声明，作用域只在该函数内部
func BlockScopMain() {
	a := 42
	fmt.Println(a)
	blockFoo()
}

// 函数内部变量的声明，作用域只在该函数内部
func blockFoo() {
	// 无法访问a
	// 这不会编译
	// 因为a没有声明
	// fmt.Println(a)
}

// 函数中变量的作用域和{}代码块中的作用域
func Closure1() {
	x1 := 42
	fmt.Println("x1在{}外", x1)
	{
		fmt.Println("x1在{}内", x1)
		y := "这是y,在{}中声明"
		fmt.Println(y)
	}
	// fmt.Println(y) // y超出了{}范围外

}

func Closure2() {
	fmt.Println(increment())
	fmt.Println(increment())
}

// 对变量c进行自增
func increment() int {
	cc++
	return cc
}

/*
闭包帮助我们限制
没有闭包的多个函数使用的变量的作用域，
对于两个或多个函数访问同一个变量，
该变量需要是包作用域
*/

// 匿名函数的使用
func Closure3() {
	cc := 0
	increment := func() int {
		cc++
		return cc
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(cc)
}

/*
闭包帮助我们限制多个函数使用的变量的作用域
如果没有闭包，两个或多个函数要访问同一个变量，
该变量需要是包作用域

匿名函数
没有名称的函数

函数表达式
将函数赋值给变量
*/

// 匿名函数和作用域内的变量
func Closure4() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

// wrapper() 返回一个匿名函数
func wrapper() func() int {
	x := 0
	fmt.Println("x1的地址：", &x)
	return func() int {
		x++
		fmt.Println("x2的地址：", &x)
		return x
	}

}
