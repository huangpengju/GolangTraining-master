// switch包用来练习Switch的使用
package switchstatements

import "fmt"

func SwitchMain() {
	switch "Mhi" {
	case "Daniel":
		fmt.Println("准备好了 Daniel")
	case "Medhi":
		fmt.Println("准备好了 Medhi")
	case "Jenny":
		fmt.Println("准备好了 Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
	/*
		本例子中没有 fallthrough
		fallthrough是可选的
		——您可以通过显式声明来指定fallthrough 穿透，下面的例子中具体展示
		不像其他语言那样需要break
	*/
}

// FallthoughTest 用于练习fallthrough
// “fallthrough:用于穿透switch
// 当switch中某个case匹配成功之后,就执行该case语句
// 如果遇到fallthrough,那么后面紧邻的case,
// 无需匹配, 执行穿透执行。
//  fallthrough应该位于某个case的最后一行”
func FallthroughTest() {
	switch "周三" {
	case "周一":
		fmt.Println("周一准备好了")
	case "周二":
		fmt.Println("周二准备好了")
	case "周三":
		fmt.Println("周三准备好了")
		fallthrough
	case "周四":
		fmt.Println("周四准备好了")
		fallthrough
	case "周五":
		fmt.Println("周五准备好了")
	case "周六":
		fmt.Println("周六准备好了")

	}

}

// MultipleEvals 多个判断条件
func MultipleEvals() {
	switch "周二" {
	case "周一", "周二":
		fmt.Println("周一准备好了，还有一个错误时间，周二")
	case "周三", "周四":
		fmt.Println("这里有周三和周四")
	case "周五", "周六":
		{
			fmt.Println("周五 周六准备好了")
		}
	}
}

// NoExpression 是练习没有表达式的switch
func NoExpression() {
	myName := "ju"
	switch {
	case len(myName) == 3:
		fmt.Println("我的名字长度是2")
	case myName == "ju1":
		fmt.Println("ju1准备好了")
	case myName == "ju2":
		fmt.Println("ju2准备好了")
	case myName == "ju3", myName == "ju4":
		fmt.Println("我的名字是ju3或者ju4")
	case myName == "ju5":
		fmt.Println("ju5准备好了")
	case myName == "ju6":
		fmt.Println("ju6准备好了")
	default:
		fmt.Println("没有找到你的名字，这里是default")
	}

}

// switch 使用类型
// 通常我们在switch中用变量的值
// go允许你在switch中用变量的类型

// contact联系结构体类型
type contact struct {
	greeting string //问候
	name     string
}

// SwitchOnType工作于接口
// 我们将在后面学习更多关于接口的知识
func switchOnType(x interface{}) {
	switch x.(type) { //这是一个断言;断言"x是这种类型"
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case contact:
		fmt.Println("contact结构类型")
	default:
		fmt.Println("未知的")
	}
}
func OnType() {
	// 判断是不是int
	switchOnType(7)
	// 判断是不是string
	switchOnType("abcd")
	// 判断是不是结构体
	var t = contact{"你最近怎么样", "周一"}
	switchOnType(t)
	// 判断是不是string
	switchOnType(t.greeting)
	// 判断是不是string
	switchOnType(t.name)
}
