package main

import "fmt"

func main() {
	name := "张三"
	age := 22
	pi := 3.1415926
	fmt.Println(name, age)

	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("pi=%.2f\n", pi)

	res := fmt.Sprintf("%s is %d years old.", name, age)
	fmt.Println(res)
}
