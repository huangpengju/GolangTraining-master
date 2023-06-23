// variables 在此处用于var 声明变量
package variables

import "fmt"

// VarZeroValue 是声明变量未初始化，变量使用该类型的默认零值
func VarZeroValue() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
