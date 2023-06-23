// model 用于程序入口直接调用
// model 操作其他包
package model

import (
	"GolangTraining-master/pkg/variables"
	"fmt"
)

// "GolangTraining-master/pkg/package/amain"

// Model 函数用于程序入口的调用
func Model() {
	// 1.输出练习
	// start1()

	// 2.包到调用练习
	// amain1()

	// 3.变量的练习
	variables1()
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
