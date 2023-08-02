package main

import "fmt"

func main() {
	// 切片定义
	var slice1 []int
	fmt.Println(slice1 == nil)
	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(slice1)

}
