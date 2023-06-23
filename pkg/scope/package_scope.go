// scope 练习包的作用范围
package scope

import (
	vis "GolangTraining-master/pkg/scope/visibility"
	"fmt"
)

// 定义个变量 全局变量可以在同一个包内 多个函数中使用，首字母无大写要求
var x = 42

func ScopeMain() {
	fmt.Println(x)
	foo()
	// 使用vis包中的函数
	vis.PrintVar()
	// 使用vis包中的变量
	fmt.Println(vis.MyName)
}

// foo 首字母小写，可以在同一个包中使用
func foo() {
	fmt.Println(x)
}
