// remainder 是练习求余数
package remainder

//
import "fmt"

// Remainder 求余数
func Remainder() {
	x := 13 % 3
	fmt.Println("x=", x)

	if x == 1 {
		fmt.Println("是奇数")
	} else {
		fmt.Println("是偶数")
	}
	fmt.Println("分割线")
	for i := 1; i < 70; i++ {
		if i%2 == 1 {
			fmt.Printf("i=%v\t", i)
			fmt.Println("奇数")
		} else {
			fmt.Printf("i=%v\t", i)

			fmt.Println("偶数")
		}
	}
}
