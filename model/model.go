// model 用于程序入口直接调用
// model 操作其他包
package model

import (
	"GolangTraining-master/pkg/package/amain"
)

// "GolangTraining-master/pkg/package/amain"

// Model 函数用于程序入口的调用
func Model() {
	// 1.输出练习
	// start1()

	// 2.包到调用练习
	amain1()
}

// start 函数用于练习 fmt.Println 和 fmt.Printf
func start1() {
	// 输出hello world!
	// start.HelloWorld()

	// 输出十进制的数
	// start.Decimal()

	// 输出二进制（同时与十进制进行对比）
	// start.Binary()

	// 输出十六进制（同时与十进制、二进制进行对比；并且加 #显示 x0 和 X0）
	// start.Hexadecimal()

	// 输出多个 十进制 以及 对应的二进制和十六进制
	// start.Loop()

	// 输出多个 十进制 以及 对应的二进制和十六进制 以及输出该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	// start.Utf_8()
}

func amain1() {
	amain.Amain()
}
