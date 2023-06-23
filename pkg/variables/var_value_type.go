// variables 在此处用于短变量声明及初始化
package variables

import "fmt"

// VarValue短变量声明及初始化
func ShortHandVal() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	// 单引号在go语言中表示golang中的rune（int32）类型
	// 单引号里面是单个字符，对应值为该字符的ASCII值
	g := 'M'
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
}

// VarValue短变量声明及初始化,查看变量的类型
func ShortHandType() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	// 单引号在go语言中表示golang中的rune（int32）类型
	// 单引号里面是单个字符，对应值为该字符的ASCII值
	g := 'M'
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
}
