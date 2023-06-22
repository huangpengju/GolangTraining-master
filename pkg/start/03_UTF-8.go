// start 包用来练习 fmt.Println 和 fmt.Printf
package start

import "fmt"

// Utf_8 输出多个 十进制 和之 对应的二进制和十六进制 以及输出该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
func Utf_8() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
