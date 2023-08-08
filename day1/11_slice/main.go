package main

import "fmt"

func main() {
	// 切片定义
	var slice1 []int
	fmt.Println(slice1 == nil)
	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(slice1)
	// make
	slice2 := make([]string, 3, 6)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	slice2 = append(slice2, "tom")
	slice3 := []string{"jack", "mary"}
	slice2 = append(slice2, slice3...)
	fmt.Println(slice2)
	for _, v := range slice1 {
		fmt.Println(v)
	}
}
