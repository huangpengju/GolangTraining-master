// constants 包用来学习常量以及常量的用法
package constants

import "fmt"

// 定义一个全局常量，使用关键字const
const p = "这是一个常量P"

// 声明多个常量，并初始化
const (
	pi       = 3.13
	language = "Go"
)

// 常量是一个简单不变的值
func Constant() {
	// 在函数内部定义一个常量
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println("使用多个常量：")
	fmt.Println(pi)
	fmt.Println(language)
}
