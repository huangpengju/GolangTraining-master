package address

import "fmt"

// showingMain 用来学习获取变量的内存地址，获得的值是十六进制数
func ShowingMain() {
	a := 43

	fmt.Println("a=", a)
	fmt.Println("a的内存地址-", &a)
	fmt.Printf("十进制%d\n", &a)

	var b = "aaa"
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Printf("%v\n", &b)
	fmt.Printf("%d\n", &b)
	fmt.Printf("%b", &b)
}
