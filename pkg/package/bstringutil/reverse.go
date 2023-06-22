// stringutil 在这里用于处理字符串，同时让别的包，调用该包中的函数
package bstringutil

//Reverse 从左到右返回其参数字符串反转符文。
func Reverse(s string) string {
	return reverseTwo(s)
}
