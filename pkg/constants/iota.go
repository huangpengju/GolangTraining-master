package constants

import "fmt"

const aa = iota
const bb = iota

// 声明3个常量，使用iota
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const (
	a1 = iota // 0
	b1 = iota // 1
	c1 = iota // 2
)

const (
	a2 = iota // 0
	b2
	c2
)
const (
	_  = iota      //0
	b3 = iota * 10 // 1 * 10
	c3 = iota * 10 // 2 * 10
)
const (
	_ = iota // 0
	// 左移运算符（<<）将一个运算对象的各二进制位全部左移若干位（左边的二进制位丢弃，右边的补0）
	KB = 1 << (iota * 10) // 1 <<(1*10)
	// KB = 1<<10   // 1的二进制是 1 ，左边舍弃10位，右边补上10个0
	// 得到：100 0000 0000，将100 0000 0000 转换为十进制是 1 * 2^10 = 1024
	// KB= 1024
	MB = 1 << (iota * 10) // 1 <<(2*10)
	GB = 1 << (iota * 10) // 1 <<(3*10)
	TB = 1 << (iota * 10) // 1 <<(4*10)
)

// 使用iota定义常量
func IotaMain() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("_______")

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println("_______")
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)

	fmt.Println("_______")
	fmt.Println(b3)
	fmt.Println(c3)

	fmt.Println("_______")
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB) // 二进制
	fmt.Printf("%d\n", KB) // 十进制

	fmt.Printf("%b\t", MB) // 二进制
	fmt.Printf("%d\n", MB) // 十进制

	fmt.Printf("%b\t", GB) // 二进制
	fmt.Printf("%d\n", GB) // 十进制

	fmt.Printf("%b\t", TB) // 二进制
	fmt.Printf("%d\n", TB) // 十进制

	fmt.Printf("%T\n", TB)

	fmt.Println(aa)
	fmt.Println(bb)
}
