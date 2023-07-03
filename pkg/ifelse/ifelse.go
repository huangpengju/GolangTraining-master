// ifelse 包用来学习if lese 分支语句
package ifelse

import "fmt"

// Ifelse函数用来学习if else
func Ifelse() {
	// true 和 false
	// if 判断是不是 真值(true)
	if true {
		fmt.Println("这是跑")
	}
	fmt.Println("111")
	// 判断条件为 false  时，不通过
	if false {
		fmt.Println("这是不跑")
	}
}

// NotExclamation 函数用来练习 !true 和 !false
func NotExclamation() {
	// true 和 false
	// !true == false
	// !flase == true
	if !true {
		fmt.Println("非真就是假")
	}

	if !false {
		fmt.Println("这里是!false,非假就是真")
	}
}

// init statement 初始化语句
func InitStatement() {
	b := true
	if food := "apple1111"; b {
		fmt.Println(food, "a")
	} else {
		fmt.Println(food, "b")
	}
}

// InitStatementErrInvalidCode 初始化语句错误，无效代码
func InitStatementErrInvalidCode() {
	b := true
	if food := "apple"; b {
		fmt.Println("实物=", food)
	}
	// 未定义的变量 food
	// fmt.Println(food)

}

// IfElse 用来练习 if else,分支语句
func IfElse() {
	// 条件为 假（false）时，运行else中的代码
	if false {
		fmt.Println("第一个输出语句")
	} else {
		fmt.Println("第二个输出语句")
	}
}

// 在IfElseIfElse 中练习 if  else if  else
// 条件为真时，运行if中的代码
// 条件为假时，运行else中的代码
func IfElseIfElse() {
	// 如果，那么如果，那么
	if false {
		fmt.Println("第一个输出语句")
	} else if true {
		fmt.Println("第二个输出语句")

	} else {
		fmt.Println("第三个输出语句")

	}
}

// IfElseIfElseIfElse 用来练习 if else if else if else
func IfElseIfElseIfElse() {
	if false {
		fmt.Println("这里的判断条件是:false")
	} else if false {
		fmt.Println("这里的判断条件是:false")
	} else if true {
		fmt.Println("这里的判断条件是:true")
	} else {
		fmt.Println("这里是最后")
	}
}

// 能被3整除的数
func DivisibleByThree() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("i = ", i)
		}
	}
}
