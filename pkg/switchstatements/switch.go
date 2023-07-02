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
		无默认 fallthrough
		fallthrough是可选的
		——您可以通过显式声明来指定故障
		不像其他语言那样需要break
	*/
}
