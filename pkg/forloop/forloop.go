// forloop 包用来练习for循环
package forloop

import (
	"fmt"
)

// InitCondition 初始化for循环条件,输出循环条件i
func InitCondition() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

// Nested 练习for的嵌套循环
func Nested() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("i=", i, " | j=", j)
		}
	}
}

// ForConditionWhile 练习for带条件的的while用法
func ForConditionWhile() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

// forNoConditionWhile 练习for不带条件的的while用法
func ForNoConditionWhile() {
	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }
}

// ForBreak 练习使用break跳出死循环
func ForBreak() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 11 {
			// break 结束循环
			break
		}
		i++
	}

}

// ForContinue 练习使用break跳出死循环，同时使用continue跳出某一轮循环，开始下一轮循环
func ForContinue() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue //跳出本轮循环,开始下轮循环
		}
		fmt.Println(i)
		if i >= 50 {
			break // 结束循环
		}
	}
}

// 循环输出字符
func RuneLoopUtf81() {
	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))

		foo := "a"
		fmt.Println(foo)
		fmt.Printf("%T \n", foo)
	}
	// a := 87
	// fmt.Println(a)
	// fmt.Printf("%T \n", string(a))
	// fmt.Println(string(a))
	// fmt.Printf("%T \n", []byte(string(a)))
	// fmt.Printf("%v \n", []byte(string(a)))

	/*
		注意:
		某些操作系统(Windows)可能不会打印i < 256的字符

		如果你有这个问题，你可以使用下面的代码:

		fmt。Println(“-”,字符串(我 ), " - ", [] int32(字符串(i)))

		UTF-8是Go使用的文本编码方案。

		UTF-8使用1 - 4字节。

		一个字节是8位。

		[]byte处理字节，即一次只处理1个字节(8位)。

		[]int32允许存储4字节的值，即4字节*每字节8位= 32位。
	*/
}

// 循环输出字符
func RuneLoopUtf82() {
	for i := 49; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}

	fmt.Println("_+++++")
	a := 84
	fmt.Println(a)
	fmt.Printf("a 的 type = %T \n", a)

	fmt.Printf("string(a) 的type = %T \n", string(a))
	fmt.Println(string(a))

	fmt.Printf("[]byte(tring(a)) 的type = %T \n", []byte(string(a)))
	fmt.Println([]byte(string(a)))
	fmt.Println("____")
	fmt.Println([]byte("T"))
}
