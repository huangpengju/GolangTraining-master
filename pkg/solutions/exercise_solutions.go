// solutions 用来综合练习
package solutions

import (
	"fmt"
)

// HelloWorld 输出Hello World
func HelloWorld() {
	fmt.Println("Hello World!")
	fmt.Println("_____________1")
}

// HelloName 输出字符串和变量的值
func HelloName() {
	name := "HuangPengju"
	fmt.Println("Hello", name)
	fmt.Println("_____________2")

}

// 接收用户输入后， 把变量内容输出
func HelloUserInput() {
	var name string
	fmt.Print("请输入名字后回车：")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
	fmt.Println("_____________3")

}

// 接收用户输入的两个数（int类型）,并计算取余%
func UserEntersNumbers() {
	var numOne int
	var numTwo int
	fmt.Print("请输入第一个数：")
	fmt.Scan(&numOne)
	fmt.Print("请输入第二个数：")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "%", numTwo, " = ", numOne%numTwo)
	fmt.Println("_____________4")

}

// 找出100以内的偶数
func EvenNumbers() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("_____________5")

}

// 找出100以内，15的倍数，3的倍数，5的倍数
func FizzBuzz() {
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, " -- FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, " -- FIZZ")
		} else if i%5 == 0 {
			fmt.Println(i, " -- BUZZ")
		} else {
			fmt.Println(i, "不是3,5,15的倍数")
		}
	}
	fmt.Println("_____________6")

}

// 找出1000以内，3和5的倍数，并求和
func ThreeFive() {
	// 计数器
	counter := 0
	for i := 0; i <= 1000; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println("counter = ", counter)
	fmt.Println("_____________7")

}

// Utf 练习fmt.Println() 返回2个参数
func Utf() {
	a, e := fmt.Println("Hello"[1]) // 输出：101 , "Hello"[1] 表示Hello的第2个元素e对应的ASCII码
	fmt.Println(a)                  // 4 表示写入到标准输出的字节数
	fmt.Println(e)                  // <nil>  表示写入到标准输出的错误

}
