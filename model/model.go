// model 用于程序入口直接调用
// model 操作其他包
package model

// "GolangTraining-master/pkg/package/amain"

// Model 函数用于程序入口的调用
func Model() {
	// 1.输出练习
	// start1()

	// 2.包到调用练习
	// amain1()

	// 3.变量的练习
	// variables1()

	// 4.包和变量的作用范围
	// scope1()

	// 5. 空白标识符的使用
	// blank1()

	// 6. 常量是一个简单不变的值
	// Constants1()

	// 7. 内存地址
	// address1()

	// 8. 指针
	// pointers1()

	// 9.求余数 %
	// remainder1()

	// 10.练习for 循环
	// initCondition1()

	// 11.练习switch
	// switch1()

	// 12.练习if else
	// ifelse1()

	// 13.综合练习
	// exercisesolutions1()
}

/*
// start 函数用于练习 fmt.Println 和 fmt.Printf
func start1() {
	// 输出hello world!
	start.HelloWorld()

	// 输出十进制的数
	start.Decimal()

	// 输出二进制（同时与十进制进行对比）
	start.Binary()

	// 输出十六进制（同时与十进制、二进制进行对比；并且加 #显示 x0 和 X0）
	start.Hexadecimal()

	// 输出多个 十进制 以及 对应的二进制和十六进制
	start.Loop()

	// 输出多个 十进制 以及 对应的二进制和十六进制 以及输出该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	start.Utf_8()
}
*/
/*
// 包的练习
func amain1() {
	amain.Amain()
}
*/
/*
// 变量的练习
func variables1() {
	// 输出短变量声明及初始化的内容
	variables.ShortHandVal()

	fmt.Println("_____")
	// 输出短变量声明及初始化的内容的类型
	variables.ShortHandType()

	fmt.Println("_____")
	// 声明变量未初始化，变量使用该类型的默认零值
	variables.VarZeroValue()

	fmt.Println("_____")
	// 声明变量后赋值
	variables.DeclareVariable()

	fmt.Println("_____")
	// 同时声明多个变量
	variables.DeclareManyAtOnce()

	fmt.Println("_____")
	// 声明变量，同时初始化（go可以根据值的类型，确定变量的类型）
	variables.InitManyAtOnce()

	fmt.Println("_____")
	// 声明变量，变量的类型由值确定(推断)
	variables.InferType()

	fmt.Println("_____")
	// 推断混合类型
	variables.InferMixedUpTypes()

	fmt.Println("_____")
	// 使用:= 声明并赋值变量时，只能在函数内部操作
	variables.InitShortHand()

	fmt.Println("_____")
	// 变量的应用
	// 单引号里面的单个字符，对应的是该字符的ASCII码
	variables.Variables()

	fmt.Println("_____")
	// 变量的练习
	// 反引号像双引号一样工作
	variables.Exercise()
}
*/
/*
// scop1 练习包的作用范围
func scope1() {
	// 变量和函数 在同一包和不同包中的使用
	scope.ScopeMain()

	fmt.Println("______")
	// 函数内部的作用域
	// 函数内部变量的声明，作用域只在该函数内部
	scope.BlockScopMain()

	fmt.Println("______")
	// 函数中变量的作用域和{}代码块中的作用域
	scope.Closure1()

	fmt.Println("______")
	// 包作用域下，同一变量（全局）可以被多个函数使用
	scope.Closure2()

	fmt.Println("______")

	// 闭包帮助我们限制多个函数使用的变量的作用域，
	// 如果没有闭包，两个或多个函数要访问同一个变量，
	// 该变量需要是包作用域
	// 匿名函数 没有名称的函数
	// 函数表达式 将函数赋值给变量
	scope.Closure3()

	fmt.Println("______")
	// 匿名函数和作用域内的变量
	scope.Closure4()

	fmt.Println("______")
	// 顺序很重要（局部变量）
	scope.OrderMattersMain()

	fmt.Println("______")
	// 变量跟踪
	// VarShadowingMain 变量与函数名一致
	scope.VarShadowingMain()
}
*/
/*
// blank1 包用来测试 _ 占位符(空白标识符)
func blank1() {
	// 在函数中如果声明的变量没有使用会报错
	blank.BlankMain()
	// 对于不使用的变量可以用_标识
	// 例子
	blank.Checking()
	// 对于例子程序中不需要使用的变量可以用_标识
	blank.NotChecking()
}
*/
/*
// 常量是一个简单不变的值
// iota 是常量计数器
// 左移运算符（<<）将一个运算对象的各二进制位全部左移若干位（左边的二进制位丢弃，右边的补0）
func Constants1() {
	// 声明一个常量，并使用
	constants.Constant()

	fmt.Println("_______")
	// 声明一个常量，iota 是常量计数器
	// iota是go语言的常量计数器,只能在常量表达式中使用 iota在const关键字出现时将被重置为0,const中每新增一行常量声明将使iota计数一次 可理解为const语句块中的行索引。
	constants.IotaMain()

}
*/
/*
func address1() {
	// showingMain 用来学习获取变量的内存地址，获得的值是十六进制数
	address.ShowingMain()

	fmt.Println("_______")
	// 使用变量的内存地址
	address.UsingMain()
}
*/

/*
// pointer1 练习指针
func pointers1() {
	// 获取变量地址，引用
	pointers.Referencing()

	fmt.Println("_______")
	// 解引用
	pointers.Dereferencing()

	fmt.Println("______")
	// 使用指针
	pointers.UsingPointer1()

	fmt.Println("______")
	// 使用指针
	pointers.UsingPointer2()
}
*/
/*
// remainder 用练习%求余数
func remainder1() {
	// 求余数
	remainder.Remainder()
}
*/
/*
// initCondition1用来练习for
func initCondition1() {
	// 练习for
	forloop.InitCondition()
	// 练习for for嵌套循环
	forloop.Nested()
	// 练习for带条件的的while用法
	forloop.ForConditionWhile()
	// 练习for不带条件的的while用法（死循环）
	forloop.ForNoConditionWhile()
	// 练习for不带条件的的while用法（break跳出死循环）
	forloop.ForBreak()

	fmt.Println("___________")
	// 练习使用break跳出死循环，同时使用continue跳出某一轮循环，开始下一轮循环
	forloop.ForContinue()

	fmt.Println("___________")
	// 输出字符 UTF-8是Go使用的文本编码方案。
	forloop.RuneLoopUtf81()
	fmt.Println("___________")
	// 输出字符 UTF-8是Go使用的文本编码方案。
	forloop.RuneLoopUtf82()
}
*/

/*
// switch 判断
func switch1() {
	// switch的使用
	switchstatements.SwitchMain()
	fmt.Println("_______")
	// fallthrough 后面的case无条件执行
	switchstatements.FallthroughTest()

	fmt.Println("________")
	// 同时练习多个判断条件，满足一个即可
	switchstatements.MultipleEvals()
	fmt.Println("______")
	// 练习switch没有表达式的情况
	switchstatements.NoExpression()
	fmt.Println("________")
	// 用swithc判断变量的类型
	switchstatements.OnType()
}
*/
/*
// 练习 if else语句
func ifelse1() {
	// true 和 false
	ifelse.Ifelse()
	fmt.Println("_________")

	// NotExclamation 函数用来练习 !true 和 !false
	ifelse.NotExclamation()

	fmt.Println("_____")
	// 初始化语句
	ifelse.InitStatement()
	// 初始化语句，变量只能在代码块中使用
	ifelse.InitStatementErrInvalidCode()
	fmt.Println("_______")

	// 练习if else 分支语句
	ifelse.IfElse()
	fmt.Println("____+++++___")
	// 练习if else if else
	ifelse.IfElseIfElse()
	fmt.Println("_______")
	// 练习if else if else if else
	ifelse.IfElseIfElseIfElse()

	fmt.Println("_______")
	// 能被3整除的数
	ifelse.DivisibleByThree()
}
*/
/*
// solutions 用来综合练习
// 练习了fmt.Println 换行输出
// 练习了fmt.Print 输出
// 练习了fmt.Scan 输入
// 练习了 取余数 %
// 练习了 if else
// 练习了 for
// 练习了 += 累加运算
func exercisesolutions1() {
	// 简单的输出
	// solutions.HelloWorld()
	// 输出变量的内容
	// solutions.HelloName()
	// 先输入，然后输出变量的内容
	// solutions.HelloUserInput()
	// 先输入，然后输出变量的内容，并取余数
	// solutions.UserEntersNumbers()
	// 求100以内的偶数
	// solutions.EvenNumbers()
	// 求1000以内3 5 15的倍数
	// solutions.FizzBuzz()

	// 求1000以内，3 和 5 的倍数之和
	// solutions.ThreeFive()

	//
	solutions.Utf()
}
*/
