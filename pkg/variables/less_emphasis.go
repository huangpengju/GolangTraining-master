package variables

import "fmt"

// 声明变量，在Variables函数中使用
var a = "这存储在a中"
var b, c string = "存储在b中", "存储在c中"
var d string

// Exercise函数中使用
var oneSolution = "这是全局变量"

// 声明变量后和赋值
func DeclareVariable() {
	/*
		// 声明一个string类型的变量
		var message string
		// 给变量赋值
		message = "Hello World."
		//打印变量
		fmt.Println(message)
	*/
	// should merge variable declaration with assignment on next line (S1021)
	// 建议上面的声明和赋值合并到一行,如下：
	var message1 string = "Hello World!"

	fmt.Println(message1)

}

// 声明多个变量，给某个变量赋值
func DeclareManyAtOnce() {
	// 声明一个string类型的变量
	var message string
	// 声明多个int类型的变量
	var a, b, c int
	// 给变量赋值
	a = 1
	message = "Hello World!"
	fmt.Println(message, a, b, c)
}

// 声明变量，同时初始化
func InitManyAtOnce() {
	// 声明一个变量，go根据变量的值，确定变量的类型
	var message = "Hello World!"

	// 同声明3个int类型的变量，并初始化
	var a, b, c int = 11, 22, 33
	fmt.Println(message, a, b, c)
}

// 声明变量，同时初始化，变量的类型，用值来推断
func InferType() {
	// 声明一个变量，变量的类型由后面的值确定
	var message = "Hello World!!!"
	var a, b, c = 11, 12, 13
	fmt.Println(message, a, b, c)
}

// 声明变量，变量的类型由值确定，推断混合类型
func InferMixedUpTypes() {
	var message = "Hello World!"
	var a, b, c = 1, false, 3
	fmt.Printf("a=%v type:%T\nb=%v type:%T\nc=%v type:%T\n", a, a, b, b, c, c)
	fmt.Println(message, a, b, c)
}

// 短变量声明赋值的作用域
// 使用:= 声明并赋值变量时，只能在函数内部操作
func InitShortHand() {
	//只能在函数内部执行此操作
	message := "Hello World!"
	a, b, c := 1, false, 3
	d := 4
	e := true
	fmt.Println(message, a, b, c, d, e)
}

// 变量的应用
func Variables() {
	// 给全局变量d赋值
	d = "存储在d中"
	// 下面声明的变量具有作用域，作用域是此函数范围内
	var e = 42
	f := 43
	g := "存储在g中"
	h, i := "存储在h中", "存储在i中"
	j, k, l, m := 44.7, true, false, 'm' //单引号里面的单个字符，对应的是该字符的ASCII码
	n := "n"                             // 双引号表示字符串
	o1 := `o`
	o2 := `1`
	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o1 - ", o1)
	fmt.Println("o2 - ", o2)
	fmt.Printf("o1=%v type=%T\n", o1, o1)
	fmt.Printf("o2=%v type=%T\n", o2, o2)
}

// exercise 变量
func Exercise() {
	// 打印输出全局变量 oneSolution
	fmt.Println("Hello ~ oneSolution:→", oneSolution)

	// 声明局部变量 anotherSolution 并初始化
	// 打印输出
	var anotherSolution1 = "这是局部变量1号，使用双引号"
	fmt.Println("Hello ~ anotherSolution:→", anotherSolution1)

	// 短变量声明局部变量 anotherSolution2 并初始化
	// 打印输出
	anotherSolution2 := "这是局部变量2号，使用双引号"
	fmt.Println("Hello ~ anotherSolution:→", anotherSolution2)

	// 短变量声明局部变量 anotherSolution3 并初始化
	// 打印输出
	anotherSolution3 := `这是局部变量3号，使用反引号` // 反勾号像双引号一样工作
	fmt.Println("Hello ~ anotherSolution:→", anotherSolution3)

}
