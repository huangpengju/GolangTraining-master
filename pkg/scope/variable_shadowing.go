package scope

import "fmt"

func max(x int) int {
	return 42 + x
}

// VarShadowingMain 变量与函数名一致
func VarShadowingMain() {
	max := max(7)
	fmt.Println(max) // Max现在是结果，而不是函数

	// 调用same.go
	fmt.Println("same.go : sameX=", sameX)
}

// 不要这样做;
// 隐藏变量的不良编码实践
