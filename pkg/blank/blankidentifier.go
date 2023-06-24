// blank 用来练习 _ 空白标识符
package blank

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func BlankMain() {
	a := "这里存储a"
	// b := "这里存储b"  // B没有被使用，不然会报错
	fmt.Println("a - ", a)
}

// 带有错误检查的例子
func Checking() {
	res, err := http.Get("http://www.geekwiseacademy.com/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}

func NotChecking() {
	res, _ := http.Get("http://www.geekwiseacademy.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
