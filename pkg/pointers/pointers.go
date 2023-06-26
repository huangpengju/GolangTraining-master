// pointers包用来练习指针
package pointers

import "fmt"

// Referencing() 取地址 引用
func Referencing() {
	a := 43
	fmt.Println("a=", a)
	fmt.Println("a的地址：", &a)

	var b = &a
	fmt.Println("变量b=", b)
	fmt.Printf("b的类型%T\n", b)
	//上面的代码使b成为存储int的内存地址的指针

	//b的类型为“int pointer”

	//*int--*是类型的一部分--b是*int类型
}

// 解引用
func Dereferencing() {
	a := 43
	fmt.Println("a=", a)
	fmt.Println("a的地址是：", &a)

	// 声明一个变量，存放a的地址
	var b = &a

	fmt.Println("b=", b)
	fmt.Println("解析b=", *b)

	// fmt.Println("解析a=", *&a)
	//b是一个int指针；

	//b指向存储int的内存地址

	//要查看该内存地址中的值，请在b前面添加*

	//这就是所谓的取消引用

	//在这种情况下，*是运算符（）
}

// UsingPointer 使用指针
func UsingPointer1() {
	a := 43
	fmt.Println("a=", 43)
	fmt.Println("a取地址=", &a)

	var b = &a
	fmt.Println("b=", b)
	fmt.Println("b解引用=", *b)

	// 给b中的这个地址的值改为42
	*b = 42
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("b解引用=", b)

	//这很有用

	//我们可以传递一个内存地址，而不是一堆值（我们可以传递引用）

	//然后我们仍然可以更改存储在该内存地址的任何值

	//这使我们的程序更具性能

	//我们不必传递大量数据

	//我们只需要传递地址

	//顺便说一句go中的一切都是通过价值

	//当我们传递内存地址时，我们传递的是一个值
}

// 没有使用指针
func UsingPointer2() {
	x := 5
	// 没有使用指针
	// zero1(x)
	fmt.Println("x的地址是：", &x)
	// fmt.Printf("x的地址：%p\n", &x) //x的地址是main函数中x的地址
	// zero2(x)
	// fmt.Println("x=", 5)

	// 传递x的地址
	// zero3(&x)
	// fmt.Println("x=", x)

	// 传递x的地址
	zero4(&x)

	fmt.Println("x=", x)
}

// 不会覆盖x
// func zero1(z int) {
// 	z = 0
// }

// zero中参数z 新的变量
// func zero2(z int) {
// 	fmt.Printf("zero=,%p\n", &z) //z的地址是函数zero中参数的地址
// 	fmt.Println("zero:=", &z)
// 	z = 0
// }

// zero3 的参数z 是一个int类型的指针
// func zero3(z *int) {
// 	*z = 0
// }

func zero4(z *int) {
	fmt.Println("z=", z)
	*z = 0
}
