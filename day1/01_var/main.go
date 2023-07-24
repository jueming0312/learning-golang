package main

import "fmt"

func main() {
	//第一种方式
	var name string = "张三"

	// 第二种方式类型推导
	var age = 18
	IpAddress := "10.1.1.1"
	fmt.Println(name, age, IpAddress)

	// 定义多个变量
	var (
		a, b, c = "李四", 13, "上海"
	)
	fmt.Println(a, b, c)
}
