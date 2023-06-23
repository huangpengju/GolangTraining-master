package vis

import "fmt"

// PrintVar 首字母大写，可以被导出
func PrintVar() {
	// 变量MyName 和 yourName 在同一个包中均可使用
	fmt.Println(MyName)
	fmt.Println(yourName)
}
