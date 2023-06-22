// start 包用来练习 fmt.Println 和 fmt.Printf
package start

import "fmt"

// Decimal 输出十进制数
func Decimal() {
	fmt.Println(42)
}

// 输出十进制数和二进制数，对比
func Binary() {
	fmt.Printf("%d-%b\n", 42, 42)
}

// 输出十六进制数
func Hexadecimal() {
	// 输出十进制 二进制 十六进制
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	// 输出十进制 二进制 十六进制(小写的0x前缀)
	fmt.Printf("%d - %b - %#x\n", 42, 42, 42)
	// 输出十进制 二进制 十六进制(大写的0x前缀)
	fmt.Printf("%d - %b - %#X\n", 42, 42, 42)
	// 输出十进制 二进制 十六进制(小写的0x前缀) \t 一个tab 8
	fmt.Printf("%d \t %b \t %#X\n", 42, 42, 42)

}

// 循环输出 十进制 以及对应的 二进制 和 十六进制
func Loop() {
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %x\n", i, i, i)
	}
}
