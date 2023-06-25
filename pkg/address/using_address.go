package address

import "fmt"

const metersToYards float64 = 1.09361

// 使用变量的内存地址
func UsingMain() {

	fmt.Print(1, "aa")
	var meters float64
	fmt.Print("Enter meters swam:")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, "meter is", yards, "years.")
}
