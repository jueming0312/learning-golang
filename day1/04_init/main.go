package main

import "fmt"

var (
	user string
	home string
)

func init() {
	if user == "" {
		fmt.Println("user not set") // 这句话会在main函数执行前被输出
	}
	if home == "" {
		home = "/home" + user
	}
}
func main() {
	fmt.Printf("user=%s home=%s\n", user, home)
}
