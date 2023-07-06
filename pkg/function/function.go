// function 包用来练习函数
package function

import "fmt"

// main 是程序的入口点
func FuncMain() {
	fmt.Println("main()是入口函数")
	fmt.Println("Hello world!")
	fmt.Println("下面是带参数的函数↓↓↓")
	greet("黄鹏举")
	greet("黄姝禾")
	greetTwoParams1("黄鹏举", "黄姝禾")
	greetTwoParams2("huangpengju", "huangshuhe")

	fmt.Println("下面是带返回值的函数↓↓↓↓")
	fmt.Println(returnGreet("黄鹏举", "黄小米"))
	fmt.Println(fmt.Sprint(1, 2))
	fmt.Println(fmt.Sprint("1", "2"))
	fmt.Println(fmt.Sprint(1, "2"))
	fmt.Println(fmt.Sprint("1", 2))

	fmt.Println(returnGreets("黄鹏举", "黄姝禾"))

	// fmt.Println("可以返回变量类型的0值")

	fmt.Println(returnGreetss("黄鹏举", "黄小米"))

	// 求平均数 ，传入多个值
	n := average(1, 2, 3)
	fmt.Println(n)
	// 声明一个float64类型的切片
	data := []float64{1, 2, 3}
	n1 := average1(data...)
	fmt.Println(n1)

	//
	data1 := []float64{1, 2, 3}
	n2 := average2(data1)
	fmt.Println(n2)
	fmt.Println("___________")
	greeting()
	greeting1()
	greeting2()
	greeting3()
	greeting4()
}

// greet函数的声明包括一个string类型的参数
// 当调用greet时，传入一个参数
func greet(name string) {
	fmt.Println(name)
}

// 声明一个带有2个参数的函数
func greetTwoParams1(fname string, dname string) {
	fmt.Println(fname, dname)
}

// 声明一个带有2个参数的函数
// 同类型的参数，除最后一个外，其他参数的类型可以省略
func greetTwoParams2(fname, dname string) {
	fmt.Println(fname, dname)
}

// fmt.Sprint
// Sprint采用默认格式将其参数格式化，
// 串联所有输出生成并返回一个字符串。
// 如果两个相邻的参数都不是字符串，
// 会在它们的输出之间添加空格。
// 然后一个string类型的值
func returnGreet(fname, dname string) string {
	return fmt.Sprint(fname, dname)
}

// 函数可以返回一个指定类型的变量
func returnGreets(fname string, dname string) (s string) {
	s = fmt.Sprint(fname, dname)
	return

	/*
	   注意：应该
	   避免使用命名返回。
	   偶尔命名的返回是有用的。阅读本文了解更多信息：
	   https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html

	*/
}

// 返回两个string类型的值
func returnGreetss(fname string, dname string) (string, string) {
	return fmt.Sprint(fname, dname), fmt.Sprint(dname, fname)
}

// 求平均数
// 定义一个函数
// 这个函数的参数是可变的
// 返回一个float64类型的值
func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// 求平均数
// 这个函数的参数是可变的
func average1(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// 求平均数
// 这个函数的参数是切片
func average2(sf []float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// 函数表达式
// 定义一个普通函数
func greeting() {
	fmt.Println("Hello World!")
}

// 定义一个函数，在函数内部使用匿名函数
// 匿名函数作为表达式赋值给一个变量
func greeting1() {
	greeting := func() {
		fmt.Println("这里是匿名函数")
	}
	greeting()
}

// 查看匿名函数表达是什么类型
func greeting2() {
	greeting := func() {
		fmt.Println("这里是第二个匿名函数")
	}
	greeting()
	fmt.Printf("greeting的 Type=%T\n", greeting)
}

// makeGreeter的返回值是匿名函数作
// 匿名函数的返回值是string类型
func makeGreeter() func() string {
	return func() string {
		return "Hello 这是匿名函数的返回值"
	}
}

// greeting3调用另一个函数
func greeting3() {
	// 把一个函数作为表达式赋值给变量
	// 这个函数返回的是一个匿名函数
	greet := makeGreeter()
	fmt.Println(greet())
}

// greeting4调用另一个函数
func greeting4() {
	// 定义一个变量
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("greet的类型：%T\n", greet)
}
