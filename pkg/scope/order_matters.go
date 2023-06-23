package scope

import "fmt"

// OrderMattersMain 函数内部变量的声明和使用顺序很重要
func OrderMattersMain() {
	// fmt.Println(orderX)
	fmt.Println(orderY)
	// orderX := 42
}

var orderY = 42
