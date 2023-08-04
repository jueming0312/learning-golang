package main

import "fmt"

func main() {
	// 数组定义
	var a [5]int
	var b [3]int

	b[0] = 3
	b[1] = 6
	b[2] = 9
	fmt.Println(a)
	fmt.Println(b)

	var num1 = [3]int{1, 2, 3}
	var num2 = [...]string{"tom", "jack", "mary"}
	num3 := [3]float64{1.23, 3.24, 56.11}
	fmt.Println(num1, num2, num3)

	for i, v := range num3 {
		fmt.Printf("index=%d value=%v\n", i, v)
	}

}
