// amain 用于练习包的运用
// 比如：a包可以使用b包和c包下的函数或变量（大写字母开头）
// b包可以使用b包下的函数（小写字母开头）
package amain

import (
	"GolangTraining-master/pkg/package/bstringutil"
	"GolangTraining-master/pkg/package/cvar"
	"fmt"
)

// AMain 用于让Model 调用
func Amain() {
	// 调用bstringutil包中的变量
	fmt.Println(bstringutil.Name)
	// 调用bstringutil包中的函数
	fmt.Println(bstringutil.Reverse("!oG,olleH"))

	// 调用cvar包中的变量
	fmt.Println(cvar.Name)
}
