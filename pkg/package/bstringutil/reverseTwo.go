// bstringutil 用于演示，函数在同一个包中使用的情况
package bstringutil

// reverseTwo 首字母小写，只能在同一个包中使用
// reverseTwo 不能被其他包使用
func reverseTwo(s string) string {
	// 定义一个切片 rune的别名是int32
	r := []rune(s)

	// 定义两个变量 i和j
	// i赋值为0 是切片的第一个索引
	// j赋值为len(r)-1 是切片的最后一个索引
	// i < len(r)/2  i 表示这个切片 前半部分的索引
	// 每循环一次，让改变i和j的值，i+1表示 切片的索引从前往后移动1个，  j-1表示 切片的索引从后往前移动1个
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		// 交换切片的元素值
		// 每迭代一次，同时把切片中以后面索引的值，给前面索引。同理，前面索引的值，给后面索引，
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
